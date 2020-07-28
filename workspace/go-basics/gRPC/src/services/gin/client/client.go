package main

import (
	"context"
	"../proto"
	//	"./client"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strconv"
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

	fmt.Println("On gRPC client")
	conn, err := grpc.Dial("localhost:50051",grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Cannot talk with gRPC server %v", err)
	}
	defer conn.Close()
	c1 := proto.NewCustomerServiceClient(conn)
	id, err := strconv.ParseInt(cid,10,64)
	if err != nil {
		log.Fatalf("cid not found! %v", err)
	}
	fmt.Println(id)

	name, orders := getCust(c1)
//	name, orders := client.INIT(cid)

	c.JSON(http.StatusOK, gin.H{
		"id":cid,
		"name": name, //res.Cust.Name,
		"orders": orders, //res.Cust.Orders,
	})
}


func getCust(c proto.CustomerServiceClient) (string, int64) {
	req:= &proto.CustomerRequest{
		Cid: 6,
	}

	res, err := c.GetDetails(context.Background(),req)

	if err!= nil {
		log.Fatalf("Error While Getting customer details : %v ", err)
	}

	log.Println("Response From the gRPC server ", res)

	return res.Cust.Name, res.Cust.Orders
}