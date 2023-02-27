package controller

import (
	"github.com/gin-gonic/gin"
	"go_practice/gorm_src/helper"
	"go_practice/gorm_src/models"
	"go_practice/gorm_src/services"
	"net/http"
)

type ContactController struct {
	cs *services.ContactService
}

func NewContactController(cs *services.ContactService) *ContactController {
	return &ContactController{cs: cs}
}

func (cc *ContactController) CreateContact(c *gin.Context) {
	var contact models.Contact

	err := c.ShouldBindJSON(&contact)

	if err != nil {
		resp := helper.GenerateValidationResponse(err)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	code := http.StatusOK

	resp := cc.cs.CreateContact(&contact)

	if !resp.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, resp)
}
