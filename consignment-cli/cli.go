package main

import (
	"context"
	"encoding/json"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	pb "github.com/nithinchandra/shipping/consignment-service/proto/consignment"
	"io/ioutil"
	"log"
	"os"
)

const (
	address	=	"localhost:50051"
	defaultFilename =	"consignment.json"
)

func parseFile(file string) (*pb.Consignment, error){
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil{
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment,err
}

func main(){
	cmd.Init()


	client := pb.NewShippingServiceClient("consignment", microclient.DefaultClient)

	// Contact the server and print out its response.
	file := defaultFilename
	if len(os.Args) > 1{
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil{
		log.Fatalf("could not parse L %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil{
		log.Fatalf("could not greet : %v", err)
	}
	log.Printf("Created : %t", r.Created)

	getAll ,err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil{
		log.Fatalf("could not list consignments: %v", err)
	}

	for _,v := range getAll.Consignments{
		log.Println(v)
	}
}