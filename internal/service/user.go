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

package service

import (
	"github.com/durudex/durudex-email-service/internal/config"
	"github.com/durudex/durudex-email-service/internal/domain"
	"github.com/durudex/durudex-email-service/pkg/email"
)

// Email interface.
type User interface {
	Code(to, username string, code uint64) error
	Register(to, username string) error
	LoggedIn(to, ip string) error
}

// Email user service structure.
type UserService struct {
	email email.Email
	cfg   config.EmailConfig
}

// Creating a new email user service.
func NewUserService(email email.Email, cfg config.EmailConfig) *UserService {
	return &UserService{email: email, cfg: cfg}
}

// Send to user email verification code.
func (s *UserService) Code(to, username string, code uint64) error {
	// Create a new email message.
	msg := email.SendEmailInput{To: to, Subject: "Verification Code"}

	// Create html template input.
	templateInput := domain.VerificationEmailInput{Username: username, Code: code}

	// Generate email html template.
	if err := msg.GenerateBodyFromHTML(s.cfg.Template.Verification, templateInput); err != nil {
		return err
	}

	// Send email message.
	return s.email.Send(msg)
}

// Send to user email register information.
func (s *UserService) Register(to, username string) error {
	// Create a new email message.
	msg := email.SendEmailInput{To: to, Subject: "You have successfully created a new account"}

	// Create html template input.
	templateInput := domain.RegisterEmailInput{Username: username}

	// Generate email html template.
	if err := msg.GenerateBodyFromHTML(s.cfg.Template.Register, templateInput); err != nil {
		return err
	}

	// Send email message.
	return s.email.Send(msg)
}

// Send to user email logged in information.
func (s *UserService) LoggedIn(to, ip string) error {
	// Create a new email message.
	msg := email.SendEmailInput{To: to, Subject: "You have successfully logged in"}

	// Create html template input.
	templateInput := domain.LoggedInEmailInput{Ip: ip}

	// Generate email html template.
	if err := msg.GenerateBodyFromHTML(s.cfg.Template.LoggedIn, templateInput); err != nil {
		return err
	}

	// Send email message.
	return s.email.Send(msg)
}
