package grpc

import (
	"context"
	"log/slog"
	"net"

	lyk_grpc "github.com/Lykeion-org/go-shared/pkg/grpc"
	kosmos "github.com/Lykeion-org/go-shared/pkg/grpc/generated/kosmos"
	"github.com/Lykeion-org/kosmos/internal/handlers"
	"google.golang.org/grpc"
)


type kosmosServer struct {
	kosmos.UnimplementedKosmosServiceServer
	userHandler handlers.UserHandler
	languageHandler *handlers.LanguageHandler
	learningHandler handlers.LearningHandler
	server *grpc.Server
}

func NewKosmosServer(uh handlers.UserHandler, lh *handlers.LanguageHandler, lrh handlers.LearningHandler) lyk_grpc.Server{
	return &kosmosServer{
		userHandler: uh,
		languageHandler: lh,
		learningHandler: lrh,
	}
}

func (s *kosmosServer) StartServer(target string) error{
	grpcServer := grpc.NewServer()
	kosmos.RegisterKosmosServiceServer(grpcServer, s)

	
	listener, err := net.Listen("tcp", target)
	if err != nil {
		slog.Error("Failed to open tcp listener for grpc server", "error", err)
		return err
	}

	go func(){
		err := grpcServer.Serve(listener)
		if err != nil {
			panic(err)
		}
	}()

	s.server = grpcServer
	return nil
}
func (s *kosmosServer) StopServer() error{
	s.server.GracefulStop()
	return nil
}

func(s *kosmosServer) RegisterUser(ctx context.Context, req *kosmos.RegisterUserRequest) (*kosmos.RegisterUserResponse,error) {
	slog.Debug("received a register user request")
	res, err := s.userHandler.RegisterUser(ctx, req.Email, req.Password)			//index 0: access_token, index 1: refresh_token
	if err != nil {
		return nil, err
	}

	return &kosmos.RegisterUserResponse{
		AccessToken: res[0],
		RefreshToken: res[1],
	}, nil
}

func (s *kosmosServer) LoginUser(ctx context.Context, req *kosmos.LoginUserRequest) (*kosmos.LoginUserResponse, error){
	slog.Debug("received a login user request")
	res, err := s.userHandler.LoginUser(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &kosmos.LoginUserResponse{
		AccessToken: res[0],
		RefreshToken: res[1],
	}, nil
}

func (s *kosmosServer) GetUserClaims(ctx context.Context, req *kosmos.RequestToken) (*kosmos.GetUserClaimsResponse, error){
	slog.Debug("received a get claim request", "token", req.AccessToken)
	res, err := s.userHandler.GetUserClaims(ctx, req.AccessToken)
	if err != nil {
		return nil, err
	}

	return &kosmos.GetUserClaimsResponse{
		UserId: res.UserUid,
		Role: res.UserRole,
		IsAuthenticaed: true,
		IsValid:true,
	}, nil
}

func (s *kosmosServer) GetUserInfo(ctx context.Context, req *kosmos.GetUserInfoRequest) (*kosmos.GetUserInfoResonse, error){
	slog.Debug("received a get user info request", "userId", req.UserUid)
	res, err := s.userHandler.GetUserInfo(ctx, req.UserUid)
	if err != nil{
		return nil, err
	}

	slog.Debug("received a user from demes service", "user", res)

	return &kosmos.GetUserInfoResonse{
		UserId				: res.Uid,
		EmailAddress		: res.EmailAddress,
		AvatarSource		: res.AvatarSource,
		NativeLanguage		: res.NativeLanguage,
		ReferenceLanguage	: res.ReferenceLanguage,
		TargetLanguage		: res.TargetLanguage,
		FirstName			: res.FirstName,
		LastName			: res.LastName,
		ExtraName			: res.ExtraName,
		DisplayName			: res.DisplayName,
		Role				: res.Role,
	}, nil
}

