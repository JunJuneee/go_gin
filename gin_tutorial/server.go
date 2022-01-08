package main

import (
	"gin/controller"
	"gin/midwares"
	"gin/service"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	// gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService service.VideoService = service.New()
	loginService service.LoginService = service.NewLoginService()
	jwtService   service.JWTService   = service.NewJWTService()

	videoController controller.VideoController = controller.New(videoService)
	loginContoller  controller.LoginController = controller.NewLoginContoller(loginService, jwtService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()
	server := gin.New()
	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), midwares.Logger())

	server.POST("/login", func(c *gin.Context) {
		token := loginContoller.Login(c)
		if token != "" {
			c.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			c.JSON(http.StatusUnauthorized, nil)
		}
	})

	apiRoutes := server.Group("/api", midwares.AuthorizeJWT())
	{
		apiRoutes.GET("/videos", func(c *gin.Context) {
			c.JSON(http.StatusOK, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(c *gin.Context) {
			err := videoController.Save(c)

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{"message": "video Input is Valid!"})
			}

		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	server.Run(":8080")
}
