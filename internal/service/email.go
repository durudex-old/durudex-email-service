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

package service

import (
	"fmt"

	"github.com/Durudex/durudex-notif-service/internal/config"
	"github.com/Durudex/durudex-notif-service/pkg/email"
)

type EmailService struct {
	email  email.Email
	config config.EmailConfig
}

// Creating a new email service.
func NewEmailService(email email.Email, emailConfig config.EmailConfig) *EmailService {
	return &EmailService{
		email:  email,
		config: emailConfig,
	}
}

type verificationEmailInput struct {
	Name string
	Code int32
}

// Send to user email verification code.
func (s *EmailService) UserVerifyCode(to, name string, code int32) (bool, error) {
	msg := email.SendEmailInput{
		To:      to,
		Subject: "Verification Code",
	}

	// Generate email html template.
	templateInput := verificationEmailInput{Name: name, Code: code}
	msg.GenerateBodyFromHTML(s.config.Template.Verification, templateInput)

	// Send email message.
	if err := s.email.Send(msg); err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil
}
