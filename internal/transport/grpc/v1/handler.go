/*
 * Copyright © 2022 Durudex

 * This file is part of Durudex: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.

 * Durudex is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.

 * You should have received a copy of the GNU Affero General Public License
 * along with Durudex. If not, see <https://www.gnu.org/licenses/>.
 */

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
