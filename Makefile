# Copyright © 2021 Durudex

# This file is part of Durudex: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as
# published by the Free Software Foundation, either version 3 of the
# License, or (at your option) any later version.

# Durudex is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
# GNU Affero General Public License for more details.

# You should have received a copy of the GNU Affero General Public License
# along with Durudex. If not, see <https://www.gnu.org/licenses/>.

.PHONY: download
download:
	go mod download

.PHONY: build
build: download
	go build -o .bin/notifservice.exe ./cmd/notifservice/main.go

.PHONY: run
run: download
	go run ./cmd/notifservice/main.go

.PHONY: test
test:
	go test -v ./...

.PHONY: protoc
protoc:
	protoc \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		internal/delivery/grpc/protobuf/*.proto

.PHONY: protoc-types
protoc-types:
	protoc \
		--go_out=. \
		--go_opt=paths=source_relative \
		internal/delivery/grpc/protobuf/types/*.proto

.DEFAULT_GOAL := run