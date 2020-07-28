package main

import (
	"../proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"time"
)

func main()  {
	fmt.Println("In Server ..")

	lis, err := net.Listen("tcp","0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Error failed to load server %v",err)
	}

	s := grpc.NewServer()

//	proto.RegisterCustomerServiceServer(s,)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Error failed to Serve")
	}
}

//type CustomerRequest struct {
//	cid int64
//}
//type Customer struct {
//	Name string
//	id int64
//	Orders int64
//}
//type CustomerResponse struct {
//	cust Customer
//}

type server struct {

}
func (*server) GetDetails (c context.Context, req *proto.CustomerRequest)(*proto.CustomerResponse, error) {
	fmt.Println("Function called ..")
	cid := req.GetCid()
	names := make([]string, 0)
	names = append(names, "Manish", "Vysakh", "Vishnu", "Sathya")
	rand.Seed(time.Now().Unix())
	name := names[rand.Intn(len(names))]
	orders := rand.Int63n(1000)

	res := &proto.CustomerResponse{
		Cust: &proto.Customer{
			Id: cid,
			Orders: orders,
			Name: name,
		},
	}
	return res, nil
}