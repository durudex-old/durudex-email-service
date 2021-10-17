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

package notifservice

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Durudex/durudex-notif-service/internal/config"
	"github.com/Durudex/durudex-notif-service/internal/delivery/grpc"
	"github.com/Durudex/durudex-notif-service/internal/server"
	"github.com/Durudex/durudex-notif-service/internal/service"
	"github.com/rs/zerolog/log"
)

// Run durudex auth service application.
func Run(configPath string) {
	// Initialize config.
	cfg := config.Init(configPath)

	// Service, Handler.
	service := service.NewService()
	grpcHandler := grpc.NewGRPCHandler(service)

	// Create and run server.
	srv := server.NewServer(cfg, grpcHandler)
	go func() {
		srv.Run()
	}()

	// Quit in application.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	// Stoping server.
	srv.Stop()

	log.Info().Msg("Durudex Notif Service stoping!")
}
