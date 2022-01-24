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

// Testing checking email length and characters.
func TestValidate_IsEmailValid(t *testing.T) {
	// Testing args.
	type args struct{ email string }

	// Tests structures.
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "OK",
			args: args{email: "example@example.example"},
			want: true,
		},
		{
			name: "Not correct email",
			args: args{email: "example.example"},
			want: false,
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Check email length and characters.
			got := IsEmailValid(tt.args.email)

			// Check status.
			if got != tt.want {
				t.Error("error status are not similar")
			}
		})
	}
}
