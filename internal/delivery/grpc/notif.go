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
	"github.com/Durudex/durudex-notif-service/internal/delivery/grpc/protobuf/types"
	"github.com/Durudex/durudex-notif-service/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type EmailHandler struct {
	service *service.Service
	pb.UnimplementedNotifServiceServer
}

// Creating a new auth handler.
func NewEmailHandler(service *service.Service) *EmailHandler {
	return &EmailHandler{
		service: service,
	}
}

// Sending a user code to the email.
func (h *EmailHandler) UserCode(ctx context.Context, input *pb.UserCodeRequest) (*types.Status, error) {
	// Send to user email code.
	emailStatus, err := h.service.Email.UserCode(input.Email, input.Name, input.Code)
	if err != nil {
		return &types.Status{Status: emailStatus}, status.Error(codes.Internal, err.Error())
	}

	return &types.Status{Status: emailStatus}, nil
}
