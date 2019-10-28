module app

go 1.12

require (
	google.golang.org/grpc v1.24.0
	helloworld v0.0.0
)

replace helloworld v0.0.0 => ./proto/helloworld
