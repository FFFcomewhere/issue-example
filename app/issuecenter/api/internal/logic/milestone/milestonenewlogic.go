package milestone

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

type MilestonenewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMilestonenewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MilestonenewLogic {
	return &MilestonenewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MilestonenewLogic) Milestonenew(req *types.MilestoneNewReq) (resp *types.MilestoneNewResp, err error) {

	//开启事务　插入数据
	if err := l.svcCtx.MilestoneModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		milestone := new(model.Milestone)
		milestone.Name = req.Name

		_, err := l.svcCtx.MilestoneModel.Insert(context, session, milestone)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "new milestone db milestone Insert err:%v,milestone:%+v", err, milestone)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &types.MilestoneNewResp{}, nil
}
