package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/user/:name/:action", func(ctx *gin.Context) {
		name := ctx.Param("name")
		action := ctx.Param("action")
		ctx.String(http.StatusOK, "Hello %v parabens pela sua ação de %v", name, action)
	})

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Servidor nao foi iniciado corretamente, erro: %v", err)
	}

}
