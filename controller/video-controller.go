package controller

import (
	"github.com/Ammce/pilot/entity"
	"github.com/Ammce/pilot/service"
	"github.com/Ammce/pilot/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
	FindById(ctx *gin.Context) string
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-adult", validators.IsAdultValidator)
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	err = validate.Struct(&video)
	if err != nil {
		return err
	}

	c.service.Save(video)
	return nil
}

func (c *controller) FindById(ctx *gin.Context) string {
	var Id string = "Hello"
	return c.service.FindById(&Id)
}
