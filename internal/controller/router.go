package controller

import (
	"github.com/Orken1119/Websocket/internal/controller/middleware"
	"github.com/Orken1119/Websocket/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gorilla/websocket"

	"github.com/Orken1119/Websocket/internal/controller/auth"
	"github.com/Orken1119/Websocket/internal/repository"
)

func Setup(app pkg.Application, router *gin.Engine) {
	db := app.DB

	authController := &auth.AuthController{
		UserRepository: repository.NewUserRepository(db),
	}

	router.POST("/signup", authController.Signup)
	router.POST("/signin", authController.Signin)


	http.HandleFunc("/ws", handler)
	router.Use(middleware.JWTAuth(`access-secret-key`))
}



var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return 
	}

	defer conn.Close()

	go handleConnection(conn)
}

func handleConnection(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			return 
		}

		go func (msg []byte) {
			err = conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				return
			}
		}(message)
	}
}