.PHONY = generate

generate:
	goa gen github.com/allinbits/sdk-service-meta

install-goa:
	go get -u goa.design/goa/v3
	go get -u goa.design/goa/v3/...
	go install -v github.com/golang/protobuf/protoc-gen-go@latest
	go install -v google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest