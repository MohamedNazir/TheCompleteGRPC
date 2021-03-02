package main

import (
	"context"
	"flag"
	"fmt"
	"io"
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

	for i := 0; i < 10; i++ {
		createtRandomLaptop(laptopClient)
	}

	filter := &pb.Filter{
		MaxPriceUsd: 3000,
		MinCpuCores: 4,
		MinCpuGhz:   2.5,
		MinRam: &pb.Memory{
			Value: 8,
			Unit:  pb.Memory_GIGABYTE},
	}

	SearchLaptop(laptopClient, filter)

}

func SearchLaptop(laptopClient pb.LaptopServiceClient, filter *pb.Filter) {
	log.Printf("search filer :", filter)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.SerachLaptopRequest{
		Filter: filter,
	}
	stream, err := laptopClient.SearchLaptop(ctx, req)
	if err != nil {
		log.Fatalf("Error in laptop search ", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("Error", err)
		}
		laptop := res.GetLaptop()
		log.Print("- found:", laptop.GetId())
		log.Print(" + brand:", laptop.GetBrand())
		log.Print(" + name:", laptop.GetName())
		log.Print(" + cpu cores:", laptop.GetCpu().GetNumberCores())
		log.Print(" + cpu  min ghz:", laptop.GetCpu().GetMinGhz())
		log.Print(" + ram:", laptop.GetRam().GetValue())
		log.Print(" + price:", laptop.GetPriceUsd())
	}
}

func createtRandomLaptop(laptopClient pb.LaptopServiceClient) {
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
