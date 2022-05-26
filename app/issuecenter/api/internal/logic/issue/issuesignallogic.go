package issue

import (
	"context"
	"github.com/FFFcomewhere/issue-example/app/issuecenter/model"
	"github.com/FFFcomewhere/issue-example/app/usercenter/rpc/types/builder"
	"github.com/FFFcomewhere/issue-example/common/xerr"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/svc"
	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrUserNoExistsError = xerr.NewErrMsg("用户不存在")
var ErrIssueUpdataBaseMessageError = xerr.NewErrMsg("issue updata base message faild")
var ErrIssueDeleteError = xerr.NewErrMsg("issue delete faild")
var ErrCommentAddError = xerr.NewErrMsg("comment add faild")
var ErrCommentUpdateError = xerr.NewErrMsg("comment update faild")
var ErrCommentDeleteError = xerr.NewErrMsg("comment Delete faild")
var FindCommentListByIssueid = xerr.NewErrMsg("issue find comment list by issueid faild")
var ErrNewCommentInfoError = xerr.NewErrMsg("new commentinfo faild")

type IssuesignalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIssuesignalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IssuesignalLogic {
	return &IssuesignalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IssuesignalLogic) Issuesignal(req *types.IssueSignalReq) (resp *types.IssueSignalResp, err error) {
	//修改issue基本信息
	if req.ReName != "" || req.ReTag != "" || req.ReMilestone != "" {
		err := l.updataIssue(req)
		if err != nil {
			return nil, err
		}
	}

	//删除提案
	if req.IfDelete == true {
		err := l.deleteIssue(req)
		if err != nil {
			return nil, err
		}
	}

	//添加评论
	if req.AddComment != "" {
		err := l.addComment(req)
		if err != nil {
			return nil, err
		}
	}

	//删除评论
	if req.DeleteCommentid != 0 {
		err := l.deleteComment(req)
		if err != nil {
			return nil, err
		}
	}

	//修改评论
	if req.UpdateCommentid != 0 {
		err := l.deleteComment(req)
		if err != nil {
			return nil, err
		}
	}

	//封装响应
	var tempResp types.IssueSignalResp

	issue, err := l.svcCtx.IssueModel.FindOne(l.ctx, req.Issueid)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "issue_id:%s,err:%v", req.Issueid, err)
	}
	tempIssueInfo, err := l.newIssueInfo(issue, req)
	if err != nil {
		return nil, err
	}
	tempResp.Issue = *tempIssueInfo

	whereBuilder := l.svcCtx.CommentModel.RowBuilder().Where(squirrel.Eq{})
	tempCommentList, err := l.svcCtx.CommentModel.FindListByIssueid(l.ctx, whereBuilder, req.Issueid, "id desc")
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Issueid", req.Issueid)
	}

	var tempCommentInfoList []types.CommentInfo
	if len(tempCommentList) > 0 {
		for _, e := range tempCommentList {
			tempCommentInfo, err := l.newCommentInfo(e, req)
			if err != nil {
				return nil, err
			}
			tempCommentInfoList = append(tempCommentInfoList, *tempCommentInfo)
		}
	}

	return &types.IssueSignalResp{
		Issue:       *tempIssueInfo,
		CommentList: tempCommentInfoList,
	}, nil
}

//修改issue基本信息
func (l *IssuesignalLogic) updataIssue(req *types.IssueSignalReq) error {
	newIssue, err := l.svcCtx.IssueModel.FindOne(l.ctx, req.Issueid)
	if err != nil && err != model.ErrNotFound {
		return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get issue db err. rowType: %s ,err : %v", "issueid", err)
	}

	if newIssue == nil {
		return errors.Wrapf(ErrUserNoExistsError, "issueid:%d", req.Issueid)
	}

	if req.ReName != "" {
		newIssue.Name = req.ReName
	}

	if req.ReTag != "" {
		tempTag, err := l.svcCtx.TagModel.FindOneByName(l.ctx, req.ReTag)
		if err == nil {
			newIssue.Tagid = tempTag.Id
		}
	}

	if req.ReMilestone != "" {
		tempTag, err := l.svcCtx.MilestoneModel.FindOneByName(l.ctx, req.ReMilestone)
		if err == nil {
			newIssue.Milestoneid = tempTag.Id
		}
	}

	l.svcCtx.IssueModel.Update(l.ctx, nil, newIssue)

	return nil
}

