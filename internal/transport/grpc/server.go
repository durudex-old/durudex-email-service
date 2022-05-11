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
	"net"

	"github.com/durudex/durudex-email-service/internal/config"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// gRPC server structure.
type Server struct {
	server  *grpc.Server
	config  config.GRPCConfig
	handler *Handler
}

// Creating a new gRPC server.
func NewServer(cfg config.GRPCConfig, handler *Handler) *Server {
	options := getOptions(cfg.TLS)

	return &Server{
		server:  grpc.NewServer(options...),
		config:  cfg,
		handler: handler,
	}
}

// Running gRPC server.
func (s *Server) Run() {
	log.Info().Msg("Running gRPC server...")

	address := s.config.Host + ":" + s.config.Port

	// Creating a new TCP listener.
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal().Err(err).Msg("error creating tcp listener")
	}

	// Registering gRPC handlers.
	s.handler.RegisterHandlers(s.server)

	// Running gRPC server.
	if err := s.server.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("error running gRPC server")
	}
}

// Stoping gRPC server.
func (s *Server) Stop() {
	log.Info().Msg("Stopping gRPC server...")

	s.server.Stop()
}
