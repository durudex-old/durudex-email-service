# Certificates

You need to generate certificates for this program to work. How to do it you can find here - [click](https://github.com/durudex/durudex-gateway/blob/main/certs/README.md).

### Required certificates:
+ email.service-cert.pem
+ email.service-key.pem
+ rootCA.pem

**If you do not want to use tls connection change [server.tls](https://github.com/durudex/durudex-email-service/blob/main/configs/main.yml) configuration to `false`**:
```yml
server:
    tls: false
```
