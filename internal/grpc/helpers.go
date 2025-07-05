package grpc

import (
	"log/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func newClientConnection(target string) (*grpc.ClientConn, error){
	creds := insecure.NewCredentials()
	conn, err:= grpc.NewClient(target, grpc.WithTransportCredentials(creds))
	if err != nil {
		slog.Debug("Failed to open grpc client connection", "target", target, "error",err)
	}
	return conn, nil
}
