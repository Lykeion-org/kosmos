package handlers

import (
	"errors"
	"context"
	"github.com/Lykeion-org/go-shared/pkg/grpc/generated/shared"

)

type FlashcardOptions struct {
	QuestionType int32
	HintType int32
	AnswerType int32
	QuestionSource int32
	HintSource int32
	AnswerSource int32
	TargetLanguage int32
	ReferenceLanguage int32
}

type LearningHandler interface {
	CreateLanguageFlashcards(ctx context.Context, userUid string, symbols []*shared.Symbol, flashcardOption FlashcardOptions) error
	GetUserFlashcards(ctx context.Context, userUid string, limit int32, offset int32) error
	GetScheduledFlashcards(ctx context.Context, userUid string, limit int32, offset int32) error
	CheckAnswer(ctx context.Context, userUid string, flashcardUid string, answer string) error
	}

type learningHandler struct {
}

func NewLearningHandler() LearningHandler {
	return &learningHandler{}
}

func (h *learningHandler) CreateLanguageFlashcards(ctx context.Context, userUid string, symbols []*shared.Symbol, flashcardOption FlashcardOptions) error {
	return errors.New("not implemented yet")
}
func(u *learningHandler) GetUserFlashcards(ctx context.Context, userUid string, limit int32, offset int32) error {
	return errors.New("not implemented yet")
}
func(u *learningHandler) GetScheduledFlashcards(ctx context.Context, userUid string, limit int32, offset int32) error {
	return errors.New("not implemented yet")
}
func(u *learningHandler) CheckAnswer(ctx context.Context, userUid string, flashcardUid string, answer string) error {
	return errors.New("not implemented yet")
}
