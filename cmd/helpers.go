package main

import (
	"log/slog"
	"github.com/Lykeion-org/kosmos/internal/grpc"
)

type Services struct {
	SecurityService *grpc.AegisService
	UserService *grpc.DemesService
	LanguageService *grpc.LexoraService

}

func initializeServices() *Services {
	demes := grpc.NewDemes(Configuration.UserServiceTarget)
	if demes == nil {
		slog.Info("Starting kosmos without user service")
	}
	
	lexora := grpc.NewLexora(Configuration.LanguageServiceTarget)
	if lexora == nil {
		slog.Info("Starting kosmos without language service")
	}

	aegis := grpc.NewAegis(Configuration.SecurityServiceTarget) 

	return &Services{
		UserService: demes,
		LanguageService: lexora,
		SecurityService: aegis,
	}

}

