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
	"github.com/Durudex/durudex-notif-service/internal/config"
	"github.com/Durudex/durudex-notif-service/pkg/email"
)

type Email interface {
	UserCode(to, name string, code uint64) (bool, error)
}

type Service struct {
	Email
}

// Creating a new service.
func NewService(email email.Email, emailConfig config.EmailConfig) *Service {
	return &Service{
		Email: NewEmailService(email, emailConfig),
	}
}
