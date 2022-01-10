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

package emailservice

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/durudex/durudex-email-service/internal/config"
	"github.com/durudex/durudex-email-service/internal/delivery/grpc"
	"github.com/durudex/durudex-email-service/internal/server"
	"github.com/durudex/durudex-email-service/internal/service"
	"github.com/durudex/durudex-email-service/pkg/email"

	"github.com/rs/zerolog/log"
)

// A function that running the application.
func Run(configPath string) {
	// Initialize config.
	cfg, err := config.Init(configPath)
	if err != nil {
		log.Error().Msgf("error initialize config: %s", err.Error())
	}

	// Create a new email client.
	emailClient, err := createEmailClient(cfg.SMTP)
	if err != nil {
		log.Error().Msgf("error creating email client: %s", err.Error())
	}

	// Creating a new service and handlers.
	service := service.NewService(emailClient, cfg)
	grpcHandler := grpc.NewHandler(service)

	// Creating a new server.
	srv, err := server.NewServer(cfg, grpcHandler)
	if err != nil {
		log.Fatal().Msgf("error creating a new server: %s", err.Error())
	}

	// Run server.
	go srv.Run()

	// Quit in application.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	// Stoping server.
	srv.Stop()
}

// Creating a new email client.
func createEmailClient(cfg config.SMTPConfig) (*email.Client, error) {
	return email.NewClient(&email.SMTPConfig{
		Host:           cfg.Host,
		Port:           cfg.Port,
		Username:       cfg.Username,
		Password:       cfg.Password,
		ConnectTimeout: cfg.ConnectTimeout,
		SendTimeout:    cfg.SendTimeout,
		Helo:           cfg.Helo,
		KeepAlive:      cfg.KeepAlive,
	})
}
