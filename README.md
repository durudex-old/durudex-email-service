<div align="center">
    <a href="https://discord.gg/4qcXbeVehZ">
        <img alt="Discord" src="https://img.shields.io/discord/882288646517035028?label=%F0%9F%92%AC%20discord">
    </a>
    <a href="https://github.com/durudex/durudex-email-service/blob/main/COPYING">
        <img alt="License" src="https://img.shields.io/github/license/durudex/durudex-email-service?label=%F0%9F%93%95%20license">
    </a>
    <a href="https://github.com/durudex/durudex-email-service/stargazers">
        <img alt="GitHub Stars" src="https://img.shields.io/github/stars/durudex/durudex-email-service?label=%E2%AD%90%20stars&logo=sdf">
    </a>
    <a href="https://github.com/durudex/durudex-email-service/network">
        <img alt="GitHub Forks" src="https://img.shields.io/github/forks/durudex/durudex-email-service?label=%F0%9F%93%81%20forks">
    </a>
</div>

<h1 align="center">⚡️ Durudex Email Service</h1>

<p align="center">
Service for sending email messages.
</p>

### 💡 Prerequisites
+ [Go 1.17](https://golang.org/)
+ [Docker](https://www.docker.com)
+ [golangci-lint](https://golangci-lint.run/usage/install/)
+ [grpc](https://grpc.io/docs/languages/go/quickstart/)

## ⚙️ Build & Run
1) Create an `.env` file in the root directory and add the following values from `.env.example`:
```sh
# Config variables.
CONFIG_PATH=configs/main

# Email config variables.
EMAIL_USERNAME=
EMAIL_PASSWORD=
```
2) If you use a default `smtp.host` and `smtp.port` configuration,
create SMTP server with address: `email.durudex.local:25`.
3) Generate certificates, information can be found at [certs/README.md](certs/README.md)

Use `make run` to run and `make build` to build project.

## 👍 Contribute
If you want to say thank you and/or support the active development of [Durudex](https://github.com/durudex):
1) Add a [GitHub Star](https://github.com/durudex/durudex-notif-service/stargazers) to the project.
2) Join the [Discord Server](https://discord.gg/4qcXbeVehZ).

## ⚠️ License
Copyright © 2021-2022 [Durudex](https://github.com/durudex). Released under the [GNU AGPL v3](https://www.gnu.org/licenses/agpl-3.0.html) license.

#### Third-party library licenses
+ [grpc-go](https://github.com/grpc/grpc-go/blob/master/LICENSE)
+ [protobuf](https://github.com/protocolbuffers/protobuf/blob/master/LICENSE)
+ [zerolog](https://github.com/rs/zerolog/blob/master/LICENSE)
+ [viper](https://github.com/spf13/viper/blob/master/LICENSE)
+ [go-simple-mail](https://github.com/xhit/go-simple-mail/blob/master/LICENSE)
