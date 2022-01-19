package primary

import (
	"log"
	"net"

	"github.com/daibertdiego/golang_hex_architecture/src/adapters/framework/primary/grpc/pb"
	"github.com/daibertdiego/golang_hex_architecture/src/ports"
	"google.golang.org/grpc"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

func (grpca Adapter) Run() {
	var err error
	listen, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatal("Failed to listen on port 9000: %v", err)
	}

	arithmeticServiceServer := grpca
	grpcServer := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(grpcServer, arithmeticServiceServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatal("Unable to server GRPC Server over port 9000: %v", err)
	}
}
