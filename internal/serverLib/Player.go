package serverLib

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Player struct {
	isRobot      bool
	User         *User
	PlayerAssets *PlayerAssets
	position     int
	game         *Game
	gameAction   chan *GameAction
	endGame      chan struct{}
	myRound      chan struct{}
}

func NewPlayer(user *User) *Player {
	isRobot := false
	if user == nil {
		isRobot = true
	}
	player := &Player{
		isRobot:      isRobot,
		User:         user,
		gameAction:   make(chan *GameAction),
		PlayerAssets: NewPlayerAssets(),
		endGame:      make(chan struct{}),
		myRound:      make(chan struct{}),
	}
	return player
}

const (
	RobotAutoPlay = 100 //ms
	HumanAutoPlay = 100 //ms
)

func (p *Player) Init(ctx context.Context) {
	if !p.isRobot {
		go p.ServeCmd(ctx)
	}
	go p.ServeRound(ctx)
}

func (p *Player) ServeRound(ctx context.Context) {
	run := true
	for run {
		select {
		case <-ctx.Done():
			fmt.Println("got the stop channel")
			return
		case <-p.myRound:
			if p.game.GameTable.roundRemain < 4 {
				p.game.GameTable.roundRemain -= 1
			}
			p.RoundStart()
			if p.game.GameTable.roundRemain == 0 {
				run = false
				break
			}
			p.game.turn = (p.game.turn + 1) % p.game.maxPlayer
			p.game.Players[p.game.turn].myRound <- struct{}{}
		}
	}
	// for _, p := range p.game.Players {
	// 	fmt.Println(p.PlayerAssets.Score)
	// }
}

func (p *Player) RoundStart() {
	IntervalTime := HumanAutoPlay * time.Millisecond // 觸發間隔時間
	ticker := time.NewTicker(IntervalTime)           // 設定 秒觸發一次
	select {
	case gameAction := <-p.gameAction:
		if gameAction.actionCode == GameActionPlay {
		}
		p.game.Action <- gameAction
	case <-ticker.C:
		p.AutoPlay()
	}
}

func (p *Player) ServeCmd(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case cmd := <-p.User.gameCmd:
			p.handleGameCmd(cmd)
		}
	}
}

func (p *Player) handleGameCmd(cmd string) {
	res := strings.SplitN(cmd, " ", 2)
	switch res[0] {
	case "game_play_card":
		p.cmdGamePlayCard(res[1])
	}
}

func (p *Player) cmdGamePlayCard(data string) {
	if !p.game.CheckTurn(p.position) {
		return // not your turn
	}
	var cmdInGamePlayCard CmdInGamePlayCard
	err := json.Unmarshal([]byte(data), &cmdInGamePlayCard)
	if err != nil {
		return
	}
	playCard := cmdInGamePlayCard.Card

	ga := &GameAction{
		actionCode: GameActionPlay,
		card:       playCard,
	}
	timer := time.NewTimer(500 * time.Millisecond) //防呆 應該能快速處理完?
	select {
	case p.gameAction <- ga:
		break
	case <-timer.C:
		break
	}
}

func (p *Player) SetGame(game *Game) {
	p.game = game
}

func (p *Player) SetPosition(position int) {
	p.position = position
}

func (p *Player) GetPosition() int {
	return p.position
}

func (p *Player) AutoPlay() {
	defer p.game.GameTable.mutex.RUnlock()
	defer p.PlayerAssets.mutex.RUnlock()
	p.game.GameTable.mutex.RLock()
	p.PlayerAssets.mutex.RLock()
	cardLevels := p.getCanBuyCards()
	cards := []*Card{}
	for _, cardLevel := range cardLevels {
		if len(cardLevel) > 0 {
			cards = append(cards, cardLevel...)
		}
	}
	if len(cards) > 0 {
		randomIndex := rand.Intn(len(cards))
		p.buyCard(cards[randomIndex])
		noble := p.getValidNoble()
		if len(noble) > 0 {
			rand.Seed(time.Now().UnixNano())
			index := rand.Intn(len(noble))
			p.getNoble(noble[index])
		}
		p.game.Broadcast(p.game.GameTable.reportDeck())
	} else {
		p.randomTakeActions()
		p.game.Broadcast(p.game.GameTable.reportDeck())
	}
	return
}

