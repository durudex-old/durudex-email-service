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
	"bytes"
	"errors"
	"text/template"
)

type SendEmailInput struct {
	To      string
	Subject string
	Body    string
}

type Email interface {
	Send(input SendEmailInput) (bool, error)
}

// Generating body from html templates.
func (s *SendEmailInput) GenerateBodyFromHTML(templateFileName string, data interface{}) error {
	// Parsing htpml file.
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}

	// Executing date in html template.
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}

	s.Body = buf.String()

	return nil
}

// Validation email input.
func (e *SendEmailInput) Validate() error {
	// If the email is not specified.
	if e.To == "" {
		return errors.New("empty to")
	}
	// If subject and body are not specified.
	if e.Subject == "" || e.Body == "" {
		return errors.New("empty subject/body")
	}
	// Check email length and characters.
	if !IsEmailValid(e.To) {
		return errors.New("invalid to email")
	}

	return nil
}
