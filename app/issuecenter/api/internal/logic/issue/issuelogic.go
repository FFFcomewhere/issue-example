package issue

import (
	"context"
	"github.com/FFFcomewhere/issue-example/app/issuecenter/model"
	"github.com/FFFcomewhere/issue-example/app/usercenter/rpc/types/builder"
	"github.com/FFFcomewhere/issue-example/common/xerr"
	"github.com/pkg/errors"

	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/svc"
	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrNewIssueInfoError = xerr.NewErrMsg("new issue info err ")

type IssueLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIssueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IssueLogic {
	return &IssueLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IssueLogic) Issue(req *types.IssueReq) (resp *types.IssueResp, err error) {
	page := req.Page
	if page < 1 {
		page = 1
	}

	whereBuilder := l.svcCtx.IssueModel.RowBuilder()
	issueList, err := l.svcCtx.IssueModel.FindPageListByIdASC(l.ctx, whereBuilder, (req.Page-1)*req.PageSize, req.PageSize)

	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get IssueList db err. rowType: %s ,err : %v", "issueid", err)
	}

	var issueInfoList []types.IssueInfo
	if len(issueList) > 0 {
		for _, e := range issueList {
			issue, err := l.NewIssueInfo(e, req)
			if err != nil {
				return nil, err
			}

			issueInfoList = append(issueInfoList, *issue)
		}
	}

	return &types.IssueResp{
		List: issueInfoList,
	}, nil
}

func (l *IssueLogic) NewIssueInfo(issue *model.Issue, req *types.IssueReq) (*types.IssueInfo, error) {
	var tempIssue types.IssueInfo
	tempIssue.Issueid = issue.Id
	tempIssue.IssueName = issue.Name
	tempIssue.UpdateTime = issue.UpdateTime.String()

	//获取用户信息
	tempGetUserInfoResp, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &builder.GetUserInfoReq{
		Userid:   issue.Userid,
		Username: "",
	})
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "newIssueInfo. user db err. rowType: %s ,err : %v", "issueid", err)
	}
	tempIssue.UserName = tempGetUserInfoResp.Username

	//获取tag信息
	tempTag, err := l.svcCtx.TagModel.FindOne(l.ctx, issue.Tagid)
	if err != nil && err != model.ErrNotFound {
		return nil, err
	}
	if tempTag == nil {
		tempIssue.TagName = ""
	} else {
		tempIssue.TagName = tempTag.Name
	}

	//获取milestone信息
	tempMilestone, err := l.svcCtx.MilestoneModel.FindOne(l.ctx, issue.Milestoneid)
	if err != nil && err != model.ErrNotFound {
		return nil, err
	}
	if tempMilestone == nil {
		tempIssue.MilestoneName = ""
	} else {
		tempIssue.MilestoneName = tempMilestone.Name
	}

	return &tempIssue, nil
}
