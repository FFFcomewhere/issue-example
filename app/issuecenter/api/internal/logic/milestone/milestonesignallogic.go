package milestone

import (
	"context"
	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/logic/issue"
	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/svc"
	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/types"
	"github.com/FFFcomewhere/issue-example/app/issuecenter/model"
	"github.com/FFFcomewhere/issue-example/common/xerr"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type MilestonesignalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMilestonesignalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MilestonesignalLogic {
	return &MilestonesignalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MilestonesignalLogic) Milestonesignal(req *types.MilestoneSignalReq) (resp *types.MilestoneSignalResp, err error) {
	//修改issue基本信息
	if req.ReName != "" {
		err := l.updataMilestone(req)
		if err != nil {
			return nil, err
		}
	}

	//删除提案
	if req.IfDelete == true {
		err := l.deleteMilestone(req)
		if err != nil {
			return nil, err
		}
	}

	milestone, err := l.svcCtx.MilestoneModel.FindOne(l.ctx, req.Milestoneid)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get milestone db err. rowType: %s ,err : %v", "", err)
	}
	milestoneinfo, err := l.NewMilestoneInfo(milestone, req)

	whereBuilde := l.svcCtx.IssueModel.RowBuilder()
	issueList, err := l.svcCtx.IssueModel.FindListByMilestoneid(l.ctx, whereBuilde, req.Milestoneid, "id DESC")
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Milestoneid", req.Milestoneid)
	}

	var issueInfoList []types.IssueInfo
	if len(issueList) > 0 {
		issueLogic := issue.NewIssueLogic(l.ctx, l.svcCtx)
		issueReq := types.IssueReq{}
		for _, e := range issueList {
			issue, err := issueLogic.NewIssueInfo(e, &issueReq)
			if err != nil {
				return nil, err
			}

			issueInfoList = append(issueInfoList, *issue)
		}

	}

	return &types.MilestoneSignalResp{
		Milestone: *milestoneinfo,
		IssueList: issueInfoList,
	}, nil
}

func (l *MilestonesignalLogic) updataMilestone(req *types.MilestoneSignalReq) error {
	milestone, err := l.svcCtx.MilestoneModel.FindOne(l.ctx, req.Milestoneid)
	if err != nil && err != model.ErrNotFound {
		return err
	}

	l.svcCtx.MilestoneModel.Update(l.ctx, nil, milestone)
	return nil
}

func (l *MilestonesignalLogic) deleteMilestone(req *types.MilestoneSignalReq) error {
	err := l.svcCtx.MilestoneModel.Delete(l.ctx, nil, req.Milestoneid)
	if err != nil && err != model.ErrNotFound {
		return err
	}

	return nil
}

func (l *MilestonesignalLogic) NewMilestoneInfo(milestone *model.Milestone, req *types.MilestoneSignalReq) (*types.MilestoneInfo, error) {
	var milestoneinfo types.MilestoneInfo
	milestoneinfo.Milestoneid = milestone.Id
	milestoneinfo.Name = milestone.Name
	milestoneinfo.UpdateTime = milestone.UpdateTime.String()
	return &milestoneinfo, nil
}
