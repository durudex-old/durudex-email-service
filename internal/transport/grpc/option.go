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

package grpc

import (
	"context"

	"github.com/durudex/durudex-email-service/internal/config"
	"github.com/durudex/durudex-email-service/pkg/tls"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Getting gRPC server options.
func getOptions(cfg config.TLSConfig) []grpc.ServerOption {
	log.Debug().Msg("Getting gRPC server options...")

	var opts []grpc.ServerOption

	// Added basic server options.
	opts = append(opts,
		// Unary interceptor.
		grpc.UnaryInterceptor(unaryInterceptor),
		// Stream interceptor.
		grpc.StreamInterceptor(streamInterceptor),
	)

	if cfg.Enable {
		creds, err := tls.LoadTLSConfig(cfg.CACert, cfg.Cert, cfg.Key)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to load TLS credentials")
		}

		// Append server credential options.
		opts = append(opts, grpc.Creds(credentials.NewTLS(creds)))
	}

	return opts
}

// Unary gRPC server interceptor.
func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Info().Str("method", info.FullMethod).Msg("Unary interceptor")

	return handler(ctx, req)
}

// Stream gRPC server interceptor.
func streamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Info().Str("method", info.FullMethod).Msg("Stream interceptor")

	return handler(srv, ss)
}
