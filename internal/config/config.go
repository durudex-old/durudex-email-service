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

package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Default variables.
const (
	defaultGRPCPort = "8002"
	defaultGRPCHost = "notifservice.durudex.local"
	defaultGRPCTLS  = true
)

type (
	Config struct {
		GRPC  GRPCConfig
		SMTP  SMTPConfig
		Email EmailConfig
	}

	// CRPC server config variables.
	GRPCConfig struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
		TLS  bool   `mapstructure:"tls"`
	}

	// SMTP server config variables.
	SMTPConfig struct {
		Host           string        `mapstructure:"host"`
		Port           int           `mapstructure:"port"`
		ConnectTimeout time.Duration `mapstructure:"connectTimeout"`
		SendTimeout    time.Duration `mapstructure:"sendTimeout"`
		Helo           string        `mapstructure:"helo"`
		KeepAlive      bool          `mapstructure:"keepAlive"`
		Username       string
		Password       string
	}

	// Email config variables.
	EmailConfig struct {
		Template EmailTemplate
	}

	// Email templates config variables.
	EmailTemplate struct {
		Verification string `mapstructure:"verification"`
		LoggedIn     string `mapstructure:"loggedIn"`
		Register     string `mapstructure:"register"`
	}
)

// Initialize config.
func Init(configPath string) (*Config, error) {
	log.Debug().Msg("Initialize config...")

	// Populate defaults config variables.
	populateDefaults()

	// Parsing config file.
	if err := parseConfigFile(configPath); err != nil {
		return nil, fmt.Errorf("error parsing config file: %s", err.Error())
	}

	var cfg Config
	// Unmarshal config keys.
	if err := unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("error unmarshal config keys: %s", err.Error())
	}

	// Set env configurations.
	setFromEnv(&cfg)

	return &cfg, nil
}

// Parsing config file.
func parseConfigFile(configPath string) error {
	log.Debug().Msgf("Parsing config file: %s", configPath)

	// Split path to folder and file.
	path := strings.Split(configPath, "/")

	viper.AddConfigPath(path[0]) // folder
	viper.SetConfigName(path[1]) // file

	// Read config file.
	return viper.ReadInConfig()
}

// Unmarshal config keys.
func unmarshal(cfg *Config) error {
	log.Debug().Msg("Unmarshal config keys...")

	// Unmarshal grpc keys.
	if err := viper.UnmarshalKey("grpc", &cfg.GRPC); err != nil {
		return err
	}
	// Unmarshal smpt keys.
	if err := viper.UnmarshalKey("smtp", &cfg.SMTP); err != nil {
		return err
	}
	// Unmarshal email keys.
	return viper.UnmarshalKey("email.template", &cfg.Email.Template)
}

// Seting environment variables from .env file.
func setFromEnv(cfg *Config) {
	log.Debug().Msg("Set from environment configurations...")

	cfg.SMTP.Username = os.Getenv("SMTP_USERNAME")
	cfg.SMTP.Password = os.Getenv("SMTP_PASSWORD")
}

// Populate defaults config variables.
func populateDefaults() {
	log.Debug().Msg("Populate defaults config variables...")

	viper.SetDefault("grpc.host", defaultGRPCHost)
	viper.SetDefault("grpc.port", defaultGRPCPort)
	viper.SetDefault("grpc.tls", defaultGRPCTLS)
}
