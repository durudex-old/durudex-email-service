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

package config

import (
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Default config path.
const defaultConfigPath string = "configs/main"

type (
	// Config variables.
	Config struct {
		GRPC  GRPCConfig
		SMTP  SMTPConfig
		Email EmailConfig
	}

	// gRPC config variables.
	GRPCConfig struct {
		Host string    `mapstructure:"host"`
		Port string    `mapstructure:"port"`
		TLS  TLSConfig `mapstructure:"tls"`
	}

	// TLS config variables.
	TLSConfig struct {
		Enable bool   `mapstructure:"enable"`
		CACert string `mapstructure:"ca-cert"`
		Cert   string `mapstructure:"cert"`
		Key    string `mapstructure:"key"`
	}

	// SMTP config variables.
	SMTPConfig struct {
		ConnectTimeout time.Duration `mapstructure:"connect-timeout"`
		SendTimeout    time.Duration `mapstructure:"send-timeout"`
		Helo           string        `mapstructure:"helo"`
		KeepAlive      bool          `mapstructure:"keep-alive"`
		Host           string
		Port           int
		Username       string
		Password       string
	}

	// Email config variables.
	EmailConfig struct {
		Template EmailTemplate `mapstructure:"template"`
	}

	// Email templates config variables.
	EmailTemplate struct {
		Verification string `mapstructure:"verification"`
		LoggedIn     string `mapstructure:"logged-in"`
		Register     string `mapstructure:"register"`
	}
)

// Initialize config.
func Init() (*Config, error) {
	log.Debug().Msg("Initialize config...")

	// Parsing specified when starting the config file.
	if err := parseConfigFile(); err != nil {
		return nil, err
	}

	var cfg Config

	// Unmarshal config keys.
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	// Set configurations from environment.
	if err := setFromEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// Parsing specified when starting the config file.
func parseConfigFile() error {
	// Get config path variable.
	configPath := os.Getenv("CONFIG_PATH")

	// Check is config path variable empty.
	if configPath == "" {
		configPath = defaultConfigPath
	}

	log.Debug().Msgf("Parsing config file: %s", configPath)

	// Split path to dir and file.
	dir, file := filepath.Split(configPath)

	viper.AddConfigPath(dir)
	viper.SetConfigName(file)

	// Read in config file.
	return viper.ReadInConfig()
}

// Unmarshal config keys.
func unmarshal(cfg *Config) error {
	log.Debug().Msg("Unmarshal config keys...")

	// Unmarshal smtp keys.
	if err := viper.UnmarshalKey("smtp", &cfg.SMTP); err != nil {
		return err
	}
	// Unmarshal email template keys.
	if err := viper.UnmarshalKey("email", &cfg.Email); err != nil {
		return err
	}
	// Unmarshal gRPC server keys.
	return viper.UnmarshalKey("grpc", &cfg.GRPC)
}

// Set configurations from environment.
func setFromEnv(cfg *Config) error {
	log.Debug().Msg("Set configurations from environment...")

	// Get smtp port from environments.
	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return err
	}

	// SMTP configurations.
	cfg.SMTP.Host = os.Getenv("SMTP_HOST")
	cfg.SMTP.Port = smtpPort

	// Email configurations.
	cfg.SMTP.Username = os.Getenv("EMAIL_USERNAME")
	cfg.SMTP.Password = os.Getenv("EMAIL_PASSWORD")

	return nil
}
