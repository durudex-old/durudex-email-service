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

package server

import (
	"context"

	"github.com/durudex/durudex-email-service/internal/config"
	"github.com/durudex/durudex-email-service/pkg/tls"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// gRPC server structure.
type gRPCServer struct{ Server *grpc.Server }

// Creating a new gRPC server.
func NewGRPC(cfg *config.TLSConfig) (*gRPCServer, error) {
	serverOptions := []grpc.ServerOption{}

	// Added basic server options.
	serverOptions = append(serverOptions, grpc.UnaryInterceptor(unaryInterceptor))

	if cfg.Enable {
		// Loading TLS credentials.
		creds, err := tls.LoadTLSCredentials(cfg.CACert, cfg.Cert, cfg.Key)
		if err != nil {
			return nil, err
		}

		// Append server options.
		// Append server credential options.
		serverOptions = append(serverOptions, grpc.Creds(creds))
	}

	return &gRPCServer{Server: grpc.NewServer(serverOptions...)}, nil
}

// Unary gRPC interceptor.
func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Info().Msgf("unary interceptor: %s", info.FullMethod)

	return handler(ctx, req)
}
