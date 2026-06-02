# golang_grpc_mongodb_crud

# Generating Go Files from a Protocol Buffer (.proto) File

Before generating Go code from a `.proto` file, ensure that **Protocol Buffers Compiler (`protoc`)** is installed on your system and added to your system's `PATH`.

## Install Required Go Plugins

Install the following plugins, which are required for generating Go and gRPC code:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Make sure the Go binaries directory (`$GOPATH/bin` or `%USERPROFILE%\go\bin`) is included in your system's `PATH`.

## Generate Go Files

Run the following command from the root of your project:

```bash
protoc --go_out=. --go-grpc_out=. proto/user.proto
```

This command generates:

* `user.pb.go` – Go code for Protocol Buffers messages.
* `user_grpc.pb.go` – Go code for gRPC client and server interfaces.

These generated files can then be used in your Go gRPC application.

