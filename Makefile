proto-go:
	protoc \
		--go_out=. --go_opt=paths=import \
        --go-grpc_out=. --go-grpc_opt=paths=import \
        proto/*/*.proto

install-go:
	cd proto-go && \
		go get google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc && \
		go install google.golang.org/protobuf/cmd/protoc-gen-go && \
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc && \
		go mod tidy

install: install-go

proto: proto-go

.PHONY: proto-go install-go install proto
