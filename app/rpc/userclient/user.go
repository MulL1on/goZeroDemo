// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userclient

import (
	"context"

	"goZeroDemo/app/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	LoginReq    = user.LoginReq
	LoginRes    = user.LoginRes
	RegisterReq = user.RegisterReq
	RegisterRes = user.RegisterRes

	User interface {
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRes, error)
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRes, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUser) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}
