package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/MohamedNazir/TheCompleteGRPC/pb/github.com/MohamedNazir/TheCompleteGRPC/proto/pb"
	"github.com/MohamedNazir/TheCompleteGRPC/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("Hello Gopeher, client!")

	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dial server %s", *serverAddress)

	con, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server :", err)
	}

	laptopClient := pb.NewLaptopServiceClient(con)

	laptop := sample.NewLaptop()
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	// set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := laptopClient.CreateLaptop(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			//not a big deal
			log.Printf("laptop alrady exists")
		} else {
			log.Fatal("canot create Laptop :", err)
		}
		return
	}
	log.Printf("Laptop created with Id : %s", res.Id)
}

func createtRandomLaptop(client pb.LaptopServiceClient) {

}
