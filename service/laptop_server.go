package service

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/MohamedNazir/TheCompleteGRPC/pb/github.com/MohamedNazir/TheCompleteGRPC/proto/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LaptopServer struct {
	Store LaptopStore
}

func NewLaptopServer(store LaptopStore) *LaptopServer {
	return &LaptopServer{store}
}

func (server *LaptopServer) CreateLaptop(ctx context.Context, req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("receive a requst to create a laptop with id %s \n", laptop.Id)
	if len(laptop.Id) > 0 {
		// check if it is avalid uuid
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not a valid UUID : %v ", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Error(codes.Internal, "Error generating uuid ")
		}
		laptop.Id = id.String()
	}

	// START :
	// Lets assume, the server takes time to process the request, to test the testcase with timeout in the context
	time.Sleep(6 * time.Second)
	// END:

	if err := contextError(ctx); err != nil {
		return nil, err
	}

	// save the laptop to the DB, in this case inmomory
	err := server.Store.Save(laptop)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "cannot save laptop to database")
	}
	log.Printf("Laptop saved successfully with id : %s \n", laptop.Id)

	res := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}
	return res, nil
}

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return logError(status.Error(codes.Canceled, "request is canceled"))
	case context.DeadlineExceeded:
		return logError(status.Error(codes.DeadlineExceeded, "deadline is exceeded"))
	default:
		return nil
	}
}

func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}

//func (server *LaptopServer) SearchLaptop(req *pb.SerachLaptopRequest, stream *pb.LaptopService_SearchLaptopServer)

func (server *LaptopServer) SearchLaptop(req *pb.SerachLaptopRequest, stream pb.LaptopService_SearchLaptopServer) error {
	filter := req.GetFilter()
	log.Printf("received a search laptop request with filter :%w", filter)

	err := server.Store.Search(filter,
		func(laptop *pb.Laptop) error {
			res := &pb.SearchLaptopResponse{
				Laptop: laptop,
			}
			err := stream.Send(res)
			if err != nil {
				return err
			}
			log.Printf("sent laptop with id %s :", laptop.Id)
			return nil
		})

	if err != nil {
		return status.Errorf(codes.Internal, "unexpected Error :%v", err)
	}
	return nil
}
