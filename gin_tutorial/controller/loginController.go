package controller

import (
	"gin/dto"
	"gin/service"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(c *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jwtService   service.JWTService
}

func NewLoginContoller(loginService service.LoginService, jwtService service.JWTService) LoginController {
	return &loginController{
		loginService,
		jwtService,
	}
}

func (ct *loginController) Login(c *gin.Context) string {
	var credentials dto.Credentials
	err := c.ShouldBind(&credentials)
	if err != nil {
		return ""
	}
	isAuthenticated := ct.loginService.Login(credentials.Username, credentials.Password)
	if isAuthenticated {
		return ct.jwtService.GenerateToken(credentials.Username, true)
	}
	return ""
}
