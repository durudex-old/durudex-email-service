package v1

import (
	"github.com/durudex/durudex-email-service/internal/service"
	v1 "github.com/durudex/durudex-email-service/pkg/pb/v1"

	"google.golang.org/grpc"
)

// gRPC handler structure.
type Handler struct{ service *service.Service }

// Creating a new gRPC handler.
func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

// Registering gRPC handlers.
func (h *Handler) RegisterHandlers(srv *grpc.Server) {
	v1.RegisterEmailServiceServer(srv, NewEmailHandler(h.service))
}
