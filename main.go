package main

import (
	"context"
	"flag"
	"fmt"
	"hello_grpc/models"
	pb "hello_grpc/proto/buffs/proto"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "gRPC server port")
)

type server struct {
	pb.MovieServiceServer
}

// TODO(lesley): move to controller
func (*server) GetMovie(ctx context.Context, req *pb.ReadMovieRequest) (*pb.ReadMovieResponse, error) {
	fmt.Println("Read Movie", req.GetId())

	var movie models.Movie

	// TODO(lesley): read from the DB instead
	movie.ID = "1"
	movie.Title = "Spiderman"
	movie.Genre = "Action"
	movie.CreateAt = time.Now()
	movie.UpdateAt = time.Now()

	return &pb.ReadMovieResponse{
		Movie: &pb.Movie{
			Id:    movie.ID,
			Title: movie.Title,
			Genre: movie.Genre,
		},
	}, nil

}

func main() {
	fmt.Println("gRPC Server running...")

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterMovieServiceServer(s, &server{})

	log.Printf("Server listening at Port: %v", listen.Addr())

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}

}
