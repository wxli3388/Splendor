package serverLib

import "fmt"

type Server struct {
	roomManager *RoomManager
	users       map[*User]bool
	Register    chan *User
	Unregister  chan *User
	Message     chan string
}

func NewServer() *Server {
	roomManager := NewRoomManager()
	server := &Server{
		roomManager: roomManager,
		users:       make(map[*User]bool),
		Register:    make(chan *User),
		Unregister:  make(chan *User),
		Message:     make(chan string),
	}
	roomManager.SetServer(server)
	go roomManager.listen()
	return server
}

func (server *Server) Run() {
	for {
		select {
		case user := <-server.Register:
			server.registerUser(user)

		case user := <-server.Unregister:
			server.unregisterUser(user)
		case msg := <-server.Message:
			for users, _ := range server.users {
				users.Write(msg)
			}
		}
	}
}

func (server *Server) Broadcast(param any) {
	server.Message <- fmt.Sprintf("%v", param)
}

func (server *Server) registerUser(user *User) {
	server.users[user] = true
}

func (server *Server) unregisterUser(user *User) {
	if _, ok := server.users[user]; ok {
		delete(server.users, user)
	}
}
