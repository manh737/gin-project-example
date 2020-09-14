package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/manh737/gin-project-example/entity"
	"github.com/manh737/gin-project-example/service"
)

type VideoController interface {
	GetAll() []entity.Video
	Save(ctx *gin.Context) error
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) GetAll() []entity.Video {
	return c.service.GetAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	if err := ctx.ShouldBindJSON(&video); err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}
