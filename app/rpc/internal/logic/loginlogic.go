package logic

import (
	"context"
	"goZeroDemo/app/rpc/user"
	"goZeroDemo/model"

	"github.com/zeromicro/go-zero/core/logx"
	"goZeroDemo/app/rpc/internal/svc"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginRes, error) {
	// todo: add your logic here and delete this line
	username := in.Username
	password := in.Password
	var userSubject model.User
	l.svcCtx.Db.Model(&model.User{}).Where("username=?", username).First(&userSubject)
	if password != userSubject.Password {
		return &user.LoginRes{
			Code: 400,
			Msg:  "password incorrect",
		}, nil
	}
	return &user.LoginRes{
		Code: 200,
		Msg:  "login successfully",
	}, nil
}
