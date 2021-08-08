package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Ammce/pilot/controller"
	"github.com/Ammce/pilot/middleware"
	"github.com/Ammce/pilot/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func setupLogOutpu() {
	f, _ := os.Create("logger.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutpu()
	server := gin.New()

	server.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuth())

	server.GET("/posts", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())
	})

	server.GET("/posts/{id}", func(ctx *gin.Context) {
		fmt.Println(ctx.Query("id"))
		ctx.JSON(200, VideoController.FindById(ctx))
	})

	server.POST("/posts", func(ctx *gin.Context) {
		err := VideoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Post input is valid"})
		}
	})

	server.Run(":8080")
}
