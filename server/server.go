package server

import (
	"context"
	"grpc_go/db"
	pb "grpc_go/grpc_mongo_crud/proto"
	"grpc_go/models"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	user := models.User{
		UserId:   req.UserId,
		Name:     req.Name,
		Age:      req.Age,
		Email:    req.Email,
		Password: req.Password,
	}
	_, err := db.UserCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{
		UserId: user.UserId,
		Name:   user.Name,
		Age:    user.Age,
		Email:  user.Email,
	}, nil

}
