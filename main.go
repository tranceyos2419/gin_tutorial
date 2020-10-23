package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yos/gin_tutorial_reviews/controller"
	"github.com/yos/gin_tutorial_reviews/service"
)

func main() {
	server := gin.Default()

	var (
		videoService   service.VideoService      = service.New() //?
		videoContoller controller.VideoContoller = controller.New(videoService)
	)

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoContoller.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoContoller.Save(ctx))
	})

	server.Run(":8000")
}
