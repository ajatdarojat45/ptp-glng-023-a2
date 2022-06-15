package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"c2/controllers"
	"c2/config"
)

func main(){
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}
	
	r := gin.Default()
	r.GET("/", HelloWorld)
	r.GET("/orders", inDB.GetOrders)
	r.POST("/orders", inDB.CreateOrder)
	r.PUT("/orders/:orderId", inDB.UpdateOrder)
	r.DELETE("/orders/:orderId", inDB.DeleteOrder)
	r.Run(":3000")
}

func HelloWorld (c *gin.Context) {
	var result gin.H
	result = gin.H{
		"message": "Hello world",
	}

	c.JSON(http.StatusOK, result)
}