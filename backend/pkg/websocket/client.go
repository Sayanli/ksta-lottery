package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"lottery/internal/api/handler"

	"github.com/gorilla/websocket"
)

type Client struct {
	Handler *handler.Handler
	Conn    *websocket.Conn
	Pool    *Pool
	Number  uint
}

type Message struct {
	Method string          `json:"method"`
	Body   json.RawMessage `json:"body"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (c *Client) Read(Handler *handler.Handler) {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		var message Message
		err := c.Conn.ReadJSON(&message)
		if err != nil {
			log.Println(err, "failed")
			return
		}
		switch m := message.Method; m {
		case "GetNumberByToken":
			c.GetNumberByToken(&message)

		case "GenerateNumber":
			err := c.GenerateNumberAndToken(&message)
			if err != nil {
				return
			}

		case "LotteryCompleted":
			c.Pool.Broadcast <- c.Handler.LotteryCompleted()
			c.Conn.WriteJSON(Response{Status: "success", Message: "lottery completed"})

		default:
			message := fmt.Sprintf("method '%s' unsupported", m)
			c.Conn.WriteJSON(Response{Status: "error", Message: message})
		}
	}
}

func (c *Client) GetNumberByToken(message *Message) {
	number, err := c.Handler.GetNumberByToken(message.Body)
	if err != nil {
		c.Conn.WriteJSON(Response{Status: "error", Message: err.Error()})
		return
	}
	nt := struct {
		Number uint `json:"number"`
	}{
		number,
	}
	c.Number = number
	c.Conn.WriteJSON(Response{Status: "success", Message: "number", Data: nt})
	if c.Handler.CheckWinner(number) {
		c.Conn.WriteJSON(Response{Status: "success", Message: "win"})
	}
}

func (c *Client) GenerateNumberAndToken(message *Message) error {
	number, token, err := c.Handler.GenerateNumberAndToken()
	if err != nil {
		c.Conn.WriteJSON(Response{Status: "error", Message: err.Error()})
		return err
	}
	c.Number = number
	nt := struct {
		Number uint   `json:"number"`
		Token  string `json:"token"`
	}{
		number,
		token,
	}
	c.Conn.WriteJSON(Response{Status: "success", Message: "number", Data: nt})
	return nil
}
