/*
 * Copyright © 2022 Durudex

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

import "time"

const (
	// Config defaults.
	defaultConfigPath string = "configs/main"

	// Server defaults.
	defaultServerHost string = "email-service.durudex.local"
	defaultServerPort string = "8002"
	defaultServerTLS  bool   = true

	// SMTP defaults.
	defaultSMTPHost           string        = "email.durudex.local"
	defaultSMTPPort           int           = 25
	defaultSMTPConnectTimeout time.Duration = time.Second * 10
	defaultSMTPSendTimeout    time.Duration = time.Second * 10
	defaultSMTPHelo           string        = "durudex"
	defaultSMTPKeepAlive      bool          = true

	// Email templates defaults.
	defaultEmailTemplateVerification string = "./web/template/verification.html"
	defaultEmailTemplateLoggedIn     string = "./web/template/logged_in.html"
	defaultEmailTemplateRegister     string = "./web/template/register.html"
)
