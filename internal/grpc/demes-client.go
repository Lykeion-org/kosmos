package grpc

import (
	"google.golang.org/grpc"
	"log/slog"
	"github.com/Lykeion-org/go-shared/pkg/grpc/generated/demes"
)

type DemesService struct {
	connection *grpc.ClientConn
	Client demes.DemesServiceClient
}

func NewDemes(target string) *DemesService {
	conn, err := newClientConnection(target)
	if err != nil {
		slog.Warn("Failed to start Demes service on target", "target", target, "error",err)
		return nil
	}
	client := demes.NewDemesServiceClient(conn)

	return &DemesService{
		Client: client,
		connection: conn,
	}
}
func (c *DemesService) StopClient() {
	c.connection.Close()
	c.connection = nil
}