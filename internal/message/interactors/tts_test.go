package interactors

import (
	"context"
	"errors"
	"github.com/BobrePatre/grouse-backend/internal/message/interactors/mocks"
	"go.uber.org/mock/gomock"
	"os"
	"testing"

	"github.com/BobrePatre/grouse-backend/internal/message/dto"
	"github.com/stretchr/testify/assert"
	"log/slog"
)

func TestSendTtsNotification(t *testing.T) {

	type fields struct {
		notifications *mocks.MockNotificationTtsGateway
		logger        *slog.Logger
	}

	type args struct {
		ctx context.Context
		dto dto.SendTtsNotificationRequest
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
		on      func(*fields)
	}{
		{
			name: "successful notification",
			args: args{
				ctx: context.Background(),
				dto: dto.SendTtsNotificationRequest{
					UserId:  123,
					Content: "Test message",
				},
			},
			wantErr: false,
			on: func(f *fields) {
				f.notifications.EXPECT().NotifyClient(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
		},
		{
			name: "failed notification",
			args: args{
				ctx: context.Background(),
				dto: dto.SendTtsNotificationRequest{
					UserId:  123,
					Content: "Test message",
				},
			},
			wantErr: true,
			on: func(f *fields) {
				f.notifications.EXPECT().NotifyClient(gomock.Any(), gomock.Any()).Return(errors.New("notification failed")).Times(1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Mock controller
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := &fields{
				notifications: mocks.NewMockNotificationTtsGateway(ctrl),
				logger:        slog.New(slog.NewJSONHandler(os.Stdout, nil)),
			}

			oms := NewTtsInteractor(f.logger, f.notifications)

			if tt.on != nil {
				tt.on(f)
			}

			// actor
			err := oms.SendNotification(tt.args.ctx, tt.args.dto)

			// assert
			if tt.wantErr {
				if assert.Error(t, err) {
					t.Logf("expected error: %v", err)
				} else {
					t.Fatalf("expected an error but got none")
				}
			} else {
				if assert.NoError(t, err) {
					t.Log("notification sent successfully")
				} else {
					t.Fatalf("did not expect an error but got: %v", err)
				}
			}
		})
	}
}
