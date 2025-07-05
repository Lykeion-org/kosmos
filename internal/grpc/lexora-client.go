package grpc

import (
	"google.golang.org/grpc"
	"log/slog"
	"github.com/Lykeion-org/go-shared/pkg/grpc/generated/lexora"
)

type LexoraService struct {
	connection *grpc.ClientConn
	Client lexora.LexoraServiceClient
}

func NewLexora(target string) *LexoraService {
	conn, err := newClientConnection(target)
	if err != nil {
		slog.Warn("Failed to start Lexora service on target", "target", target, "error",err)
		return nil
	}
	client := lexora.NewLexoraServiceClient(conn)

	return &LexoraService{
		Client: client,
		connection: conn,
	}
}
func (c *LexoraService) StopClient() {
	c.connection.Close()
	c.connection = nil
}