package monitoring

import(
	lyk_grpc "github.com/Lykeion-org/go-shared/pkg/grpc"
)

type watchdog struct {
	services []lyk_grpc.Client
}


func NewWatchdog(serviceTarget ...lyk_grpc.Client) *watchdog {
	return &watchdog{
		services: serviceTarget,
	}
}

func(w *watchdog) RegisterService(service lyk_grpc.Client) {
	w.services = append(w.services, service)
}

func (w *watchdog) Start() error {
	go w.CheckHealth()
	return nil
}

func (w *watchdog) CheckHealth() error {

}