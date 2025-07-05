package handlers

import (
	"context"
	"errors"
	"log/slog"

	"github.com/Lykeion-org/go-shared/pkg/grpc/generated/aegis"
	"github.com/Lykeion-org/go-shared/pkg/grpc/generated/demes"
)

type Tokens struct {
}

type UserHandler interface {
	RegisterUser(ctx context.Context, email string, password string) ([2]string, error)
	LoginUser(ctx context.Context, email string, password string) ([2]string, error)
	GetUserClaims(ctx context.Context, token string) (*aegis.ValidateTokenResponse, error)
	GetUserInfo(ctx context.Context, uid string) (*demes.User, error)
	RequestPasswordReset(ctx context.Context, token string) error
	RequestEmailVerification(ctx context.Context, token string) error
	ConfirmPasswordReset(ctx context.Context, token string, newPassword string) error
	ConfirmEmailVerification(ctx context.Context, token string) error
}

type userHandler struct {
	authSvc aegis.AegisServiceClient
	userSvc demes.DemesServiceClient
}

func NewUserHandler(authSvc aegis.AegisServiceClient, userService demes.DemesServiceClient) *userHandler {
	return &userHandler{
		authSvc: authSvc,
		userSvc: userService,
	}
}

func (h *userHandler) RegisterUser(ctx context.Context, email string, password string) ([2]string, error) {
	// Create a new user in the database
	user, err := h.userSvc.CreateUser(ctx, &demes.CreateUserRequest{Email: email, Password: password})
	if err != nil {
		slog.Debug("Error receiving user response from user service", "error", err)
		return [2]string{}, err
	}

	// Generate a token for the user
	tokens, err := h.authSvc.GenerateTokens(ctx, &aegis.GenerateTokensRequest{UserUid: user.Uid, UserRole: user.Role})
	if err != nil {
		slog.Debug("Error receiving token response from auth service", "error", err)
		return [2]string{}, err
	}

	return [2]string{tokens.AccessToken, tokens.RefreshToken}, nil
}
func (h *userHandler) LoginUser(ctx context.Context, email string, password string) ([2]string, error) {
	res, err := h.userSvc.ValidatePassword(ctx, &demes.ValidatePasswordRequest{Email: email, Password: password})
	if err != nil {
		slog.Debug("Error validating user password", "error", err)
		return [2]string{}, err
	}

	if !res.Correct {
		slog.Debug("Error from user service, password might be incorrect", "error", err)
		return [2]string{}, errors.New("incorrect password")
	}

	user, err := h.userSvc.ReadUserByEmailAddress(ctx, &demes.UserMailRequest{EmailAddress: email})
	if err != nil {
		return [2]string{}, err
	}

	tokens, err := h.authSvc.GenerateTokens(ctx, &aegis.GenerateTokensRequest{UserUid: user.Uid, UserRole: user.Role})
	if err != nil {
		slog.Debug("Error receiving token response from auth service", "error", err)
		return [2]string{}, err
	}

	return [2]string{tokens.AccessToken, tokens.RefreshToken}, nil
}
func (h *userHandler) GetUserClaims(ctx context.Context, token string) (*aegis.ValidateTokenResponse, error) {
	claims, err := h.authSvc.ValidateToken(ctx, &aegis.ValidateTokenRequest{AccessToken: token})
	if err != nil {
		return nil, err
	}
	slog.Debug("claims response from aegis service", "claims", claims)

	return claims, nil

}
func (h *userHandler) GetUserInfo(ctx context.Context, userUid string) (*demes.User, error) {

	user, err := h.userSvc.ReadUser(ctx, &demes.UserUidRequest{Uid: userUid})
	if err != nil {
		return nil, err
	}

	return user, nil

}
func (h *userHandler) RequestPasswordReset(ctx context.Context, token string) error {
	return errors.New("not implemented yet")

}
func (h *userHandler) RequestEmailVerification(ctx context.Context, token string) error {
	return errors.New("not implemented yet")
}
func (h *userHandler) ConfirmPasswordReset(ctx context.Context, token string, newPassword string) error {
	return errors.New("not implemented yet")
}
func (h *userHandler) ConfirmEmailVerification(ctx context.Context, token string) error {
	return errors.New("not implemented yet")
}
