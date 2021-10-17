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

package server

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"net"

	"github.com/Durudex/durudex-notif-service/internal/config"
	handler "github.com/Durudex/durudex-notif-service/internal/delivery/grpc"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	CACertFile           = "cert/rootCA.pem"
	notifserviceCertFile = "cert/notifservice-cert.pem"
	notifserviceKeyFile  = "cert/notifservice-key.pem"
)

type Server struct {
	tcpServer   *net.Listener
	grpcServer  *grpc.Server
	grpcHandler *handler.Handler
}

// Creating a new server.
func NewServer(cfg *config.Config, grpcHandler *handler.Handler) *Server {
	serverOptions := []grpc.ServerOption{}

	// If TLS is true.
	if cfg.Server.TLS {
		tlsCredentials, err := LoadTLSCredentials()
		if err != nil {
			log.Fatal().Msgf("error load tls credentials: %s", err.Error())
		}

		// GRPC server options.
		serverOptions = append(
			serverOptions,
			grpc.Creds(tlsCredentials),
			grpc.UnaryInterceptor(grpcHandler.UnaryInterceptor),
		)
	}

	// Server address.
	address := cfg.Server.Host + ":" + cfg.Server.Port

	return &Server{
		tcpServer:   NewTCPServer(address),
		grpcServer:  grpc.NewServer(serverOptions...),
		grpcHandler: grpcHandler,
	}
}

// Loading TLS credentials.
func LoadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate od the CA who signed client's certificate.
	pemCA, err := ioutil.ReadFile(CACertFile)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemCA) {
		return nil, errors.New("error to add server CA's certificate")
	}

	// Load server's certificate and private key.
	serverCert, err := tls.LoadX509KeyPair(notifserviceCertFile, notifserviceKeyFile)
	if err != nil {
		return nil, err
	}

	// Create the credentials and returning it.
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	return credentials.NewTLS(config), nil
}

// Creating a new tpc server.
func NewTCPServer(address string) *net.Listener {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal().Msgf("error creating new tcp server: %s", err.Error())
		return nil
	}

	return &lis
}

// Run grpc server.
func (srv *Server) Run() {
	log.Debug().Msg("Running server...")

	// Regisctration grpc handlers.
	srv.grpcHandler.RegisterHandlers(srv.grpcServer)

	if err := srv.grpcServer.Serve(*srv.tcpServer); err != nil {
		log.Fatal().Msgf("error running grpc server: %s", err.Error())
	}
}

// Stop grpc server.
func (srv *Server) Stop() {
	log.Info().Msg("Stoping grpc server...")
	srv.grpcServer.Stop()
}
