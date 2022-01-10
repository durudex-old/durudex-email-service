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

package service

import (
	"github.com/durudex/durudex-email-service/internal/config"
	"github.com/durudex/durudex-email-service/internal/domain"
	"github.com/durudex/durudex-email-service/pkg/email"
)

// Email service structure.
type EmailService struct {
	email email.Email
	cfg   config.EmailConfig
}

// Creating a new email service.
func NewEmailService(email email.Email, cfg config.EmailConfig) *EmailService {
	return &EmailService{email: email, cfg: cfg}
}

// Send to user email verification code.
func (s *EmailService) UserCode(to, username string, code uint64) (bool, error) {
	// Create a new email message.
	msg := email.SendEmailInput{To: to, Subject: "Verification Code"}

	// Create html template input.
	templateInput := domain.VerificationEmailInput{Username: username, Code: code}

	// Generate email html template.
	err := msg.GenerateBodyFromHTML(s.cfg.Template.Verification, templateInput)
	if err != nil {
		return false, err
	}

	// Send email message.
	return s.email.Send(msg)
}

// Send to user email register information.
func (s *EmailService) UserRegister(to, username string) (bool, error) {
	// Create a new email message.
	msg := email.SendEmailInput{To: to, Subject: "You have successfully created a new account"}

	// Create html template input.
	templateInput := domain.RegisterEmailInput{Username: username}

	// Generate email html template.
	err := msg.GenerateBodyFromHTML(s.cfg.Template.Register, templateInput)
	if err != nil {
		return false, err
	}

	// Send email message.
	return s.email.Send(msg)
}

// Send to user email logged in information.
func (s *EmailService) UserLoggedIn(to, ip string) (bool, error) {
	// Create a new email message.
	msg := email.SendEmailInput{To: to, Subject: "You have successfully logged in"}

	// Create html template input.
	templateInput := domain.LoggedInEmailInput{IP: ip}

	// Generate email html template.
	err := msg.GenerateBodyFromHTML(s.cfg.Template.LoggedIn, templateInput)
	if err != nil {
		return false, err
	}

	// Send email message.
	return s.email.Send(msg)
}
