# Certificates

You need to generate certificates for this program to work. How to do it you can find here - [click](https://github.com/Durudex/durudex-gateway/blob/main/cert/README.md).

### Required certificates:
+ emailservice-cert.pem
+ emailservice-key.pem
+ rootCA.pem

**If you do not want to use tls connection change [server.tls](https://github.com/Durudex/durudex-email-service/blob/main/configs/main.yml) configuration to `false`**:
```yml
server:
    tls: false
```
