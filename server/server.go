package server

import (
	"context"
	"grpc_go/db"
	pb "grpc_go/grpc_mongo_crud/proto"
	"grpc_go/models"

	"go.mongodb.org/mongo-driver/bson"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

// Create User
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

// Get User
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	user := models.User{}
	err := db.UserCollection.FindOne(ctx, bson.M{"user_id": req.UserId}).Decode(&user)
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

// Update User
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	filter := bson.M{"user_id": req.UserId}
	update := bson.M{"$set": bson.M{
		"name":     req.Name,
		"age":      req.Age,
		"password": req.Password,
		"email":    req.Email,
	}}
	_, err := db.UserCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{
		UserId: req.UserId,
		Name:   req.Name,
		Age:    req.Age,
		Email:  req.Email,
	}, nil
}

// Delete User
func (s *UserService) DeleteUser(ctx context.Context, req *pb.GetUserRequest) (*pb.DeleteResponse, error) {
	_, err := db.UserCollection.DeleteOne(ctx, bson.M{"user_id": req.UserId})
	if err != nil {
		return nil, err
	}
	return &pb.DeleteResponse{Message: "User Delete Successful"}, nil
}

// Get User List
// cursor is need while returning multiple document
func (s *UserService) GetUserList(ctx context.Context, req *pb.GetUserListRequest) (*pb.UserListResponse, error) {
	cursor, err := db.UserCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var users []*pb.UserResponse
	for cursor.Next(ctx) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, &pb.UserResponse{
			UserId: user.UserId,
			Name:   user.Name,
			Age:    user.Age,
			Email:  user.Email,
		})
	}
	return &pb.UserListResponse{
		Users: users,
	}, nil
}
