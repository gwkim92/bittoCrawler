package main

import (
	conf "bittoCralwer/config"
	"bittoCralwer/ether/api"
	ether "bittoCralwer/ether/proto"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

var (
	configFlag = flag.String("config", "./config/config.toml", "toml file to use for configuration")
	httpFlag   = flag.Int("http", 0, "router http port")
)

func main() {
	config := conf.NewConfig(*configFlag)

	// http
	if *httpFlag != 0 {
		config.Port.Http = *httpFlag
	}

	go startScrapingBlocks(config) // Starting the block scraping in a separate goroutine

	// gRPC server setup (unchanged from your code)
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		// TODO log 수정
		fmt.Println("err : ", err)
	}

	grpcServer := grpc.NewServer()
	ether.RegisterBlockServiceServer(grpcServer, &api.BlockServer{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}

	// TODO health check
}

func startScrapingBlocks(config *conf.Config) {
	server := &api.BlockServer{
		Config: config,
	}

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			blockResponse, err := server.ImportLatestBlock(context.Background(), &ether.ImportBlockRequest{})
			if err != nil {
				fmt.Println("error : ", err)
			} else {
				fmt.Println("block number : ", blockResponse.BlockNumber, blockResponse.BlockHash, blockResponse.Status)
			}
		}
	}
}