//删除提案
func (l *IssuesignalLogic) deleteIssue(req *types.IssueSignalReq) error {
	err := l.svcCtx.IssueModel.Delete(l.ctx, nil, req.Issueid)

	if err != nil && err != model.ErrNotFound {
		return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "delete issue db err. rowType: %s ,err : %v", "issueid", err)
	}
	return nil
}

//添加评论
func (l *IssuesignalLogic) addComment(req *types.IssueSignalReq) error {
	if err := l.svcCtx.CommentModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		newComment := model.Comment{
			Issueid: req.Issueid,
			Content: req.AddComment,
		}

		tempIssue, err := l.svcCtx.IssueModel.FindOne(l.ctx, req.Issueid)
		if err != nil && err != model.ErrNotFound {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "new comment db err. rowType: %s ,err : %v", "issue", err)
		}
		if tempIssue != nil {
			newComment.Userid = tempIssue.Userid
		} else {
			newComment.Userid = 0
		}

		_, err = l.svcCtx.CommentModel.Insert(context, session, &newComment)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "new db comment Insert err:%v,comment:%+v", err, newComment)
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

//删除评论
func (l *IssuesignalLogic) deleteComment(req *types.IssueSignalReq) error {
	err := l.svcCtx.CommentModel.Delete(l.ctx, nil, req.DeleteCommentid)
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "new comment db err. rowType: %s ,err : %v", "issue", err)
	}
	return nil
}

//修改评论
func (l *IssuesignalLogic) updateComment(req *types.IssueSignalReq) error {
	newComment, err := l.svcCtx.CommentModel.FindOne(l.ctx, req.UpdateCommentid)
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "new comment db err. rowType: %s ,err : %v", "issue", err)
	}

	newComment.Content = req.UpdateComment
	l.svcCtx.CommentModel.Update(l.ctx, nil, newComment)

	return nil
}

func (l *IssuesignalLogic) newIssueInfo(issue *model.Issue, req *types.IssueSignalReq) (*types.IssueInfo, error) {
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
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "new comment db err. rowType: %s ,err : %v", "issue", err)
	}
	tempIssue.UserName = tempGetUserInfoResp.Username

	//获取tag信息
	tempTag, err := l.svcCtx.TagModel.FindOne(l.ctx, issue.Tagid)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "new comment db err. rowType: %s ,err : %v", "issue", err)
	}
	if tempTag == nil {
		tempIssue.TagName = ""
	} else {
		tempIssue.TagName = tempTag.Name
	}

	//获取milestone信息
	tempMilestone, err := l.svcCtx.MilestoneModel.FindOne(l.ctx, issue.Milestoneid)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "new comment db err. rowType: %s ,err : %v", "issue", err)
	}
	if tempMilestone == nil {
		tempIssue.MilestoneName = ""
	} else {
		tempIssue.MilestoneName = tempMilestone.Name
	}

	return &tempIssue, nil
}

func (l *IssuesignalLogic) newCommentInfo(comment *model.Comment, req *types.IssueSignalReq) (*types.CommentInfo, error) {
	var tempCommentInfo types.CommentInfo

	//获取用户名
	tempGetUserInfoResp, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &builder.GetUserInfoReq{
		Userid:   comment.Userid,
		Mobile:   "",
		Username: "",
	})
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "new comment db err. rowType: %s ,err : %v", "issue", err)
	}
	tempCommentInfo.UserName = tempGetUserInfoResp.Username

	tempCommentInfo.Commentid = comment.Id
	tempCommentInfo.Content = comment.Content
	tempCommentInfo.UpdateTime = comment.UpdateTime.String()

	return &tempCommentInfo, nil
}
