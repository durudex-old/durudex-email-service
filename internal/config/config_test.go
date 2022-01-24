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

package config

import (
	"os"
	"reflect"
	"testing"
)

// Testing initialize config.
func TestConfig_Init(t *testing.T) {
	// Environment configurations.
	type env struct {
		smtpUsername string
		smtpPassword string
	}

	// Testing args.
	type args struct {
		path string
		env  env
	}

	// Set environment configuration.
	setEnv := func(env env) {
		os.Setenv("SMTP_USERNAME", env.smtpUsername)
		os.Setenv("SMTP_PASSWORD", env.smtpPassword)
	}

	// Tests structures.
	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				path: "fixtures/main",
				env: env{
					smtpUsername: "user",
					smtpPassword: "1234567890",
				},
			},
			want: &Config{
				Server: ServerConfig{
					Host: defaultServerHost,
					Port: defaultServerPort,
					TLS:  defaultServerTLS,
				},
				SMTP: SMTPConfig{
					Host:           defaultSMTPHost,
					Port:           defaultSMTPPort,
					ConnectTimeout: defaultSMTPConnectTimeout,
					SendTimeout:    defaultSMTPSendTimeout,
					Helo:           defaultSMTPHelo,
					KeepAlive:      defaultSMTPKeepAlive,
					Username:       "user",
					Password:       "1234567890",
				},
				Email: EmailConfig{
					Template: EmailTemplate{
						Verification: defaultEmailTemplateVerification,
						LoggedIn:     defaultEmailTemplateLoggedIn,
						Register:     defaultEmailTemplateRegister,
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
			got, err := Init(tt.args.path)
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
