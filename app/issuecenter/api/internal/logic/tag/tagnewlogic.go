package tag

import (
	"context"

	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/svc"
	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TagnewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTagnewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TagnewLogic {
	return &TagnewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TagnewLogic) Tagnew(req *types.TagNewReq) (resp *types.TagNewResp, err error) {
	// todo: add your logic here and delete this line

	return
}
