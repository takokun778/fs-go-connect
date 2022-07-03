package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	greetv1 "connect/gen/greet/v1"
	"connect/gen/greet/v1/greetv1connect"

	grpchealth "github.com/bufbuild/connect-grpchealth-go"

	"github.com/bufbuild/connect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type GreetServer struct{}

func (s *GreetServer) Greet(
	ctx context.Context,
	req *connect.Request[greetv1.GreetRequest],
) (*connect.Response[greetv1.GreetResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&greetv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}

func main() {
	compressMin := connect.WithCompressMinBytes(1024)

	greeter := &GreetServer{}

	mux := http.NewServeMux()

	mux.Handle(greetv1connect.NewGreetServiceHandler(greeter, compressMin))

	mux.Handle(grpchealth.NewHandler(
		grpchealth.NewStaticChecker(greetv1connect.GreetServiceName),
		compressMin,
	))

	http.ListenAndServe(
		":8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
