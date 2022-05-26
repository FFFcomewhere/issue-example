package logic

import (
	"context"
	"github.com/FFFcomewhere/issue-example/app/usercenter/model"
	"github.com/FFFcomewhere/issue-example/common/xerr"
	"github.com/pkg/errors"

	"github.com/FFFcomewhere/issue-example/app/usercenter/rpc/internal/svc"
	"github.com/FFFcomewhere/issue-example/app/usercenter/rpc/types/builder"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrUserGetInfoError = xerr.NewErrMsg("user get info error")
var ErrUserNoExistsError = xerr.NewErrMsg("user no exist")

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *builder.GetUserInfoReq) (*builder.GetUserInfoResp, error) {
	var user *model.User
	var err error
	//分三种情况查询
	if in.Userid != 0 {
		user, err = l.svcCtx.UserModel.FindOne(l.ctx, in.Userid)
	} else if in.Mobile != "" {
		user, err = l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	} else if in.Username != "" {
		user, err = l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	}

	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetUserInfo: find user err , id:%d , err:%v", in.Userid, err)
	}
	if user == nil {
		return nil, errors.Wrapf(ErrUserGetInfoError, "userid:%v, mobile:%v, username:%v ", in.Userid, user.Mobile, user.Username)
	}

	return &builder.GetUserInfoResp{
		Userid:   user.Id,
		Mobile:   user.Mobile,
		Username: user.Username,
	}, nil

}
