package logic

import (
	"context"
	"goZeroDemo/app/api/internal/svc"
	"goZeroDemo/app/api/internal/types"
	"goZeroDemo/app/rpc/userclient"
	"log"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRes, err error) {

	if req.Username == "" || req.Password == "" {
		return &types.LoginRes{
			Code: 400,
			Msg:  "username and password cannot be null",
		}, nil
	}

	res, err := l.svcCtx.UserRpc.Register(l.ctx, &userclient.RegisterReq{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		log.Printf("register error %v", err)
		return &types.LoginRes{
			Code: res.Code,
			Msg:  "internal error",
		}, err
	}

	return &types.LoginRes{
		Code: res.Code,
		Msg:  "register successfully",
	}, nil
}
