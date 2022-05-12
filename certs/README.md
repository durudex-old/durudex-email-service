# Certificates

You need to move your certificates to this directory, and make configuration changes:
```yml
grpc:
  tls:
    ca-cert: "./certs/you-ca-cert.pem"
    cert: "./certs/you-cert.pem"
    key: "./certs/you-key.pem"
```

**If you do not want to use tls connection change**:
```yml
grpc:
  tls:
    enable: false
```