func (p *Player) randomTakeActions() {

	validOne := []Gems{} //valid gems for take one of each type
	validTwo := []Gems{} //valid gems for take two

	// card could be reserved
	cardMap := map[int][]*Card{0: p.game.GameTable.cardLow, 1: p.game.GameTable.cardMid, 2: p.game.GameTable.cardHigh}
	cardCandidate := []*Card{}
	for i := 0; i < 3; i += 1 {
		for _, card := range cardMap[i] {
			if card == nil {
				continue
			}
			cardCandidate = append(cardCandidate, card)
		}
	}
	validReserve := p.game.GameTable.Gems[Gold] > 0 && len(p.PlayerAssets.ReservedCard) < 3 && len(cardCandidate) > 0

	for k, v := range p.game.GameTable.Gems {
		if v == 0 {
			continue
		}
		if k != Gold {
			validOne = append(validOne, k)
			if v >= 4 {
				validTwo = append(validTwo, k)
			}
		}
	}

	actions := []int{}
	if len(validOne) >= 3 {
		actions = append(actions, 0) // take 3 different types of gems
	}
	if len(validTwo) >= 1 {
		actions = append(actions, 1) // take 2 gems of same type
	}
	if validReserve {
		actions = append(actions, 2) // reserve card
	}
	var validCards []*Card
	if len(p.PlayerAssets.ReservedCard) > 0 {
		validCards = p.getCanBuyReservedCards()
		if len(validCards) > 0 {
			actions = append(actions, 3) // buy reserve card
		}
	}

	rand.Seed(time.Now().UnixNano())
	// h := map[Gems]int{}
	// for k, v := range p.PlayerAssets.Deduct {
	// 	h[k] = v + p.PlayerAssets.Gems[k]
	// }
	// if len(actions) == 0 {
	// 	fmt.Println(h)
	// 	fmt.Println(p.game.GameTable.Gems)
	// 	fmt.Println(validOne)
	// 	for _, card := range p.PlayerAssets.ReservedCard {
	// 		fmt.Println(card.Cost)
	// 	}
	// }

	randomIndex := rand.Intn(len(actions))

	switch actions[randomIndex] {
	case 0:
		if len(validOne) == 3 {
			for _, gemName := range validOne {
				p.PlayerAssets.Gems[gemName] += 1
				p.game.GameTable.Gems[gemName] -= 1
			}
		} else {
			rand.Seed(time.Now().UnixNano())
			rand.Shuffle(len(validOne), func(i, j int) {
				validOne[i], validOne[j] = validOne[j], validOne[i]
			})
			for i := 0; i < 3; i += 1 {
				gemName := validOne[i]
				p.PlayerAssets.Gems[gemName] += 1
				p.game.GameTable.Gems[gemName] -= 1
			}
		}
	case 1:
		if len(validTwo) == 1 {
			gemName := validTwo[0]
			p.PlayerAssets.Gems[gemName] += 2
			p.game.GameTable.Gems[gemName] -= 2
		} else {
			rand.Seed(time.Now().UnixNano())
			validTwoIndex := rand.Intn(len(validTwo))
			gemName := validTwo[validTwoIndex]
			p.PlayerAssets.Gems[gemName] += 2
			p.game.GameTable.Gems[gemName] -= 2
		}
	case 2:
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(len(cardCandidate))
		p.reserveCard(cardCandidate[index])
	case 3:
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(len(validCards))
		p.buyCard(validCards[index])
		noble := p.getValidNoble()
		if len(noble) > 0 {
			rand.Seed(time.Now().UnixNano())
			index := rand.Intn(len(noble))
			p.getNoble(noble[index])
		}
	}

	return
}

