package main

import (
	"log"
	"log/slog"
	"github.com/Lykeion-org/go-shared/pkg/config"
	"github.com/Lykeion-org/go-shared/pkg/helpers"
	"github.com/Lykeion-org/kosmos/internal/grpc"
	"github.com/Lykeion-org/kosmos/internal/handlers"
)


func main(){
	log.Println("Starting application...\nLoading config file")
	
	//Set variables and open grpc clients
	cfg, err := config.LoadConfigFile[Config]("../dev_config.yaml")
	if err != nil {
		slog.Error("Failed to load application configuration file", "error", err)
		panic("Could not load application configuration")
	}
	Configuration = cfg;
	helpers.InitializeStandardLogger("debug")

	slog.Info("Opening grpc clients...")
	
	services := initializeServices()
	defer services.UserService.StopClient() ; defer services.SecurityService.StopClient()

	slog.Info("Connecting to message queues...")

	//Create server and server handlers
	userHandler := handlers.NewUserHandler(services.SecurityService.Client, services.UserService.Client)
	languageHandler := handlers.NewLanguageHandler(services.LanguageService.Client, nil, nil)
	learningHandler := handlers.NewLearningHandler()
	server := grpc.NewKosmosServer(userHandler, languageHandler, learningHandler)
	
	//Start grpc server
	slog.Info("Starting grpc server", "port", cfg.KosmostServerAddress)
	server.StartServer(cfg.KosmostServerAddress)

	
	//Handle queues


	select{}


}
