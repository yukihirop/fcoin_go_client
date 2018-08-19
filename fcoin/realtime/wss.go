package fcoin_realtime

import (
	"fmt"

	"github.com/gorilla/websocket"
)

const (
	writeWait  = 10
	pingPeriod = 10
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	conn         *websocket.Conn
	readMessages chan map[string]interface{}
	topic        string
	callbacks    map[string]HandleFunc
}

func (conf *Configure) createWebSocketClient(topic string, handler HandleFunc) (client *Client) {
	conn, _, err := websocket.DefaultDialer.Dial(conf.WssEndpoint, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	client = &Client{conn: conn, readMessages: make(chan map[string]interface{}), topic: topic, callbacks: map[string]HandleFunc{}}
	client.callbacks[topic] = handler
	return
}

func (c *Client) read() {
	defer func() {
		c.conn.Close()
	}()
	for {
		message := map[string]interface{}{}
		err := c.conn.ReadJSON(&message)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Println("error: %v", err)
			}
			break
		} else {
			topic, _ := message["type"].(string)
			if topic == c.topic {
				c.readMessages <- message
			}
		}
	}
}

func (c *Client) write(message map[string]interface{}) {
	if err := c.conn.WriteJSON(message); err != nil {
		return
	}
}

func (c *Client) callCallback(topic string) {
	cb := c.callbacks[topic]
	cb(<-c.readMessages)
}