func (p *Player) takeGems(take map[Gems]int) bool {
	p.PlayerAssets.mutex.Lock()
	defer p.PlayerAssets.mutex.Unlock()
	p.game.GameTable.mutex.Lock()
	defer p.game.GameTable.mutex.Unlock()

	if v, ok := take[Gold]; ok && v > 0 {
		return false // disable to take gold, use reverse card method
	}

	single := 0
	double := 0
	for k, v := range take {
		if v == 1 {
			single += 1
		} else if v == 2 {
			double += 1
		} else if v != 0 {
			return false
		}
		if p.game.GameTable.Gems[k] < v || v == 2 && p.game.GameTable.Gems[k] < 4 {
			return false
		}
	}
	if !(single == 3 && len(take) == 3 || double == 1 && len(take) == 1) {
		return false
	}
	for k, v := range take {
		p.game.GameTable.Gems[k] -= v
		p.PlayerAssets.Gems[k] += v
	}
	return true
}

// caller handle lock
func (p *Player) reserveCard(card *Card) bool {
	if len(p.PlayerAssets.ReservedCard) >= 3 {
		return false
	}
	if !p.game.GameTable.getGold() {
		return false
	}

	p.PlayerAssets.Gems[Gold] += 1
	p.PlayerAssets.ReservedCard = append(p.PlayerAssets.ReservedCard, card)
	p.PlayerAssets.ReservedCardPosition[card.Number] = len(p.PlayerAssets.ReservedCard) - 1

	index := p.game.GameTable.cardNumberPosition[card.Number]
	delete(p.game.GameTable.cardNumberPosition, card.Number)
	if card.Number >= CardHighStartNumber {
		p.game.GameTable.DrawCard(&p.game.GameTable.deckHigh, p.game.GameTable.cardHigh, index)
	} else if card.Number >= CardMidStartNumber {
		p.game.GameTable.DrawCard(&p.game.GameTable.deckMid, p.game.GameTable.cardMid, index)
	} else {
		p.game.GameTable.DrawCard(&p.game.GameTable.deckLow, p.game.GameTable.cardLow, index)
	}
	return true
}

func (p *Player) buyReservedCard(card *Card) bool {
	p.game.GameTable.mutex.Lock()
	defer p.game.GameTable.mutex.Unlock()
	p.PlayerAssets.mutex.Lock()
	defer p.PlayerAssets.mutex.Unlock()
	var cardIndex int
	var ok bool
	if cardIndex, ok = p.PlayerAssets.ReservedCardPosition[card.Number]; !ok {
		return false
	}
	if !p.accumulateBuying(card) {
		return false
	}
	p.PlayerAssets.ReservedCard = append(p.PlayerAssets.ReservedCard[:cardIndex], p.PlayerAssets.ReservedCard[cardIndex+1:]...)
	delete(p.PlayerAssets.ReservedCardPosition, card.Number)

	return true
}

func (p *Player) getCanBuyReservedCards() []*Card {
	defer p.PlayerAssets.mutex.RUnlock()
	p.PlayerAssets.mutex.RLock()
	if len(p.PlayerAssets.ReservedCard) == 0 {
		return nil
	}
	arr := []*Card{}
	for _, card := range p.PlayerAssets.ReservedCard {
		if p.canBuyCard(card) {
			arr = append(arr, card)
		}
	}
	return arr
}

func (p *Player) buyCard(card *Card) bool {
	if _, ok := p.game.GameTable.cardNumberPosition[card.Number]; !ok {
		return false
	}

	if !p.accumulateBuying(card) {
		return false
	}

	index := p.game.GameTable.cardNumberPosition[card.Number]
	delete(p.game.GameTable.cardNumberPosition, card.Number)
	if card.Number >= CardHighStartNumber {
		p.game.GameTable.DrawCard(&p.game.GameTable.deckHigh, p.game.GameTable.cardHigh, index)
		// if len(p.game.GameTable.deckHigh) > 0 {
		// 	newCard := p.game.GameTable.deckHigh[len(p.game.GameTable.deckHigh)-1]
		// 	p.game.GameTable.cardNumberPosition[newCard.Number] = index
		// 	p.game.GameTable.deckHigh = p.game.GameTable.deckHigh[:len(p.game.GameTable.deckHigh)-1]
		// 	p.game.GameTable.cardHigh[index] = newCard
		// } else {
		// 	p.game.GameTable.cardHigh[index] = nil
		// }
	} else if card.Number >= CardMidStartNumber {
		p.game.GameTable.DrawCard(&p.game.GameTable.deckMid, p.game.GameTable.cardMid, index)
	} else {
		p.game.GameTable.DrawCard(&p.game.GameTable.deckLow, p.game.GameTable.cardLow, index)
	}

	return true
}

