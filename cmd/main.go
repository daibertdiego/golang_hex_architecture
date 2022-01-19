package main

import (
	"log"
	"time"

	"github.com/daibertdiego/golang_hex_architecture/src/adapters/app/api"
	"github.com/daibertdiego/golang_hex_architecture/src/adapters/core/arithmetic"
	gRPC "github.com/daibertdiego/golang_hex_architecture/src/adapters/framework/primary/grpc"
	db "github.com/daibertdiego/golang_hex_architecture/src/adapters/framework/secondary"
	"github.com/daibertdiego/golang_hex_architecture/src/ports"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("dev.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	// ports
	var dbAdapter ports.DbPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var grpcAdapter ports.GRPCPort

	dbAdapter, err := db.NewAdapter(10 * time.Second)
	if err != nil {
		log.Fatal(err)
	}

	defer dbAdapter.CloseDbConnection()

	core = arithmetic.NewAdapter()

	appAdapter = api.NewAdapter(core, dbAdapter)

	grpcAdapter = gRPC.NewAdapter(appAdapter)

	log.Println("App running on port 9000...")

	grpcAdapter.Run()
}
