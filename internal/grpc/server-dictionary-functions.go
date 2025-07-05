package grpc

import (
	"context"
	"github.com/Lykeion-org/go-shared/pkg/grpc/generated/lexora"
	"github.com/Lykeion-org/go-shared/pkg/grpc/generated/shared"
)

func (s *kosmosServer) GetReferent(ctx context.Context, req *lexora.GetReferentRequest) (*shared.Referent, error){
	return s.languageHandler.LanguageSvc.GetReferent(ctx, req)
}
func (s *kosmosServer) GetSymbol(ctx context.Context, req *lexora.GetSymbolRequest) (*shared.Symbol, error){
	return s.languageHandler.LanguageSvc.GetSymbol(ctx, req)
}
func (s *kosmosServer) GetWord(ctx context.Context, req *lexora.GetWordRequest) (*shared.Word, error){
	return s.languageHandler.LanguageSvc.GetWord(ctx, req)
}
func (s *kosmosServer) ListReferents(ctx context.Context, req *lexora.ListReferentsRequest) (*lexora.ListReferentsResponse, error){
	return s.languageHandler.LanguageSvc.ListReferents(ctx, req)
}
func (s *kosmosServer) FindReferents(ctx context.Context, req *lexora.FindReferentsRequest) (*lexora.FindReferentsResponse, error){
	return s.languageHandler.LanguageSvc.FindReferents(ctx, req)
}
func (s *kosmosServer) CreateReferent(ctx context.Context, req *lexora.CreateReferentRequest) (*shared.Referent, error){
	return s.languageHandler.LanguageSvc.CreateReferent(ctx, req)
}
func (s *kosmosServer) CreateSymbol(ctx context.Context, req *lexora.CreateSymbolRequest) (*shared.Symbol, error){
	return s.languageHandler.LanguageSvc.CreateSymbol(ctx, req)
}
func (s *kosmosServer) CreateWord(ctx context.Context, req *lexora.CreateWordRequest) (*shared.Word, error){
	return s.languageHandler.LanguageSvc.CreateWord(ctx, req)
}
func (s *kosmosServer) CreateWordAttributes(ctx context.Context, req *lexora.CreateWordAttributesRequest) (*shared.WordAttributes, error){
	return s.languageHandler.LanguageSvc.CreateWordAttributes(ctx, req)
}
func (s *kosmosServer) UpdateReferent(ctx context.Context, req *lexora.UpdateReferentRequest) (*shared.Referent, error){
	return s.languageHandler.LanguageSvc.UpdateReferent(ctx, req)
}
func (s *kosmosServer) UpdateSymbol(ctx context.Context, req *lexora.UpdateSymbolRequest) (*shared.Symbol, error){
	return s.languageHandler.LanguageSvc.UpdateSymbol(ctx, req)
}
func (s *kosmosServer) UpdateWord(ctx context.Context, req *lexora.UpdateWordRequest) (*shared.Word, error){
	return s.languageHandler.LanguageSvc.UpdateWord(ctx, req)
}
func (s *kosmosServer) UpdateWordAttributes(ctx context.Context, req *lexora.UpdateWordAttributesRequest) (*shared.WordAttributes, error){
	return s.languageHandler.LanguageSvc.UpdateWordAttributes(ctx, req)
}
func (s *kosmosServer) DeleteReferent(ctx context.Context, req *lexora.DeleteReferentRequest) (*shared.Referent, error){
	return s.languageHandler.LanguageSvc.DeleteReferent(ctx, req)
}
func (s *kosmosServer) DeleteSymbol(ctx context.Context, req *lexora.DeleteSymbolRequest) (*shared.Symbol, error){
	return s.languageHandler.LanguageSvc.DeleteSymbol(ctx, req)
}
func (s *kosmosServer) DeleteWord(ctx context.Context, req *lexora.DeleteWordRequest) (*shared.Word, error){
	return s.languageHandler.LanguageSvc.DeleteWord(ctx, req)
}
func (s *kosmosServer) DeleteWordAttributes(ctx context.Context, req *lexora.DeleteWordAttributesRequest) (*shared.WordAttributes, error){
	return s.languageHandler.LanguageSvc.DeleteWordAttributes(ctx, req)
}
func (s *kosmosServer) LinkSymbolToReferent(ctx context.Context, req *lexora.LinkSymbolToReferentRequest) (*lexora.LinkSymbolToReferentResponse, error){
	return s.languageHandler.LanguageSvc.LinkSymbolToReferent(ctx, req)
}
func (s *kosmosServer) LinkWordToSymbol(ctx context.Context, req *lexora.LinkWordToSymbolRequest) (*lexora.LinkWordToSymbolResponse, error){
	return s.languageHandler.LanguageSvc.LinkWordToSymbol(ctx, req)
}
func (s *kosmosServer) SetSymbolLemma(ctx context.Context, req *lexora.SetSymbolLemmaRequest) (*lexora.SetSymbolLemmaResponse, error){
	return s.languageHandler.LanguageSvc.SetSymbolLemma(ctx, req)
}
func (s *kosmosServer) UnlinkSymbolFromReferent(ctx context.Context, req *lexora.UnlinkSymbolFromReferentRequest) (*lexora.UnlinkSymbolFromReferentResponse, error){
	return s.languageHandler.LanguageSvc.UnlinkSymbolFromReferent(ctx, req)
}
func (s *kosmosServer) UnlinkWordFromSymbol(ctx context.Context, req *lexora.UnlinkWordFromSymbolRequest) (*lexora.UnlinkWordFromSymbolResponse, error){
	return s.languageHandler.LanguageSvc.UnlinkWordFromSymbol(ctx, req)
}
func (s *kosmosServer) ListAllLinkedReferents(ctx context.Context, req *lexora.ListAllLinkedReferentsRequest) (*lexora.ListAllLinkedReferentsResponse, error){
	return s.languageHandler.LanguageSvc.ListAllLinkedReferents(ctx, req)
}
func (s *kosmosServer) ListAllLinkedSymbols(ctx context.Context, req *lexora.ListAllLinkedSymbolsRequest) (*lexora.ListAllLinkedSymbolsResponse, error){
	return s.languageHandler.LanguageSvc.ListAllLinkedSymbols(ctx, req)
}