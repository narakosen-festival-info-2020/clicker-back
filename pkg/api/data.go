package api

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/gorilla/websocket"
	"github.com/narakosen-festival-info-2020/clicker-back/pkg/clicker"
)

// App is websocket app data
type App struct {
	clickerData *clicker.Data
	totalUsers  int // sync target
	clients     map[*websocket.Conn]bool
	upgrader    websocket.Upgrader
	sync.RWMutex
}

// Generate is websocket app generate
func Generate() App {
	return App{
		clickerData: clicker.Generate(),
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
		log.Println(err)
		return
	}

	defer ws.Close()
	app.Lock()
	app.totalUsers++
	app.Unlock()

	app.clients[ws] = true

	type receivedMessage struct {
		Count float64 `json:"count"`
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
		go app.clickerData.AddClick(msg.Count)
	}
}

func (app *App) handleMessages() {
	for {
		// 現在接続しているクライアント全てにメッセージを送信する
		for client := range app.clients {
			time.Sleep(time.Second / time.Duration((10 * len(app.clients))))
			err := client.WriteJSON(app.clickerData.GetJSON())
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
	server := gin.Default()

	// CORS setup
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{
		"https://clicker.nitncfes.net",
		// "http://localhost:3000",
	}
	server.Use(cors.New(corsConfig))

	facilityRoute(server, app)
	statementsRoute(server, app)

	app.clickerData.InitFacility()

	generalStatements := make(map[string]func() float64)
	clickStatements := make(map[string]func() float64)
	generalStatements["total_users"] = func() float64 {
		app.RLock()
		defer app.RUnlock()
		return (float64)(app.totalUsers)
	}
	generalStatements["now_users"] = func() float64 {
		return (float64)(len(app.clients))
	}
	app.clickerData.InitStatements(generalStatements, clickStatements)

	server.GET(url, func(ctx *gin.Context) {
		app.handleConnections(ctx.Writer, ctx.Request)
	})

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})
	go app.handleMessages()
	server.Run(":80")
}
