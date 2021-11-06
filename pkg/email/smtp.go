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

package email

import (
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

type SMTP struct {
	client *mail.SMTPClient
}

type SMTPConfig struct {
	Host           string
	Port           int
	Username       string
	Password       string
	ConnectTimeout time.Duration
	SendTimeout    time.Duration
	Helo           string
	KeepAlive      bool
}

// Creating a new SMTP.
func NewSMTP(cfg *SMTPConfig) (*SMTP, error) {
	// Creating a new SMTP server.
	server := &mail.SMTPServer{
		Host:           cfg.Host,
		Port:           cfg.Port,
		Username:       cfg.Username,
		Password:       cfg.Password,
		ConnectTimeout: cfg.ConnectTimeout,
		SendTimeout:    cfg.SendTimeout,
		Helo:           cfg.Helo,
		KeepAlive:      cfg.KeepAlive,
		Authentication: mail.AuthNone,
		Encryption:     mail.EncryptionNone,
	}

	// Connecting to SMTP server.
	client, err := server.Connect()
	if err != nil {
		return nil, err
	}

	return &SMTP{client: client}, nil
}

// Sending SMTP message.
func (s *SMTP) Send(input SendEmailInput) error {
	// Creating a new message.
	msg := mail.NewMSG()

	msg.SetFrom("")
	msg.AddTo(input.To)
	msg.SetSubject(input.Subject)

	msg.SetBody(mail.TextHTML, input.Body)

	return msg.Send(s.client)
}
