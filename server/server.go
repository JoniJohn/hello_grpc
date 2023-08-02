package server

import (
	pb "hello_grpc/proto/buffs/proto"
)

type server struct {
	pb.MovieServiceServer
}
