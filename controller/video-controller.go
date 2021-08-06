package controller

import (
	"github.com/Ammce/pilot/entity"
	"github.com/Ammce/pilot/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
	FindById(ctx *gin.Context) string
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}

func (c *controller) FindById(ctx *gin.Context) string {
	var Id string = "Hello"
	return c.service.FindById(&Id)
}
