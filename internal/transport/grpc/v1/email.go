package v1

import (
	"context"

	"github.com/durudex/durudex-email-service/internal/service"
	v1 "github.com/durudex/durudex-email-service/pkg/pb/v1"
)

// Email gRPC server handler.
type EmailHandler struct {
	service service.Email
	v1.UnimplementedEmailServiceServer
}

// Creating a new email gRPC handler.
func NewEmailHandler(service service.Email) *EmailHandler {
	return &EmailHandler{service: service}
}

func (h *EmailHandler) SendEmailUserCode(ctx context.Context, input *v1.SendEmailUserCodeRequest) (*v1.SendEmailUserCodeResponse, error) {
	return &v1.SendEmailUserCodeResponse{}, nil
}

func (h *EmailHandler) SendEmailUserLoggedIn(ctx context.Context, input *v1.SendEmailUserLoggedInRequest) (*v1.SendEmailUserLoggedInResponse, error) {
	return &v1.SendEmailUserLoggedInResponse{}, nil
}

func (h *EmailHandler) SendEmailUserRegister(ctx context.Context, input *v1.SendEmailUserRegisterRequest) (*v1.SendEmailUserRegisterResponse, error) {
	return &v1.SendEmailUserRegisterResponse{}, nil
}