func (p *Player) accumulateBuying(card *Card) bool {
	gold := p.PlayerAssets.Gems[Gold]
	cost := map[Gems]int{}
	for k, v := range card.Cost {
		total := p.PlayerAssets.Gems[k] + p.PlayerAssets.Deduct[k]
		if total+gold < v {
			return false
		}
		if total >= v {
			v -= p.PlayerAssets.Deduct[k]
			if v > 0 {
				cost[k] += v // force buy without gold
			}
			continue // todo: buy with gold
		}
		goldSubstitute := v - total
		gold = gold - goldSubstitute
		cost[Gold] = goldSubstitute
	}
	for k, v := range cost {
		p.PlayerAssets.Gems[k] -= v
		p.game.GameTable.Gems[k] += v
	}
	p.PlayerAssets.Score += card.Score
	p.PlayerAssets.Deduct[card.Deduct] += 1
	p.PlayerAssets.Cards = append(p.PlayerAssets.Cards, card)
	return true
}

func (p *Player) getCanBuyCards() [][]*Card {
	p.game.GameTable.mutex.RLock()
	defer p.game.GameTable.mutex.RUnlock()
	p.PlayerAssets.mutex.RLock()
	defer p.PlayerAssets.mutex.RUnlock()
	deck := [][]*Card{[]*Card{}, []*Card{}, []*Card{}} // contains low,mid,high cards you can buy
	cardMap := map[int][]*Card{0: p.game.GameTable.cardLow, 1: p.game.GameTable.cardMid, 2: p.game.GameTable.cardHigh}

	for i := 0; i < 3; i += 1 {
		for _, card := range cardMap[i] {
			if card == nil {
				continue
			}
			if p.canBuyCard(card) {
				deck[i] = append(deck[i], card)
			}
		}
	}
	return deck
}

func (p *Player) canBuyCard(card *Card) bool {
	gold := p.PlayerAssets.Gems[Gold]
	valid := true
	for k, v := range card.Cost {
		total := p.PlayerAssets.Gems[k] + p.PlayerAssets.Deduct[k]
		if total+gold < v {
			valid = false
			break
		}
		if total < v {
			gold = gold - (v - total)
		}
	}
	return valid
}

func (p *Player) getValidNoble() []*Noble {
	if len(p.game.GameTable.noble) == 0 {
		return nil
	}
	arr := []*Noble{}
	for _, card := range p.game.GameTable.noble {
		if card == nil {
			continue //todo: nil maybe a bad practice
		}
		valid := true
		for k, v := range card.Cost {
			if p.PlayerAssets.Deduct[k] < v {
				valid = false
				break
			}
		}
		if valid {
			arr = append(arr, card)
		}
	}
	return arr
}

func (p *Player) getNoble(noble *Noble) {
	p.PlayerAssets.Score += noble.Score
	p.PlayerAssets.Noble = append(p.PlayerAssets.Noble, noble)
	index := p.game.GameTable.nobleNumberPosition[noble.Number]
	p.game.GameTable.noble[index] = nil
	delete(p.game.GameTable.nobleNumberPosition, noble.Number)
	if p.PlayerAssets.Score >= 15 && p.game.GameTable.roundRemain == 4 {
		p.game.GameTable.roundRemain -= 1
		fmt.Println("reach15")
	}
}
