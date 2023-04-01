package logic

import (
	"context"
	"goZeroDemo/app/rpc/internal/svc"
	"goZeroDemo/app/rpc/user"
	"goZeroDemo/model"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterRes, error) {
	username := in.Username
	password := in.Password

	err := l.svcCtx.Db.Create(&model.User{Username: username, Password: password}).Error

	if err != nil {
		return &user.RegisterRes{
			Msg:  "internal error",
			Code: http.StatusInternalServerError,
		}, err
	}

	return &user.RegisterRes{
		Msg:  "register successfully",
		Code: 200,
	}, nil
}
