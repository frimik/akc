package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"

	pb "github.com/frimik/akc/api"
)

func main() {
	backend := flag.String("backend", "localhost:8080", "address of the akc backend.")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Printf("Usage: akc [OPTION]... USER\n")
		os.Exit(1)
	}

	// setup logging
	//logger, err := syslog.New(syslog.LOG_ALERT|syslog.LOG_USER, "akc")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.SetOutput(logger)

	user := flag.Arg(0)
	fmt.Printf("User argument is: %s\n", user)
	fmt.Printf("Backend is: %s\n", *backend)

	conn, err := grpc.Dial(*backend, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to %s: %v", *backend, err)
	}
	defer conn.Close()

	client := pb.NewAuthorizedKeysClient(conn)

	req := &pb.AuthorizedKeysRequest{Username: user}
	res, err := client.Match(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not auth %s: %v", req.Username, err)
	}

	for _, key := range res.Keys {
		fmt.Printf("%s\n", key)
	}
}
