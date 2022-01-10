/*
 * Copyright © 2021-2022 Durudex

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

package grpc

import (
	"context"

	"github.com/durudex/durudex-email-service/internal/delivery/grpc/pb"
	"github.com/durudex/durudex-email-service/internal/delivery/grpc/pb/types"
	"github.com/durudex/durudex-email-service/internal/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Email handler structure.
type EmailHandler struct {
	service service.Email
	pb.UnimplementedEmailServiceServer
}

// Creating a new gRPC email handler.
func NewEmailHandler(service service.Email) *EmailHandler {
	return &EmailHandler{service: service}
}

// Send to user email verification code.
func (h *EmailHandler) UserCode(ctx context.Context, input *pb.UserCodeRequest) (*types.Status, error) {
	emailStatus, err := h.service.UserCode(input.Email, input.Username, input.Code)
	if err != nil {
		return &types.Status{Status: emailStatus}, status.Error(codes.Internal, err.Error())
	}

	return &types.Status{Status: emailStatus}, nil
}

// Send to user email logged in information.
func (h *EmailHandler) UserLoggedIn(ctx context.Context, input *pb.UserLoggedInRequest) (*types.Status, error) {
	emailStatus, err := h.service.UserLoggedIn(input.Email, input.Ip)
	if err != nil {
		return &types.Status{Status: emailStatus}, status.Error(codes.Internal, err.Error())
	}

	return &types.Status{Status: emailStatus}, nil
}

// Send to user email register information.
func (h *EmailHandler) UserRegister(ctx context.Context, input *pb.UserRegisterRequest) (*types.Status, error) {
	emailStatus, err := h.service.UserRegister(input.Email, input.Username)
	if err != nil {
		return &types.Status{Status: emailStatus}, status.Error(codes.Internal, err.Error())
	}

	return &types.Status{Status: false}, nil
}
