package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

func currentContext() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Printf("$PORT either undeclared or empty, falling back to default port %s\n", port)
	}
	return ":" + port
}

func keepAlive(c *websocket.Conn, timeout time.Duration) {
	lastResponse := time.Now()
	c.SetPongHandler(func(msg string) error {
		lastResponse = time.Now()
		return nil
	})

	go func() {
		for {
			err := c.WriteMessage(websocket.PingMessage, []byte("keepalive"))
			if err != nil {
				return
			}
			time.Sleep(timeout / 2)
			if time.Now().Sub(lastResponse) > timeout {
				c.Close()
				return
			}
		}
	}()
}
