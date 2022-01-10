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

package email

import (
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

// SMTP client structure.
type Client struct {
	client *mail.SMTPClient
	cfg    *SMTPConfig
}

// SMTP config.
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

// Creating a new SMTP client.
func NewClient(cfg *SMTPConfig) (*Client, error) {
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

	// Noop sending messages.
	if err := client.Noop(); err != nil {
		return nil, err
	}

	return &Client{client: client, cfg: cfg}, nil
}

// Send client message.
func (c *Client) Send(input SendEmailInput) (bool, error) {
	// Check email message input.
	if err := input.Validate(); err != nil {
		return false, err
	}

	// Creating a new message.
	msg := mail.NewMSG()

	// Set from address.
	msg.SetFrom(c.cfg.Username)
	// Set to address.
	msg.AddTo(input.To)
	// Set email subject.
	msg.SetSubject(input.Subject)

	// Set email html template.
	msg.SetBody(mail.TextHTML, input.Body)

	// Check for error in the message.
	if msg.Error != nil {
		return false, msg.Error
	}

	// Send email message.
	if err := msg.Send(c.client); err != nil {
		return false, err
	}

	return true, nil
}
