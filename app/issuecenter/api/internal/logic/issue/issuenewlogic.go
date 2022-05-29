package issue

import (
	"context"
	"github.com/FFFcomewhere/issue-example/app/issuecenter/model"
	"github.com/FFFcomewhere/issue-example/app/usercenter/rpc/types/builder"
	"github.com/FFFcomewhere/issue-example/common/xerr"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/svc"
	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IssuenewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIssuenewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IssuenewLogic {
	return &IssuenewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IssuenewLogic) Issuenew(req *types.IssueNewReq) (resp *types.IssueNewResp, err error) {

	//开启事务　插入数据
	if err := l.svcCtx.IssueModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		issue := new(model.Issue)
		issue.Name = req.Name

		//依次判断字段是否需要赋值
		if req.UserName != "" {
			userRpcReq := builder.GetUserInfoReq{
				Userid:   0,
				Mobile:   "",
				Username: req.UserName,
			}

			user, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &userRpcReq)
			if err == model.ErrNotFound {
				issue.Userid = 0
			}
			issue.Userid = user.Userid
		}

		if req.MilestoneName != "" {
			milestone, err := l.svcCtx.MilestoneModel.FindOneByName(l.ctx, req.MilestoneName)
			if err == model.ErrNotFound {
				issue.Milestoneid = 0
			}
			issue.Milestoneid = milestone.Id
		}

		if req.TagName != "" {
			milestone, err := l.svcCtx.TagModel.FindOneByName(l.ctx, req.TagName)
			if err == model.ErrNotFound {
				issue.Tagid = 0
			}
			issue.Tagid = milestone.Id
		}

		_, err := l.svcCtx.IssueModel.Insert(context, session, issue)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "new issue db issue Insert err:%v,issue:%+v", err, issue)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &types.IssueNewResp{}, nil
}
