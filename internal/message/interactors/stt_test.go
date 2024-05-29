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
		assert  func(*testing.T, *fields, error)
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
			assert: func(t *testing.T, f *fields, err error) {
				if assert.NoError(t, err) {
					t.Log("Successfully assert no error")
				} else {
					t.Fatal("Failed to assert error")
				}
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
			assert: func(t *testing.T, f *fields, err error) {
				if assert.Error(t, err) && assert.Equal(t, "notification failed", err.Error()) {
					t.Log("Successfully assert error")
				} else {
					t.Fatal("Failed to assert error")
				}

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
			if tt.assert != nil {
				tt.assert(t, f, err)
			}
		})
	}
}
