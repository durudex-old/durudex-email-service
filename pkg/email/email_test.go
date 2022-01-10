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

package email

import "testing"

// Testing generating body from html templates.
func TestGenerateBodyFromHTML(t *testing.T) {
	// Creating a new message.
	msg := SendEmailInput{To: "example@example.com", Subject: "Test Message"}

	// Test email message input.
	type TestEmailInput struct{ Name string }

	// Create html template input.
	templateInput := TestEmailInput{Name: "World"}

	// Generate email html template.
	err := msg.GenerateBodyFromHTML("./fixtures/hello_world.html", templateInput)
	if err != nil {
		t.Errorf("error generating body from html: %s", err.Error())
	}

	// Check is empty.
	if msg.Body == "" {
		t.Error("error loading body is empty!")
	}
}

// Testing validation email input.
func TestValidate(t *testing.T) {
	// Creating a new message.
	msg := SendEmailInput{To: "example@example.com", Subject: "Test Message"}

	// Test email message input.
	type TestEmailInput struct{ Name string }

	// Create html template input.
	templateInput := TestEmailInput{Name: "World"}

	// Generate email html template.
	err := msg.GenerateBodyFromHTML("./fixtures/hello_world.html", templateInput)
	if err != nil {
		t.Errorf("error generating body from html: %s", err.Error())
	}

	// Validation email input.
	if err := msg.Validate(); err != nil {
		t.Errorf("error validating input: %s", err.Error())
	}
}
