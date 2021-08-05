package main

import (
	"github.com/Ammce/pilot/controller"
	"github.com/Ammce/pilot/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/posts", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())
	})

	server.POST("/posts", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.Save(ctx))
	})

	server.Run(":8080")
}
