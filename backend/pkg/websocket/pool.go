package websocket

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan map[uint]struct{}
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan map[uint]struct{}),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
		case winner := <-pool.Broadcast:
			for client := range pool.Clients {
				if _, exists := winner[client.Number]; exists {
					client.Conn.WriteJSON(Response{Status: "success", Message: "win"})
				}
			}
		}
	}
}
