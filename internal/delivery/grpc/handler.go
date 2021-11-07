/*
	Copyright © 2021 Durudex

	This file is part of Durudex: you can redistribute it and/or modify
	it under the terms of the GNU Affero General Public License as
	published by the Free Software Foundation, either version 3 of the
	License, or (at your option) any later version.

	Durudex is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
	GNU Affero General Public License for more details.

	You should have received a copy of the GNU Affero General Public License
	along with Durudex. If not, see <https://www.gnu.org/licenses/>.
*/

package grpc

import (
	"context"

	pb "github.com/Durudex/durudex-notif-service/internal/delivery/grpc/protobuf"
	"github.com/Durudex/durudex-notif-service/internal/service"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type Handler struct {
	service *service.Service
}

// Creating a new grpc handler.
func NewGRPCHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

// Regisctration services handlers.
func (h *Handler) RegisterHandlers(srv *grpc.Server) {
	pb.RegisterNotifServiceServer(srv, NewEmailHandler(h.service))
}

// Unary grpc interceptor.
func (h *Handler) UnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Info().Msgf("unary interceptor: %s", info.FullMethod)
	return handler(ctx, req)
}
