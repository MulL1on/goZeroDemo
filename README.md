```
├─go.mod
├─go.sum
├─model
|   └user.go
├─app
|  ├─rpc
|  |  ├─user.go
|  |  ├─user.proto
|  |  ├─userclient
|  |  |     └user.go
|  |  ├─user
|  |  |  ├─user.pb.go
|  |  |  └user_grpc.pb.go
|  |  ├─internal
|  |  |    ├─svc
|  |  |    |  └servicecontext.go
|  |  |    ├─server
|  |  |    |   └userserver.go
|  |  |    ├─logic
|  |  |    |   ├─loginlogic.go
|  |  |    |   └registerlogic.go
|  |  |    ├─config
|  |  |    |   └config.go
|  |  ├─etc
|  |  |  └user.yaml
|  ├─api
|  |  ├─user.api
|  |  ├─user.go
|  |  ├─internal
|  |  |    ├─types
|  |  |    |   └types.go
|  |  |    ├─svc
|  |  |    |  └servicecontext.go
|  |  |    ├─logic
|  |  |    |   ├─loginlogic.go
|  |  |    |   └registerlogic.go
|  |  |    ├─handler
|  |  |    |    ├─loginhandler.go
|  |  |    |    ├─registerhandler.go
|  |  |    |    └routes.go
|  |  |    ├─config
|  |  |    |   └config.go
|  |  ├─etc
|  |  |  └user.yaml
```

