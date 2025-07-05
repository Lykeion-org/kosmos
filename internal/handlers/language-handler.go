package handlers

import (
	"github.com/Lykeion-org/go-shared/pkg/grpc/generated/lexora"
	"github.com/Lykeion-org/go-shared/pkg/mq"
)


type LanguageHandler struct {
	LanguageSvc lexora.LexoraServiceClient
}

func NewLanguageHandler(langSvc lexora.LexoraServiceClient, producerQueue *mq.Producer, consumerQueue *mq.Consumer) *LanguageHandler {
	return &LanguageHandler{LanguageSvc: langSvc}
}




