package controller

import (
	"gin/entity"
	"gin/service"
	"gin/validators"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(c *gin.Context) error
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func New(s service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &controller{
		service: s,
	}
}
func (ct *controller) FindAll() []entity.Video {
	return ct.service.FindAll()
}
func (ct *controller) Save(c *gin.Context) error {
	var video entity.Video
	err := c.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	ct.service.Save(video)
	return nil
}
