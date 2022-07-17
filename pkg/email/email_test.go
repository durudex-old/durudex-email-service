/*
 * Copyright © 2022 Durudex
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

package email

import "testing"

// Testing generating body from html templates.
func TestSendEmailInput_GenerateBodyFromHTML(t *testing.T) {
	// Testing body input.
	type bodyInput struct{ Name string }

	// Testing args.
	type args struct {
		path    string
		email   string
		subject string
		data    bodyInput
	}

	// Tests structures.
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				path:    "./fixtures/hello_world.html",
				email:   "example@example.example",
				subject: "Test Message",
				data:    bodyInput{Name: "World"},
			},
			wantErr: false,
		},
		{
			name: "File not found",
			args: args{
				path:    "./fixtures/not_found.html",
				email:   "example@example.example",
				subject: "Test Message",
				data:    bodyInput{},
			},
			wantErr: true,
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new message.
			msg := SendEmailInput{To: tt.args.email, Subject: tt.args.subject}

			// Generating body from html templates.
			err := msg.GenerateBodyFromHTML(tt.args.path, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Error("error status are not similar")
			}
		})
	}
}

// Testing validation email input.
func TestSendEmailInput_Validate(t *testing.T) {
	// Testing args.
	type args struct {
		email   string
		subject string
		body    string
	}

	// Tests structures.
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				email:   "example@example.example",
				subject: "Test Message",
				body:    "<h1>Hello World</h1>",
			},
			wantErr: false,
		},
		{
			name: "Email and body not correct",
			args: args{
				email:   "example.example",
				subject: "Test Message",
				body:    "",
			},
			wantErr: true,
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new message.
			msg := SendEmailInput{To: tt.args.email, Subject: tt.args.subject}

			// Set body for args.
			msg.Body = tt.args.body

			// Validation email input.
			err := msg.Validate()
			if (err != nil) != tt.wantErr {
				t.Error("error status are not similar")
			}
		})
	}
}
