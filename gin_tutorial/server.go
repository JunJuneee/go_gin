package main

import (
	"gin/controller"
	"gin/midwares"
	"gin/service"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
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
	server.Use(gin.Recovery(), midwares.Logger(), midwares.BasicAuth(), gindump.Dump())

	server.GET("/videos", func(c *gin.Context) {
		c.JSON(http.StatusOK, videoController.FindAll())
	})

	server.POST("/videos", func(c *gin.Context) {
		err := videoController.Save(c)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "video Input is Valid!"})
		}

	})
	server.Run(":8080")
}
