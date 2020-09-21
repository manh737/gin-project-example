package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/manh737/gin-project-example/controller"
	"github.com/manh737/gin-project-example/middlewares"
	"github.com/manh737/gin-project-example/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()

	server := gin.New()

	server.Static("/static", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	apiRouter := server.Group("/api")
	{
		apiRouter.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.GetAll())
		})

		apiRouter.POST("/videos", func(ctx *gin.Context) {
			if err := videoController.Save(ctx); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video input is valid!"})
			}
		})
	}

	viewRouter := server.Group("/view")
	{
		viewRouter.GET("/video", videoController.ShowAll)
	}

	server.Run(":8088")
}
