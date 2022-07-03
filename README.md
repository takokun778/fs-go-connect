# fs-go-connect

```bash
curl --header "Content-Type: application/json" --data '{"name": "Jane"}' http://localhost:8080/greet.v1.GreetService/Greet
```

```bash
grpcurl -protoset <(buf build -o -) -plaintext -d '{"name": "Jane"}' localhost:8080 greet.v1.GreetService/Greet
```
