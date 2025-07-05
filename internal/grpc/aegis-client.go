package grpc

import (
	"google.golang.org/grpc"
	"log/slog"
	"github.com/Lykeion-org/go-shared/pkg/grpc/generated/aegis"
)

type AegisService struct {
	connection *grpc.ClientConn
	Client aegis.AegisServiceClient
}

func NewAegis(target string) *AegisService {
	conn, err := newClientConnection(target)
	if err != nil {
		slog.Warn("Failed to start Demes service on target", "target", target, "error",err)
		return nil
	}
	client := aegis.NewAegisServiceClient(conn)

	return &AegisService{
		Client: client,
		connection: conn,
	}
}
func (c *AegisService) StopClient() {
	c.connection.Close()
	c.connection = nil
}