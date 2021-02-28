package service_test

import (
	"context"
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
	laptopServer, serverAddress := startTestLaptopServer(t)
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

func startTestLaptopServer(t *testing.T) (*service.LaptopServer, string) {

	laptopServer := service.NewLaptopServer(service.NewInMemLaptopStore())
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
	//require.Equal(t, laptop1, laptop2) // this will fail, because we canot compate 2 objects, because grpc proto buf will gen extra fields

	json1, err := serializer.ProtoBufToJson(laptop1)
	require.NoError(t, err)

	json2, err := serializer.ProtoBufToJson(laptop2)
	require.NoError(t, err)

	require.Equal(t, json1, json2)
}
