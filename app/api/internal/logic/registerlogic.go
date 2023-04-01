package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"goZeroDemo/app/api/internal/svc"
	"goZeroDemo/app/api/internal/types"
	"goZeroDemo/app/rpc/userclient"
	"log"
	"net/http"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterRes, err error) {

	if req.Username == "" || req.Password == "" {
		return &types.RegisterRes{
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
		return &types.RegisterRes{
			Code: http.StatusInternalServerError,
			Msg:  "internal error",
		}, err
	}

	return &types.RegisterRes{
		Code: res.Code,
		Msg:  "register successfully",
	}, nil
}
