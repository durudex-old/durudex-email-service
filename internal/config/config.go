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
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
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
	}
)

// Initialize config.
func Init(configPath string) *Config {
	log.Debug().Msg("Initialize config...")

	// Parsing config file.
	parseConfigFile(configPath)

	var cfg Config
	// Unmarshal config keys.
	unmarshal(&cfg)

	// Set env configurations.
	setFromEnv(&cfg)

	return &cfg
}

// Parsing config file.
func parseConfigFile(configPath string) {
	log.Debug().Msgf("Parsing config file: %s", configPath)

	// Split path to folder and file.
	path := strings.Split(configPath, "/")

	viper.AddConfigPath(path[0]) // folder
	viper.SetConfigName(path[1]) // file

	// Read config file.
	if err := viper.ReadInConfig(); err != nil {
		log.Error().Msgf("error parsing config file: %s", err.Error())
	}
}

// Unmarshal config keys.
func unmarshal(cfg *Config) {
	log.Debug().Msg("Unmarshal config keys...")

	// Unmarshal grpc keys.
	if err := viper.UnmarshalKey("grpc", &cfg.GRPC); err != nil {
		log.Error().Msgf("error unmarshal grpc keys: %s", err.Error())
	}
	// Unmarshal smpt keys.
	if err := viper.UnmarshalKey("smtp", &cfg.SMTP); err != nil {
		log.Error().Msgf("error unmarshal smtp keys: %s", err.Error())
	}
	// Unmarshal email keys.
	if err := viper.UnmarshalKey("email.template", &cfg.Email.Template); err != nil {
		log.Error().Msgf("error unmarshal email template keys: %s", err.Error())
	}
}

// Seting environment variables from .env file.
func setFromEnv(cfg *Config) {
	cfg.SMTP.Username = os.Getenv("SMTP_USERNAME")
	cfg.SMTP.Password = os.Getenv("SMTP_PASSWORD")
}
