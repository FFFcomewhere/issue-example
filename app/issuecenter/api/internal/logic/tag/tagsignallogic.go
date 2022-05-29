package tag

import (
	"context"
	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/logic/issue"
	"github.com/FFFcomewhere/issue-example/app/issuecenter/model"
	"github.com/FFFcomewhere/issue-example/common/xerr"
	"github.com/pkg/errors"

	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/svc"
	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrTagNoExistsError = xerr.NewErrMsg("标签不存在")

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
	//修改Tag基本信息
	if req.ReName != "" {
		err := l.updataTag(req)
		if err != nil {
			return nil, err
		}
	}

	//删除Tag
	if req.IfDelete == true {
		err := l.deleteTag(req)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}

	tag, err := l.svcCtx.TagModel.FindOneByName(l.ctx, req.Name)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get Tag db err. rowType: %s ,err : %v", "", err)
	}

	Tag, err := l.svcCtx.TagModel.FindOne(l.ctx, tag.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get Tag db err. rowType: %s ,err : %v", "", err)
	}
	Taginfo, err := l.NewTagInfo(Tag, req)

	whereBuilde := l.svcCtx.IssueModel.RowBuilder()
	issueList, err := l.svcCtx.IssueModel.FindListByTagid(l.ctx, whereBuilde, tag.Id, "id DESC")
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Tagid", tag.Id)
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

	return &types.TagSignalResp{
		Tag:       *Taginfo,
		IssueList: issueInfoList,
	}, nil
}

func (l *TagsignalLogic) updataTag(req *types.TagSignalReq) error {
	Tag, err := l.svcCtx.TagModel.FindOneByName(l.ctx, req.Name)
	if err != nil && err != model.ErrNotFound {
		return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get issue db err. rowType: %s ,err : %v", "issueid", err)
	}

	if Tag == nil {
		return errors.Wrapf(ErrTagNoExistsError, "tagName:%v", req.Name)
	}

	Tag.Name = req.ReName
	l.svcCtx.TagModel.Update(l.ctx, nil, Tag)

	return nil
}

func (l *TagsignalLogic) deleteTag(req *types.TagSignalReq) error {
	tag, err := l.svcCtx.TagModel.FindOneByName(l.ctx, req.Name)
	if err != nil && err != model.ErrNotFound {
		return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get Tag db err. rowType: %s ,err : %v", "", err)
	}

	err = l.svcCtx.TagModel.Delete(l.ctx, nil, tag.Id)
	if err != nil && err != model.ErrNotFound {
		return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get Tag db err. rowType: %s ,err : %v", "", err)
	}

	return nil
}

func (l *TagsignalLogic) NewTagInfo(Tag *model.Tag, req *types.TagSignalReq) (*types.TagInfo, error) {
	var Taginfo types.TagInfo
	Taginfo.Tagid = Tag.Id
	Taginfo.Name = Tag.Name
	return &Taginfo, nil
}
