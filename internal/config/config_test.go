/*
 * Copyright © 2021-2022 Durudex
 *
 * This file is part of Durudex: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * Durudex is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Durudex. If not, see <https://www.gnu.org/licenses/>.
 */

package config_test

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/durudex/durudex-email-service/internal/config"
)

// Testing initialize config.
func TestConfig_Init(t *testing.T) {
	// Environment configurations.
	type env struct {
		configPath    string
		smtpHost      string
		smtpPort      string
		emailUsername string
		emailPassword string
	}

	// Testing args.
	type args struct{ env env }

	// Set environment configuration.
	setEnv := func(env env) {
		os.Setenv("CONFIG_PATH", env.configPath)
		os.Setenv("SMTP_HOST", env.smtpHost)
		os.Setenv("SMTP_PORT", env.smtpPort)
		os.Setenv("EMAIL_USERNAME", env.emailUsername)
		os.Setenv("EMAIL_PASSWORD", env.emailPassword)
	}

	// Tests structures.
	tests := []struct {
		name    string
		args    args
		want    *config.Config
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				env: env{
					configPath:    "fixtures/main",
					smtpHost:      "mail.durudex.local",
					smtpPort:      "25",
					emailUsername: "user",
					emailPassword: "1234567890",
				},
			},
			want: &config.Config{
				GRPC: config.GRPCConfig{
					Host: "email.service.durudex.local",
					Port: "8002",
					TLS: config.TLSConfig{
						Enable: true,
						CACert: "./certs/rootCA.pem",
						Cert:   "./certs/email.service.durudex.local-cert.pem",
						Key:    "./certs/email.service.durudex.local-key.pem",
					},
				},
				SMTP: config.SMTPConfig{
					ConnectTimeout: time.Second * 10,
					SendTimeout:    time.Second * 10,
					Helo:           "durudex",
					KeepAlive:      true,
					Host:           "mail.durudex.local",
					Port:           25,
					Username:       "user",
					Password:       "1234567890",
				},
				Email: config.EmailConfig{
					Template: config.EmailTemplate{
						Verification: "./web/template/verification.html",
						LoggedIn:     "./web/template/logged_in.html",
						Register:     "./web/template/register.html",
					},
				},
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environment configuration.
			setEnv(tt.args.env)

			// Initialize config.
			got, err := config.Init()
			if (err != nil) != tt.wantErr {
				t.Errorf("error initialize config: %s", err.Error())
			}

			// Check for similarity of a config.
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("error config are not similar: %s", err.Error())
			}
		})
	}
}
