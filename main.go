package main

import (
	"github.com/AvinFajarF/handlers"
	"github.com/AvinFajarF/initializers"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.Conntection()
}

func main(){
	router := gin.Default()

	v1 := router.Group("/api/v1")

	v1.POST("/register", handlers.Register)
	v1.POST("/login", handlers.Login)

	router.Run(":8081")
}