package main

import (
	"fmt"
	"log"
	"net/http"

	pro "github.com/VasenthD/Sample-crud/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	r := gin.Default()
	conn, err := grpc.Dial("localhost:3030", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	
	client := pro.NewCustomerServiceClient(conn)

	r.POST("/createcustomer", func(c *gin.Context) {
		var request pro.CustomerInfo

		// Parse incoming JSON
		if err := c.ShouldBindJSON(&request); err != nil {
			
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		// Call the gRPC service
		response, err := client.CreateCustomer(c.Request.Context(), &request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			fmt.Println("im the error 😍 ")
			return
		}

		c.JSON(http.StatusOK, gin.H{"value": response})
	})
	r.Run(":8080")
}
