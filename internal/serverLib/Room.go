package serverLib

import (
	Utils "sevens/internal/utils"
	"sync"
)

type Room struct {
	roomName          string
	roomId            string
	users             map[*User]bool
	maxPlayer         int
	mutex             sync.RWMutex
	game              *Game
	gameStart         bool
	gameSignal        chan (struct{})
	roomManagerSignal chan (struct{})
}

type RoomInfo struct {
	RoomId     string `json:"roomId"`
	RoomName   string `json:"roomName"`
	UsersCount int    `json:"usersCount"`
	CanJoin    bool   `json:"canJoin"`
	MaxPlayer  int    `json:"maxPlayer"`
}

const (
	MaxPlayer = 4
)

func NewRoom(roomName string, roomManagerSignal chan struct{}) *Room {
	if roomName == "" {
		roomName = "Let's play a game"
	}
	return &Room{
		roomName:          roomName,
		roomId:            generateRoomId(),
		users:             make(map[*User]bool),
		maxPlayer:         MaxPlayer,
		gameStart:         false,
		gameSignal:        make(chan struct{}),
		roomManagerSignal: roomManagerSignal,
	}
}

func generateRoomId() string {
	return Utils.GenerateRandomString(6)
}

func (room *Room) Broadcast(message string) {
	for user := range room.users {
		user.Write(message)
	}
}

func (room *Room) GetRoomInfo() *RoomInfo {
	defer room.mutex.Unlock()
	room.mutex.Lock()
	canJoin := room.CanJoin()
	return &RoomInfo{
		RoomId:     room.roomId,
		RoomName:   room.roomName,
		UsersCount: len(room.users),
		CanJoin:    canJoin,
		MaxPlayer:  room.maxPlayer,
	}
}

func (room *Room) CanJoin() bool {
	if len(room.users) < MaxPlayer && !room.gameStart {
		return true
	}
	return false
}

func (room *Room) StartGame() {
	defer room.mutex.Unlock()
	room.mutex.Lock()
	if room.gameStart {
		return
	}

	for user, _ := range room.users {
		user.SetStatus(UserInGame)
	}
	game := NewGame(room.users, room.maxPlayer)
	room.game = game
	room.gameStart = true

	go game.Start(room.gameSignal) // hack for game start
	go room.listenSignal()

}

func (room *Room) listenSignal() {
	<-room.gameSignal

	room.mutex.Lock()
	room.gameStart = false
	for user := range room.users {
		user.SetStatus(UserInRoom)
	}
	room.mutex.Unlock()
	room.roomManagerSignal <- struct{}{}
}

func (room *Room) JoinRoom(user *User) bool {
	room.mutex.Lock()
	defer room.mutex.Unlock()

	if !room.CanJoin() {
		return false
	}
	room.users[user] = true
	user.SetStatus(UserInRoom)

	return true
}

func (room *Room) LeaveRoom(user *User) bool {
	room.mutex.Lock()
	defer room.mutex.Unlock()
	if room.gameStart {
		return false
	}
	if _, ok := room.users[user]; ok {
		delete(room.users, user)
		return true
	}
	return false
}
