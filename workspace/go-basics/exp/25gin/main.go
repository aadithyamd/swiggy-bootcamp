package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

//import 	"net/http"
type Order struct {
	OrderId int `json:"order ID"`

}

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

//slice of orders
var orders =[]Order{
	{1},{2},{3},
}



func GetOrders (c *gin.Context) {
	c.JSON(200,&orders)
}


func PostOrder(c *gin.Context) {

	body := c.Request.Body

	content, err := ioutil.ReadAll(body)

	if err != nil {
		fmt.Println("Sorry!", err.Error())
	}

	fmt.Println(content)

	c.JSON(http.StatusCreated, gin.H {
		"message" : string{content},
	})
}

func main() {
	r := gin.Default()

	api := r.Group("/api")
	api.GET("/", HomePage)
	api.GET("/orders", GetOrders)
	api.POST("/orders", PostOrder)

	//r.GET("/ping", HomePage)

	r.Run("localhost:5656") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}


