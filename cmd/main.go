package main

import (
	"log"
	"os"

	"github.com/404th/hex_arch/internal/adapters/app/api"
	"github.com/404th/hex_arch/internal/adapters/core/arithmetic"
	"github.com/404th/hex_arch/internal/adapters/f/left/grpc"
	"github.com/404th/hex_arch/internal/adapters/f/right/db"
)

func main() {

	var nd *db.DB
	var ar *arithmetic.Adapter
	var napi *api.Adapter
	var gRPC *grpc.Adapter

	// db config
	dName := os.Getenv("DB_NAME")
	dSource := os.Getenv("DB_SOURCE")
	port := os.Getenv("PORT")

	ar = arithmetic.NewAdapter()
	nd, err := db.NewDB(dName, dSource)
	if err != nil {
		log.Fatalf("Failed to connect db: %v", err)
		return
	}

	napi = api.NewAdapter(nd, ar)
	gRPC = grpc.NewAdapter(napi)

	if err := gRPC.Run(port); err != nil {
		log.Fatalf("Failed to run gRPC Server: %v", err)
		return
	}
}
