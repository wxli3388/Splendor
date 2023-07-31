package serverLib

import (
	"context"
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type Game struct {
	Players   []*Player
	GameTable *GameTable
	Action    chan *GameAction
	finish    chan struct{}

	turn      int
	maxPlayer int
	mutex     sync.RWMutex
}

const (
	GameActionPlay  = 1
	GameActionCover = 2
)

type GameAction struct {
	actionCode int
	card       string
}

func NewGame(userMap map[*User]bool, maxPlayer int) *Game {
	var players []*Player
	for user, _ := range userMap {
		players = append(players, NewPlayer(user))
	}
	for i := len(players); i < maxPlayer; i += 1 {
		players = append(players, NewPlayer(nil))
	}

	game := &Game{
		Players:   players,
		Action:    make(chan *GameAction),
		finish:    make(chan struct{}),
		maxPlayer: maxPlayer,
		GameTable: NewGameTable(maxPlayer),
		turn:      0,
	}
	for i := 0; i < maxPlayer; i += 1 {
		game.Players[i].SetGame(game)
	}
	return game
}

func (game *Game) Start(gameSignal chan struct{}) {
	game.RandomPlayer()
	fmt.Println("game start")
	game.Broadcast("gameStart")
	game.Broadcast(game.GameTable.reportDeck())
	go game.Run(gameSignal)
	game.Players[game.turn].myRound <- struct{}{}
}

func (game *Game) Run(gameSignal chan struct{}) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < game.maxPlayer; i += 1 {
		game.Players[i].Init(ctx)
	}
	// go game.Players[game.turn].RoundStart()
	run := true
	for run {
		select {
		case action := <-game.Action:
			if action.actionCode == GameActionPlay {
			}

		case <-game.finish:
			run = false
			break
		}
	}
	cancel()

	coverCard := []string{}
	score := []int{}

	gameOver := &CmdOutGameOver{
		CoverCard: coverCard,
		Score:     score,
	}
	game.Broadcast(gameOver)
	gameSignal <- struct{}{}
	for i := 0; i < len(game.Players); i += 1 {
		if !game.Players[i].isRobot {
			game.Players[i].endGame <- struct{}{}
		}
	}

	timer := time.NewTimer(3 * time.Second) //wait for game over score board
	<-timer.C
	game.Broadcast(&CmdOutBackToRoom{})

}

func (game *Game) RandomPlayer() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(game.Players), func(i, j int) {
		game.Players[i], game.Players[j] = game.Players[j], game.Players[i]
	})

	for k, player := range game.Players {
		player.SetPosition(k)
	}
	sort.Slice(game.Players, func(i, j int) bool {
		return game.Players[i].GetPosition() < game.Players[j].GetPosition()
	})
}

func (game *Game) Broadcast(message any) {
	for _, player := range game.Players {
		if !player.isRobot {
			player.User.Write(message)
		}
	}
}

func (game *Game) CheckTurn(turn int) bool {
	defer game.mutex.RUnlock()
	game.mutex.RLock()
	return game.turn == turn
}
