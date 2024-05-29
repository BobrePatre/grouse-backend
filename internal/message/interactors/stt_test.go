package interactors

import (
	"context"
	"errors"
	mocks2 "github.com/BobrePatre/grouse-backend/internal/infrastructure/logging/mocks"
	"github.com/BobrePatre/grouse-backend/internal/message/interactors/mocks"
	"go.uber.org/mock/gomock"
	"testing"

	"github.com/BobrePatre/grouse-backend/internal/message/dto"
	"github.com/stretchr/testify/assert"
	"log/slog"
)

func TestSttInteractor_SendNotification(t *testing.T) {

	type fields struct {
		notifications *mocks.MockNotificationGateway
		logger        *slog.Logger
	}

	type args struct {
		ctx context.Context
		dto dto.SendNotificationRequest
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
				dto: dto.SendNotificationRequest{
					ReciverId: 123,
					Content:   "Test message",
				},
			},
			wantErr: false,
			on: func(f *fields) {
				f.notifications.EXPECT().Notify(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
		},
		{
			name: "failed notification",
			args: args{
				ctx: context.Background(),
				dto: dto.SendNotificationRequest{
					ReciverId: 123,
					Content:   "Test message",
				},
			},
			wantErr: true,
			on: func(f *fields) {
				f.notifications.EXPECT().Notify(gomock.Any(), gomock.Any()).Return(errors.New("notification failed")).Times(1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Mock controller
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := &fields{
				notifications: mocks.NewMockNotificationGateway(ctrl),
				logger:        slog.New(mocks2.NewMockHandler(ctrl)),
			}

			oms := NewSttInteractor(f.notifications, f.logger)

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
