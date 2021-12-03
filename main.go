package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/websocket"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Static("/","../public")
	e.GET("/ws",hello)
	e.Logger.Fatal(e.Start(":1111"))
}

func hello(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			err := websocket.Message.Send(ws, "Hello,client!")
			if err != nil {
				c.Logger().Error(err)
			}

			msg := ""
			err = websocket.Message.Receive(ws,&msg)
			if err != nil {
				c.Logger().Error(err)
			}
			fmt.Println(msg)
		}
	}).ServeHTTP(c.Response(),c.Request())
	return nil
}


