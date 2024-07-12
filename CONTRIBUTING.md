# Contributions

All definitions are located under the `/proto` directory. This is the only place where you should edit files.
Other directories contain auto-generated definition for our different technical stacks.

## Project structure

The definitions are grouped by services. Under the `/proto` directory, each service that uses GRPC has its own
directory. For example, the `user` service has its definitions under `/proto/user`.

All definitions should be siblings for a given service (don't use subdirectories, as they will be ignored by the
compiler command).

## Language specifics

Some languages require specific options in the proto file, so they can be properly compiled. Every proto file
should comply to the specifics of every language that we support.

### Go

In Go protobuf, you must specify which package will be generated.

```protobuf
option go_package = "proto-go/[SERVICE];[IMPORT]";
```

Where:
    - Service is the name of the service you generate the profo for
    - Import is the import path for the generated code

For example:
    
```protobuf
option go_package = "proto-go/user;user";
```

Will allow you to import users definitions, using the `user` package.

```go
import "github.com/in-rich/proto/proto-go/user"

// Use proto definitions for user.
user_pb.User
```

> It is recommended you set the package name to end with `_pb`, so it is clear that you are importing protobuf
> definitions.

## Publishing

When you are done with adding or modifying definitions, you should regenerate the definitions.

```bash
make proto
```

This repository uses tags as versioning. When you are ready to publish a new version, you should create a new tag
and push it to the repository.

## Maintenance

### Updating dependencies

It is important to keep track of major releases for protobuf. To do so, make sure you have the latest version of
protoc installed:

```bash
protoc --version
# libprotoc 27.2
```

Then, you can upgrade dependencies for each language:

```bash
# go (cd ./proto-go)
go get -u ./... && go mod tidy
```

Then, don't forget to regenerate definitions, to ensure everything went smoothly:

```bash
make proto
```
