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

package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/durudex/durudex-email-service/internal/config"
	"github.com/durudex/durudex-email-service/internal/delivery/grpc"
	"github.com/durudex/durudex-email-service/internal/server"
	"github.com/durudex/durudex-email-service/internal/service"
	"github.com/durudex/durudex-email-service/pkg/email"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// A function that running the application.
func Run() {
	// Set logger mode.
	if os.Getenv("DEBUG") == "true" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	// Initialize config.
	cfg, err := config.Init()
	if err != nil {
		log.Error().Err(err).Msg("error initialize config")
	}

	// Create a new email client.
	emailClient, err := email.NewClient(&email.SMTPConfig{
		Host:           cfg.SMTP.Host,
		Port:           cfg.SMTP.Port,
		Username:       cfg.SMTP.Username,
		Password:       cfg.SMTP.Password,
		ConnectTimeout: cfg.SMTP.ConnectTimeout,
		SendTimeout:    cfg.SMTP.SendTimeout,
		Helo:           cfg.SMTP.Helo,
		KeepAlive:      cfg.SMTP.KeepAlive,
	})
	if err != nil {
		log.Error().Err(err).Msg("error creating email client")
	}

	// Creating a new service and handlers.
	service := service.NewService(emailClient, cfg)
	grpcHandler := grpc.NewHandler(service)

	// Creating a new server.
	srv, err := server.NewServer(cfg, grpcHandler)
	if err != nil {
		log.Fatal().Err(err).Msg("error creating a new server")
	}

	// Run server.
	go srv.Run()

	// Quit in application.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	// Stoping server.
	srv.Stop()

	log.Info().Msg("Durudex Email Service stoping!")
}
