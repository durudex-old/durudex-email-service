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
func (h *EmailHandler) SendEmailUserCode(ctx context.Context, input *pb.SendEmailUserCodeRequest) (*pb.SendEmailUserCodeResponse, error) {
	err := h.service.UserCode(input.Email, input.Username, input.Code)
	if err != nil {
		return &pb.SendEmailUserCodeResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &pb.SendEmailUserCodeResponse{}, nil
}

// Send to user email logged in information.
func (h *EmailHandler) SendEmailUserLoggedIn(ctx context.Context, input *pb.SendEmailUserLoggedInRequest) (*pb.SendEmailUserLoggedInResponse, error) {
	err := h.service.UserLoggedIn(input.Email, input.Ip)
	if err != nil {
		return &pb.SendEmailUserLoggedInResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &pb.SendEmailUserLoggedInResponse{}, nil
}

// Send to user email register information.
func (h *EmailHandler) SendEmailUserRegister(ctx context.Context, input *pb.SendEmailUserRegisterRequest) (*pb.SendEmailUserRegisterResponse, error) {
	err := h.service.UserRegister(input.Email, input.Username)
	if err != nil {
		return &pb.SendEmailUserRegisterResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &pb.SendEmailUserRegisterResponse{}, nil
}
