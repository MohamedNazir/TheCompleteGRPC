package service

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log"

	"github.com/MohamedNazir/TheCompleteGRPC/pb/github.com/MohamedNazir/TheCompleteGRPC/proto/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const MaxImageSize = 1 << 20

type LaptopServer struct {
	laptopStore LaptopStore
	imageaStore ImageStore
}

func NewLaptopServer(laptopStore LaptopStore, imageStore ImageStore) *LaptopServer {
	return &LaptopServer{laptopStore: laptopStore, imageaStore: imageStore}
}

func (server *LaptopServer) CreateLaptop(ctx context.Context, req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("receive a requst to create a laptop with id %s \n", laptop.Id)
	if len(laptop.Id) > 0 {
		// check if it is avalid uuid
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, logError(status.Errorf(codes.InvalidArgument, "laptop ID is not a valid UUID : %v ", err))
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, logError(status.Error(codes.Internal, "Error generating uuid "))
		}
		laptop.Id = id.String()
	}

	// START :
	// Lets assume, the server takes time to process the request, to test the testcase with timeout in the context
	//	time.Sleep(6 * time.Second)
	// END:

	if err := contextError(ctx); err != nil {
		return nil, err
	}

	// save the laptop to the DB, in this case inmomory
	err := server.laptopStore.Save(laptop)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}
		return nil, logError(status.Errorf(code, "cannot save laptop to database"))
	}
	log.Printf("Laptop saved successfully with id : %s \n", laptop.Id)

	res := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}
	return res, nil
}

//func (server *LaptopServer) SearchLaptop(req *pb.SerachLaptopRequest, stream *pb.LaptopService_SearchLaptopServer)

func (server *LaptopServer) SearchLaptop(req *pb.SerachLaptopRequest, stream pb.LaptopService_SearchLaptopServer) error {
	filter := req.GetFilter()
	log.Printf("received a search laptop request with filter :%v", filter)

	err := server.laptopStore.Search(stream.Context(), filter,
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
		return logError(status.Errorf(codes.Internal, "unexpected Error :%v", err))
	}
	return nil
}

func (server *LaptopServer) UploadImage(stream pb.LaptopService_UploadImageServer) error {

	req, err := stream.Recv()
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "cannot receive image info"))
	}
	laptopID := req.GetInfo().GetLaptopId()
	imgType := req.GetInfo().GetImageType()
	log.Printf("receive an upload image request for laptop %s with image type %s", laptopID, imgType)

	laptop, err := server.laptopStore.FindById(laptopID)

	if err != nil {
		return logError(status.Errorf(codes.Internal, "cannot find laptop %v", err))
	}
	if laptop == nil {
		return logError(status.Errorf(codes.NotFound, "laptop %s dosen't exist", laptopID))
	}

	imageData := bytes.Buffer{}
	imageSize := 0

	for {
		if err := contextError(stream.Context()); err != nil {
			return err
		}
		log.Print("waiting for chunk data")

		rq, err := stream.Recv()
		if err == io.EOF {
			log.Print("No More data")
			break
		}
		if err != nil {
			return logError(status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err))
		}
		chunk := rq.GetChunkData()
		size := len(chunk)
		imageSize += size
		if imageSize > MaxImageSize {
			return logError(status.Errorf(codes.InvalidArgument, "Too big file, allowed size is 1GB"))
		}

		//START
		//	time.Sleep(10 * time.Second)
		//END
		_, err = imageData.Write(chunk)
		if err != nil {
			return logError(status.Errorf(codes.Internal, "Internal server error %v", err))
		}
	}
	imageId, err := server.imageaStore.Save(laptopID, imgType, imageData)

	if err != nil {
		return logError(status.Errorf(codes.Internal, "Internal server error %v", err))
	}

	res := &pb.UploadImageResponse{Id: imageId, Size: uint64(imageSize)}

	err = stream.SendAndClose(res)
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "Unknown error %v", err))
	}

	log.Printf("The image is saved successfully with id %s", imageId)
	return nil
}

func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
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
