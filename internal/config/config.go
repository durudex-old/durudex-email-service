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
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type (
	// Config variables.
	Config struct {
		Server ServerConfig // Server config variables.
		SMTP   SMTPConfig   // SMTP config variables.
		Email  EmailConfig  // Email config variables.
	}

	// Server config variables.
	ServerConfig struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
		TLS  bool   `mapstructure:"tls"`
	}

	// SMTP config variables.
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

	// Parsing specified when starting the config file.
	if err := parseConfigFile(configPath); err != nil {
		return nil, err
	}

	var cfg Config

	// Unmarshal config keys.
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	// Set configurations from environment.
	setFromEnv(&cfg)

	return &cfg, nil
}

// Parsing specified when starting the config file.
func parseConfigFile(configPath string) error {
	log.Debug().Msgf("Parsing config file: %s", configPath)

	// Split path to folder and file.
	path := strings.Split(configPath, "/")

	viper.AddConfigPath(path[0]) // Folder.
	viper.SetConfigName(path[1]) // File.

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
	if err := viper.UnmarshalKey("email.template", &cfg.Email.Template); err != nil {
		return err
	}
	// Unmarshal server keys.
	return viper.UnmarshalKey("server", &cfg.Server)
}

// Set configurations from environment.
func setFromEnv(cfg *Config) {
	log.Debug().Msg("Set configurations from environment...")

	// SMTP configurations.
	cfg.SMTP.Username = os.Getenv("SMTP_USERNAME")
	cfg.SMTP.Password = os.Getenv("SMTP_PASSWORD")
}

// Populate defaults config variables.
func populateDefaults() {
	log.Debug().Msg("Populate defaults config variables...")

	// Server defaults.
	viper.SetDefault("server.host", defaultServerHost)
	viper.SetDefault("server.port", defaultServerPort)
	viper.SetDefault("server.tls", defaultServerTLS)

	// SMTP defaults.
	viper.SetDefault("smtp.host", defaultSMTPHost)
	viper.SetDefault("smtp.port", defaultServerPort)
	viper.SetDefault("smtp.connectTimeout", defaultSMTPConnectTimeout)
	viper.SetDefault("smtp.sendTimeout", defaultSMTPSendTimeout)
	viper.SetDefault("smtp.helo", defaultSMTPHelo)
	viper.SetDefault("smtp.keepAlive", defaultSMTPKeepAlive)

	// Email templates defaults.
	viper.SetDefault("email.template.verification", defaultEmailTemplateVerification)
	viper.SetDefault("email.template.loggedIn", defaultEmailTemplateLoggedIn)
	viper.SetDefault("email.template.register", defaultEmailTemplateRegister)

}
