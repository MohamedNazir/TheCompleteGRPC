package service_test

import (
	"context"
	"io"
	"net"
	"testing"

	"github.com/MohamedNazir/TheCompleteGRPC/pb/github.com/MohamedNazir/TheCompleteGRPC/proto/pb"
	"github.com/MohamedNazir/TheCompleteGRPC/sample"
	"github.com/MohamedNazir/TheCompleteGRPC/serializer"
	"github.com/MohamedNazir/TheCompleteGRPC/service"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()
	laptopServer, serverAddress := startTestLaptopServer(t, service.NewInMemLaptopStore())
	laptopClient := NewTestLaptopClient(t, serverAddress)

	laptop := sample.NewLaptop()
	expectedId := laptop.Id
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	res, err := laptopClient.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, expectedId, res.Id)
	// check thta the  lapto is stored in DB
	other, err := laptopServer.Store.FindById(res.Id)
	require.NoError(t, err)
	require.NotNil(t, other)

	// check whether the same lapto is received.
	requireSameLaptop(t, laptop, other)
}

func startTestLaptopServer(t *testing.T, store service.LaptopStore) (*service.LaptopServer, string) {

	laptopServer := service.NewLaptopServer(store)
	grpcServer := grpc.NewServer()

	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	listener, err := net.Listen("tcp", ":0") // :0 is use any available port
	require.NoError(t, err)

	go grpcServer.Serve(listener) //it is a blocking call, so run it in a seperate goroutine

	return laptopServer, listener.Addr().String()
}

func NewTestLaptopClient(t *testing.T, serverAddress string) pb.LaptopServiceClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	require.NoError(t, err)
	return pb.NewLaptopServiceClient(conn)
}

func requireSameLaptop(t *testing.T, laptop1 *pb.Laptop, laptop2 *pb.Laptop) {
	//require.Equal(t, laptop1, laptop2) // this will fail, because we canot compare 2 objects, because grpc proto buf will gen extra fields

	json1, err := serializer.ProtoBufToJson(laptop1)
	require.NoError(t, err)

	json2, err := serializer.ProtoBufToJson(laptop2)
	require.NoError(t, err)

	require.Equal(t, json1, json2)
}

//UnitTest for server streaming
func TestClientSearchLaptop(t *testing.T) {
	t.Parallel()
	t.Log("Starting")
	filter := &pb.Filter{
		MaxPriceUsd: 2000,
		MinCpuCores: 4,
		MinCpuGhz:   2.2,
		MinRam:      &pb.Memory{Value: 8, Unit: pb.Memory_GIGABYTE},
	}

	store := service.NewInMemLaptopStore()
	expectedIDs := make(map[string]bool)

	for i := 0; i <= 6; i++ {
		laptop := sample.NewLaptop()

		switch i {
		case 0:
			laptop.PriceUsd = 2500
		case 1:
			laptop.Cpu.NumberCores = 2
		case 2:
			laptop.Cpu.MinGhz = 2.0
		case 3:
			laptop.Ram = &pb.Memory{Value: 4096, Unit: pb.Memory_MEGABYTE}
		case 4:
			laptop.PriceUsd = 1999
			laptop.Cpu.NumberCores = 5
			laptop.Cpu.MinGhz = 2.5
			laptop.Ram = &pb.Memory{Value: 16, Unit: pb.Memory_GIGABYTE}
			expectedIDs[laptop.Id] = true
		case 5:
			laptop.PriceUsd = 2000
			laptop.Cpu.NumberCores = 6
			laptop.Cpu.MinGhz = 2.8
			laptop.Ram = &pb.Memory{Value: 64, Unit: pb.Memory_GIGABYTE}
			expectedIDs[laptop.Id] = true
		}

		err := store.Save(laptop)
		require.NoError(t, err)
	}
	_, serverAddress := startTestLaptopServer(t, store)
	laptopClient := NewTestLaptopClient(t, serverAddress)

	req := &pb.SerachLaptopRequest{Filter: filter}
	stream, err := laptopClient.SearchLaptop(context.Background(), req)
	require.NoError(t, err)
	found := 0

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		require.NoError(t, err)
		require.Contains(t, expectedIDs, res.GetLaptop().GetId())
		found = found + 1

	}
	t.Log("found is :", found)
	require.Equal(t, len(expectedIDs), found)
}
