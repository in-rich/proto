# Proto

Shared protobuf definitions for our GRPC services.

## Requirements

- [Go](https://go.dev/doc/install)
- Make:
    - macOS:
      ```bash
      brew install make
      ```
    - Ubuntu:
      ```bash
      sudo apt-get install make
      ```
    - Windows: Install [chocolatey](https://chocolatey.org/install) (from a PowerShell with admin privileges), then run:
      ```bash
      choco install make
      ```
- Protobuf: download and install [latest release](https://github.com/protocolbuffers/protobuf/releases),
  or use a package manager:
    - macOS: `brew install protobuf`
    - Linux: `apt install -y protobuf-compiler`
    - Windows: using [chocolatey](https://chocolatey.org/install) -> `choco install protoc`

## Installation

Install dependencies:

```bash
make install
```

## Usage

Once you have modified/added proto file, you must generate the corresponding definitions:

```bash
make proto
```

## Import

### Go

> You might need to set up the GOPRIVATE variable to access this repository. Make sure you are logged in 
> to a GitHub account that has access to this repository.
> 
> ```bash
> export GOPRIVATE=github.com/in-rich
> ```

```bash
go get github.com/in-rich/proto/proto-go
```
