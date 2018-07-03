package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"

	pb "github.com/frimik/akc/api"
)

func main() {
	port := flag.Int("port", 8080, "port to listen on.")
	flag.Parse()

	// setup logging
	//logger, err := syslog.New(syslog.LOG_ALERT|syslog.LOG_USER, "akc")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.SetOutput(logger)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("could not listen on port %d: %v", *port, err)
	}
	log.Infof("Listening on port: %d", *port)

	s := grpc.NewServer()
	pb.RegisterAuthorizedKeysServer(s, server{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("could not serve: %v", err)
	}
}

type server struct{}

func (server) Auth(ctx context.Context, user *pb.User) (*pb.Keys, error) {
	p := pb.Keys{
		Keys: []*pb.Keys_Key{
			{
				Key: "foobar baz",
			},
			{
				Key: "foobar baz2",
			},
		},
	}

	log.Printf("Key message created for .proto structure: %v", p)
	return &p, nil
}
