package main

import (
	"log"
	"net"

	mailbox "github.com/hamidOyeyiola/gprc-mail-service/mail_service"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to start server on port 9000: %v", err)
	}

	m := mailbox.MailServer{}
	grpcServer := grpc.NewServer()

	mailbox.RegisterMailServiceServer(grpcServer, &m)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve on port 9000:%v", err)
	}
}
