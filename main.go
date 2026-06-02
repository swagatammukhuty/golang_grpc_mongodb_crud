package main

import (
	"grpc_go/db"
	"grpc_go/server"
	"log"
	"net"

	pb "grpc_go/grpc_mongo_crud/proto"

	"google.golang.org/grpc"
)

func main() {
	// Database connection
	err := db.InitMongo()
	if err != nil {
		log.Fatal("Mongo DB Connection Fail")
	}
	lis, err := net.Listen("tcp", "50051")
	if err != nil {
		log.Fatal("Failed to listen ", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server.UserService{})
	if err := grpcServer.Serve(lis); err!=nil{
		log.Fatal("Failed to server the gRPC server", err)
	}
	log.Println("GRPC Server is running on the port 50051")
}
