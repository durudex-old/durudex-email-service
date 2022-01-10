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
	// Server defaults.
	defaultServerHost = "emailservice.durudex.local"
	defaultServerPort = "8002"
	defaultServerTLS  = true

	// SMTP defaults.
	defaultSMTPHost           = "smtp.durudex.local"
	defaultSMTPPort           = 25
	defaultSMTPConnectTimeout = time.Second * 10
	defaultSMTPSendTimeout    = time.Second * 10
	defaultSMTPHelo           = "durudex"
	defaultSMTPKeepAlive      = true

	// Email templates defaults.
	defaultEmailTemplateVerification = "./web/template/verification.html"
	defaultEmailTemplateLoggedIn     = "./web/template/logged_in.html"
	defaultEmailTemplateRegister     = "./web/template/register.html"
)
