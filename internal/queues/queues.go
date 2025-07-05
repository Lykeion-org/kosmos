package queues

import (
	"log/slog"

	"github.com/Lykeion-org/go-shared/pkg/mq"
)


type KosmosQueues interface {
}

type kosmosQueues struct {
	producerQueues *mq.Producer
	consumerQueues *mq.Consumer
}

func NewKosmosQueues() KosmosQueues{
	return &kosmosQueues{
	}
}

func connectToQueues(url string, producerQueueNames []string, consumerQueueNames []string) *kosmosQueues{
	conn, err := mq.NewConnection(url)
	if err != nil {
		slog.Error("Failed to create or connect to message queues", "error", err)
	}

	producers, err := mq.NewProducer(conn.Channel, producerQueueNames[0])
	if err != nil {
		slog.Error("Failed to create producer queue", "error", err)
	}

	consumers, err := mq.NewConsumer(conn.Channel, consumerQueueNames[0])
	if err != nil {
		slog.Error("Failed to consumer queue", "error", err)
	}

	return &kosmosQueues{
		producerQueues: producers,
		consumerQueues: consumers,
	}



}

