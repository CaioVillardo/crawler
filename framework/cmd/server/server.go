package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/reflection"
)

var db *gorm.DB

func main() {
	db = utils.ConnectDB()
	db.LogMode(true)

	port := flag.Int("port", 0, "Choose the server port")
	flag.Parse()
	log.Printf("Start server on port %d", *port)

	userServer := setUpUserServer()

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, userServer)
	reflection.Register(grpcServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal("Cannot start server: %v", err)
	}
	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("Cannot start server: %v", err)
	}
}

func setUpUserServer() *servers.UserServer {
	userRepository := repositories.UserRepositoryDb{Db: db}
	userServer := servers.NewUserServer()
	userServer.UserUseCase = usecases.UserUseCase{UserRepository: userRepository}
	return userServer
}
