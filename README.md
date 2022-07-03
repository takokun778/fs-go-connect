# fs-go-connect

```bash
curl --header "Content-Type: application/json" --data '{"name": "Jane"}' http://localhost:8080/greet.v1.GreetService/Greet
```

```bash
grpcurl -protoset <(buf build -o -) -plaintext -d '{"name": "Jane"}' localhost:8080 greet.v1.GreetService/Greet
```

```bash
curl --header "Content-Type: application/json" --data '{"service": "greet.v1.GreetService"}' http://localhost:8080/grpc.health.v1.Health/Check
```

```bash
grpcurl -protoset <(buf build -o -) -plaintext -d '{"service": "greet.v1.GreetService"}' localhost:8080 health.v1.Health/Check
```
