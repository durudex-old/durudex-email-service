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
	"strconv"
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
		Template EmailTemplate
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

	// Populate defaults config variables.
	populateDefaults()

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

// Populate defaults config variables.
func populateDefaults() {
	log.Debug().Msg("Populate defaults config variables...")

	// Server defaults.
	viper.SetDefault("server.host", defaultServerHost)
	viper.SetDefault("server.port", defaultServerPort)
	viper.SetDefault("server.tls", defaultServerTLS)

	// SMTP defaults.
	viper.SetDefault("smtp.connect-timeout", defaultSMTPConnectTimeout)
	viper.SetDefault("smtp.send-timeout", defaultSMTPSendTimeout)
	viper.SetDefault("smtp.helo", defaultSMTPHelo)
	viper.SetDefault("smtp.keep-alive", defaultSMTPKeepAlive)

	// Email templates defaults.
	viper.SetDefault("email.template.verification", defaultEmailTemplateVerification)
	viper.SetDefault("email.template.logged-in", defaultEmailTemplateLoggedIn)
	viper.SetDefault("email.template.register", defaultEmailTemplateRegister)

}
