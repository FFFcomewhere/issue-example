package logic

import (
	"context"
	"github.com/FFFcomewhere/issue-example/app/usercenter/model"
	"github.com/FFFcomewhere/issue-example/common/tool"
	"github.com/FFFcomewhere/issue-example/common/xerr"
	"github.com/pkg/errors"

	"github.com/FFFcomewhere/issue-example/app/usercenter/rpc/internal/svc"
	"github.com/FFFcomewhere/issue-example/app/usercenter/rpc/types/builder"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrUsernamePwdError = xerr.NewErrMsg("username or passowrd error")

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

func (l *LoginLogic) Login(in *builder.LoginReq) (*builder.LoginResp, error) {
	//通过用户名登录
	userid, err := l.loginByUsername(in.Username, in.Password)

	if err != nil {
		return nil, err
	}

	//生成令牌
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&builder.GenerateTokenReq{
		UserId: userid,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "GenerateToken userId : %d", userid)
	}

	return &builder.LoginResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}

func (l *LoginLogic) loginByUsername(username, password string) (int64, error) {

	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, username)
	if err != nil && err != model.ErrNotFound {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "find user by username error，username:%s,err:%v", username, err)
	}
	if user == nil {
		return 0, errors.Wrapf(ErrUserGetInfoError, "username:%s", username)
	}

	if !(tool.Md5ByString(password) == user.Password) {
		return 0, errors.Wrap(ErrUsernamePwdError, "password error")
	}

	return user.Id, nil
}
