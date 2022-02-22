# Twirp Extensions

Generates some useful code

`rpc_service.twirpex.go`

```go
package pkg

import (
    "github.com/albenik/twirpex"
)

func NewRPCServiceServerEx(svc rpcService, opts ...interface{}) twirpex.TwirpServer {
    return NewRPCServiceServer(svc, opts...).(*rpcService)
}

func (*rpcServiceServer) TwirpServiceMeta() *twirpex.ServiceMeta {
    return &twirpex.ServiceMeta{
        PackageName:     "protobuf.package.v1",
        ServiceName:     "RPCService",
        ServiceFullName: "protobuf.package.v1.RPCService",
        MethodsNames: []string{
            "Foo",
            "Bar",
        },
    }
}

```

# Install

```
go install github.com/albenik/twirpex/cmd/protoc-gen-twirpex@latest
```

# Use

## With `protoc`

```
protoc -twirpex_out=path/to/doc/folder -twirpex_opt=paths=source_relative twirp/service/v1/service.proto
```

## With `buf` https://buf.build

`buf.gen.yaml`:

```yaml
version: v1
plugins:
    -   name: twirpex
        out: path/to/doc/folder
        opt: paths=source_relative
```
