package issuestate

import (
	"context"

	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/svc"
	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IssuestatepushLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIssuestatepushLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IssuestatepushLogic {
	return &IssuestatepushLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IssuestatepushLogic) Issuestatepush(req *types.IssueStatePushReq) (resp *types.IssueStatePushResp, err error) {
	// todo: add your logic here and delete this line

	return
}
