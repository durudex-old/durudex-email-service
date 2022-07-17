# Copyright © 2021-2022 Durudex
#
# This file is part of Durudex: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as
# published by the Free Software Foundation, either version 3 of the
# License, or (at your option) any later version.
#
# Durudex is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License
# along with Durudex. If not, see <https://www.gnu.org/licenses/>.

.PHONY: build
build:
	go mod download && CGO_ENABLE=0 GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

.PHONY: run
run: build
	docker-compose up --remove-orphans app

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test: lint
	go test -v ./...

.PHONY: buf
buf: buf-lint
	buf generate proto/src/api --path proto/src/api/durudex/v1/email_user.proto

.PHONY: buf-lint
buf-lint:
	buf lint proto/src/api/durudex/v1/email_user.proto

.DEFAULT_GOAL := run
