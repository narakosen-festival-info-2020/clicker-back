package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/narakosen-festival-info-2020/clicker-back/pkg/clicker"
)

// App is websocket app data
type App struct {
	clickerData clicker.Data
	clients     map[*websocket.Conn]bool
	upgrader    websocket.Upgrader
}

// Generate is websocket app generate
func Generate(url string) App {
	return App{
		clickerData: clicker.Data{},
		clients:     make(map[*websocket.Conn]bool),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (app *App) handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := app.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	app.clients[ws] = true

	type receivedMessage struct {
		Count int `json:"count"`
	}

	for {
		var msg receivedMessage

		// 新しいメッセージをJSONとして読み込みMessageオブジェクトにマッピング
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(app.clients, ws)
			break
		}
		go app.clickerData.AddCount(msg.Count)
	}
}

func (app *App) handleMessages() {
	for {
		// 現在接続しているクライアント全てにメッセージを送信する
		for client := range app.clients {
			time.Sleep(time.Second / time.Duration((10 * len(app.clients))))
			err := client.WriteJSON(app.clickerData.GetJSON())
			fmt.Println(app.clickerData.GetJSON())
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(app.clients, client)
			}
		}
	}
}

// Up is Server Start
func (app *App) Up(url string) {
	http.HandleFunc(url, app.handleConnections)
	go app.handleMessages()

	log.Println("http server started on :80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
