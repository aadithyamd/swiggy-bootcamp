package main

import (
	"./client"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	api := r.Group("/api")
	api.GET("/",Homepage)
	api.POST("/customer",PostCustomer)

	r.Run("localhost:5656")

}

func Homepage (c *gin.Context) {
	c.JSON(200,gin.H{
		"message":"Welcome to homepage",
	})
}


func PostCustomer(c *gin.Context) {

	cid := c.Request.FormValue("cid")
	fmt.Println(cid)

	name, orders := client.INIT(cid)

	c.JSON(http.StatusOK, gin.H{
		"id":cid,
		"name": name, //res.Cust.Name,
		"orders": orders, //res.Cust.Orders,
	})
}


