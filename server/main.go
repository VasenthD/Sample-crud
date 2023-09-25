package main

import (
	"context"
	"fmt"
	"net"

	"github.com/VasenthD/Sample-crud/config"
	"github.com/VasenthD/Sample-crud/constants"
	pro "github.com/VasenthD/Sample-crud/proto" //import the modfile name and the folder name of proto
	"github.com/VasenthD/Sample-crud/services"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func initDatabase(client *mongo.Client) {
	infoleCollection := config.GetCollection(client, "college", "CustomerInfo")
	services.InitCustomerService(infoleCollection, context.Background())

}

func main(){
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	initDatabase(mongoclient)

	lis, err := net.Listen("tcp", constants.Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s, healthServer)
	pro.RegisterCustomerServiceServer(s, &services.Rpc{})
	fmt.Println("Server listening on", constants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}
