package servers

import (
	"context"
	"log"
)

type UserServer struct {
	User        domain.User
	UserUseCase usecases.UserUseCase
}

func NewUserServer() *UserServer {
	return &UserServer{}
}

func GetUserId(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	id := ctx.Param("id")

	if err != nil {
		log.Fatalf("Error during the RPC Create User %v", err)
	}

	return &pb.UserResponse{
		Token: user.Token,
	}, nil

}
