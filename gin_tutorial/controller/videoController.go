package controller

import (
	"gin/entity"
	"gin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(c *gin.Context) error
	Update(c *gin.Context) error
	Delete(c *gin.Context) error
	ShowAll(c *gin.Context)
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func New(s service.VideoService) VideoController {
	validate = validator.New()
	// validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
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

func (ct *controller) Update(c *gin.Context) error {
	var video entity.Video
	err := c.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	video.ID = id
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	ct.service.Update(video)
	return nil
}

func (ct *controller) Delete(c *gin.Context) error {
	var video entity.Video

	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	video.ID = id

	ct.service.Delete(video)
	return nil

}

func (ct *controller) ShowAll(c *gin.Context) {
	videos := ct.service.FindAll()
	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}
	c.HTML(http.StatusOK, "index.html", data)
}
