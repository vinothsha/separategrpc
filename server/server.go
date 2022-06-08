package main

import (
	"fmt"
	"net"
	pb "separategrpc/createproto"
	rb "separategrpc/readproto"

	"github.com/gocql/gocql"
	"google.golang.org/grpc"
)

type Server struct {
	pb.BlogCreateServiceServer
	rb.ReadAllBlogServiceServer
}

var addr = "0.0.0.0:50051"
var Session *gocql.Session

func main() {
	var err error
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "blog"
	cluster.Consistency = gocql.LocalOne
	Session, err = cluster.CreateSession()
	if err != nil {
		fmt.Println("error while connect to cassandra")
	}
	fmt.Println("cassandra connected")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("error while listen port 50051")
	}
	s := grpc.NewServer()
	pb.RegisterBlogCreateServiceServer(s, &Server{})
	rb.RegisterReadAllBlogServiceServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		fmt.Println("error while serve ther server")
	}
	fmt.Println("port 50051 is running")
}
