package main

import (
	"../../proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

func main(){ //(cid string) (string, int64) {
	cid :="56"
	fmt.Println("On gRPC client")
	conn, err := grpc.Dial("localhost:50051",grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Cannot talk with gRPC server %v", err)
	}
	defer conn.Close()
	c := proto.NewCustomerServiceClient(conn)
	id, err := strconv.ParseInt(cid,10,64)
	if err != nil {
		log.Fatalf("cid not found! %v", err)
	}
	fmt.Println(id)

	//name, orders :=
		getCust(c)

	//fmt.Println(name,orders)
//	return name, orders
}

func getCust(c proto.CustomerServiceClient) {//(string, int64) {
	req:= &proto.CustomerRequest{
		Cid: 6,
	}

	res, err := c.GetDetails(context.Background(),req)

	if err!= nil {
		log.Fatalf("Error While Getting customer details : %v ", err)
	}

	log.Println("Response From the gRPC server ", res)

	//return res.Cust.Name, res.Cust.Orders
}
