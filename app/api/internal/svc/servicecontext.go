package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"goZeroDemo/app/api/internal/config"
	"goZeroDemo/app/rpc/userclient"
)

type ServiceContext struct {
	UserRpc userclient.User
	Config  config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		Config:  c,
	}
}
