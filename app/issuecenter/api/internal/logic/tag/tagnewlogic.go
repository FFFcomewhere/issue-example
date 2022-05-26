package tag

import (
	"context"
	"github.com/FFFcomewhere/issue-example/app/issuecenter/model"
	"github.com/FFFcomewhere/issue-example/common/xerr"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

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
	//开启事务　插入数据
	if err := l.svcCtx.TagModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		tag := new(model.Tag)
		tag.Name = req.Name

		_, err := l.svcCtx.TagModel.Insert(context, session, tag)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "new tag db tag Insert err:%v,tag:%+v", err, tag)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &types.TagNewResp{}, nil
}
