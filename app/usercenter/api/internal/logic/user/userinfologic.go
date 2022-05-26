package user

import (
	"context"
	"github.com/FFFcomewhere/issue-example/app/usercenter/rpc/types/builder"

	"github.com/FFFcomewhere/issue-example/app/usercenter/api/internal/svc"
	"github.com/FFFcomewhere/issue-example/app/usercenter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.GetUserInfoReq) (resp *types.GetUserInfoResp, err error) {
	InfoResp, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &builder.GetUserInfoReq{
		Userid:   req.Userid,
		Mobile:   req.Mobile,
		Username: req.Username,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetUserInfoResp{
		Userid:   InfoResp.Userid,
		Mobile:   InfoResp.Mobile,
		Username: InfoResp.Username,
	}, nil
}
