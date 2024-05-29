package delivery

import (
	"context"
	infraGrpc "github.com/BobrePatre/grouse-backend/internal/infrastructure/grpc"
	"github.com/BobrePatre/grouse-backend/internal/message/dto"
	"github.com/BobrePatre/grouse-backend/internal/message/interactors"
	stt "github.com/BobrePatre/grouse-backend/pkg/api/grpc/golang"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
)

type SttHandlers struct {
	stt.UnimplementedSttServiceServer
	interactor *interactors.SttInteractor
	logger     *slog.Logger
}

func RegisterSttHandlers(srv *grpc.Server, srvCfg *infraGrpc.Config, gtw *runtime.ServeMux, logger *slog.Logger, interactor *interactors.SttInteractor) error {
	impl := &SttHandlers{
		interactor: interactor,
		logger:     logger,
	}
	ctx := context.Background()

	stt.RegisterSttServiceServer(srv, impl)
	err := stt.RegisterSttServiceHandlerFromEndpoint(ctx, gtw, srvCfg.Address(), []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		logger.Error("failed to register stt service handler", "error", err)
		return err
	}
	return nil
}

func (h *SttHandlers) NotifyInText(ctx context.Context, req *stt.NotifyInTextRq) (*emptypb.Empty, error) {
	err := h.interactor.SendNotification(ctx, dto.SendNotificationRequest{
		ReciverId: req.TelegramId,
		Content:   req.Body,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to send notification: %v", err)
	}

	return &emptypb.Empty{}, nil
}
