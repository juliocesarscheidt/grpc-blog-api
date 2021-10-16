package main

import (
	"os"
	"fmt"
	"log"
	"net"
	"time"
	"context"
	"os/signal"

	"github.com/juliocesarscheidt/blog/blogpb"
	"github.com/juliocesarscheidt/blog/utils"
	"github.com/juliocesarscheidt/blog/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// if the program crashes, it shows the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Server is Alive")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// init client and retrieve collection
	client, _ := utils.GetMongoClient(ctx)
	collection := utils.GetMongoCollection(client, "item")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)
	reflection.Register(server)

	s := &services.Server{
		Collection: collection,
	}
	blogpb.RegisterBlogServiceServer(server, s)

	// shutdown hook
	go func() {
		fmt.Println("Listening on 0.0.0.0:50051")
		if err := server.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// wait for control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// block until a signal is received
	<-ch

	fmt.Println("Stopping the Mongo Client")
	client.Disconnect(ctx)

	fmt.Println("Stopping the Program")
	server.Stop()

	fmt.Println("Closing the Listener")
	lis.Close()

	fmt.Println("End of Program")
}
