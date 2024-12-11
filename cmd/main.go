package main

import (
	"fmt"
	"log"

	"github.com/Orken1119/Websocket/internal/controller"
	"github.com/Orken1119/Websocket/pkg"
	"github.com/gin-gonic/gin"
)

func main() {
	app, err := pkg.App()

	if err != nil {
		log.Fatal(err)
	}
	defer app.CloseDBConnection()

	ginRouter := gin.Default()

	controller.Setup(app, ginRouter)

	ginRouter.Run(fmt.Sprintf(":%d", 1150))
}
