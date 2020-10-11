package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var counterChan = make(chan int)
var broadcast = make(chan int)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type ReceivedMessage struct {
	Count int `json:"count"`
}

type SendMessage struct {
	Count int `json:"count"`
}

func app() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	http.HandleFunc("/ws", handleConnections)
	go handleMessages()
	go count()

	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	clients[ws] = true

	for {
		var msg ReceivedMessage

		// 新しいメッセージをJSONとして読み込みMessageオブジェクトにマッピング
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		counterChan <- msg.Count
	}
}

func count() {
	count := 0
	for {
		diff := <-counterChan
		count += diff
		broadcast <- count
	}
}

func handleMessages() {
	for {
		// ブロードキャストチャネルから更新後のcountを受け取る
		count := <-broadcast

		msg := SendMessage{Count: count}

		// 現在接続しているクライアント全てにメッセージを送信する
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
