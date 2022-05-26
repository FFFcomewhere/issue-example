package tag

import (
	"context"

	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/svc"
	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TagsignalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTagsignalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TagsignalLogic {
	return &TagsignalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TagsignalLogic) Tagsignal(req *types.TagSignalReq) (resp *types.TagSignalResp, err error) {
	// todo: add your logic here and delete this line

	return
}
