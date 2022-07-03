module example

go 1.18

require (
	connect v0.0.0-00010101000000-000000000000
	github.com/bufbuild/connect-go v0.1.1
	golang.org/x/net v0.0.0-20220630215102-69896b714898
)

require (
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

replace connect => ./grpc
