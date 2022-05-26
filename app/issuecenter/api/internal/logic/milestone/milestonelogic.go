package milestone

import (
	"context"
	"github.com/FFFcomewhere/issue-example/app/issuecenter/model"
	"github.com/FFFcomewhere/issue-example/common/xerr"
	"github.com/pkg/errors"

	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/svc"
	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MilestoneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMilestoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MilestoneLogic {
	return &MilestoneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MilestoneLogic) Milestone(req *types.MilestoneReq) (resp *types.MilestoneResp, err error) {
	page := req.Page
	if page < 1 {
		page = 1
	}

	whereBuilder := l.svcCtx.MilestoneModel.RowBuilder()
	milestoneList, err := l.svcCtx.MilestoneModel.FindPageListByIdASC(l.ctx, whereBuilder, (req.Page-1)*req.PageSize, req.PageSize)

	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get IssueList db err. rowType: %s ,err : %v", "issueid", err)
	}

	var milestoneInfoList []types.MilestoneInfo
	if len(milestoneList) > 0 {
		for _, e := range milestoneList {
			milestone, err := l.newMilestoneInfo(e, req)
			if err != nil {
				return nil, err
			}

			milestoneInfoList = append(milestoneInfoList, *milestone)
		}
	}

	return &types.MilestoneResp{
		List: milestoneInfoList,
	}, nil
}

func (l *MilestoneLogic) newMilestoneInfo(milestone *model.Milestone, req *types.MilestoneReq) (*types.MilestoneInfo, error) {
	var tempMilestone types.MilestoneInfo
	tempMilestone.Name = milestone.Name
	tempMilestone.Milestoneid = milestone.Id
	tempMilestone.UpdateTime = milestone.UpdateTime.String()

	return &tempMilestone, nil
}
