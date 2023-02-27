package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go_practice/gorm_src/controller"
	"go_practice/gorm_src/repositories"
	"go_practice/gorm_src/services"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	route := gin.Default()

	cr := repositories.NewContactRepository(db)
	cs := services.NewContactService(cr)
	cc := controller.NewContactController(cs)
	route.POST("/create", cc.CreateContact)

	return route
}
