package services

import (
	"context"
	"fmt"

	pro "github.com/VasenthD/Sample-crud/proto" //import the modfile name and the folder name of proto
	"go.mongodb.org/mongo-driver/mongo"
)

type Rpc struct {
	pro.UnimplementedCustomerServiceServer
	coll *mongo.Collection
	ctx  context.Context
}

func InitCustomerService(coll *mongo.Collection, ctx context.Context) *Rpc {
	return &Rpc{
		UnimplementedCustomerServiceServer: pro.UnimplementedCustomerServiceServer{},
		coll:                               coll,
		ctx:                                ctx,
	}
}

func (r *Rpc) CreateCustomer(ctx context.Context, req *pro.CustomerInfo) (*pro.DBResponse, error) {

	
	_, err := r.coll.InsertOne(r.ctx,req)
	if err!=nil{
		fmt.Println("error in insetone 🥶")
		return nil,err
	}
	var outputs *pro.DBResponse
	outputs.SsnNumber=req.SsnNumber
	outputs.PhoneNumber=req.PhoneNumber
	outputs.Email=req.Email

	return outputs, nil

}
