package logic

import (
	"context"
	"github.com/FFFcomewhere/issue-example/app/usercenter/model"
	"github.com/FFFcomewhere/issue-example/common/tool"
	"github.com/FFFcomewhere/issue-example/common/xerr"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/FFFcomewhere/issue-example/app/usercenter/rpc/internal/svc"
	"github.com/FFFcomewhere/issue-example/app/usercenter/rpc/types/builder"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrUserAlreadyRegisterError = xerr.NewErrMsg("user has been registered")

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

func (l *RegisterLogic) Register(in *builder.RegisterReq) (*builder.RegisterResp, error) {
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "username:%s,err:%v", in.Username, err)
	}
	if user != nil {
		return nil, errors.Wrapf(ErrUserAlreadyRegisterError, "Register user exists username:%s,err:%v", in.Username, err)
	}

	//开启事务　插入数据
	if err := l.svcCtx.UserModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		user := new(model.User)
		user.Username = in.Username

		if len(in.Password) > 0 {
			user.Password = tool.Md5ByString(in.Password)
		}

		_, err := l.svcCtx.UserModel.Insert(context, session, user)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Register db user Insert err:%v,user:%+v", err, user)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	//获取用户id
	getUserInfoLogic := NewGetUserInfoLogic(l.ctx, l.svcCtx)
	userinfo, err := getUserInfoLogic.GetUserInfo(&builder.GetUserInfoReq{
		Userid:   0,
		Mobile:   "",
		Username: in.Username,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrUserGetInfoError, "get username  : %v", in.Username)
	}

	//生成令牌
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&builder.GenerateTokenReq{
		UserId: userinfo.Userid,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "GenerateToken userId : %d", userinfo.Userid)
	}

	return &builder.RegisterResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}
