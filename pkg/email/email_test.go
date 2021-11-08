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

import "testing"

var templateFile = "./templates/test.html"

type testEmailInput struct {
	Name string
}

// Testing generating body from html templates.
func TestGenerateBodyFromHTML(t *testing.T) {
	// Email input.
	input := SendEmailInput{
		To:      "v1def@durudex.local",
		Subject: "Testing subject",
	}

	// Generating body from html template.
	err := input.GenerateBodyFromHTML(templateFile, testEmailInput{Name: "V1def"})
	if err != nil {
		t.Fatalf("error generating body from html template: %s", err.Error())
	}

	// Check body.
	if input.Body == "" {
		t.Fatal("error generating body is nil")
	}
}

// Tesing validation email input.
func TestValidate(t *testing.T) {
	// Email input.
	input := SendEmailInput{
		To:      "v1def@durudex.local",
		Subject: "Testing subject",
	}

	// Generating body from html template.
	templateInput := testEmailInput{Name: "V1def"}
	err := input.GenerateBodyFromHTML(templateFile, templateInput)
	if err != nil {
		t.Fatalf("error generating body from html template: %s", err.Error())
	}

	// Validation email input
	if err := input.Validate(); err != nil {
		t.Fatalf("error validating email input: %s", err.Error())
	}
}
