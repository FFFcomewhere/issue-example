package other

import (
	"context"
	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/logic/issue"
	"github.com/FFFcomewhere/issue-example/app/issuecenter/model"
	"github.com/FFFcomewhere/issue-example/common/xerr"
	es "github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"reflect"
	"strconv"

	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/svc"
	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IssuesearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIssuesearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IssuesearchLogic {
	return &IssuesearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IssuesearchLogic) Issuesearch(req *types.IssueSearchReq) (resp *types.IssueSearchResp, err error) {
	var issues []model.Issue
	//在issue.name中查询关键词
	q := es.NewQueryStringQuery("name:" + req.Keyword)
	res, err := l.svcCtx.ES.Search("issue").Query(q).Do(l.ctx)
	if err != nil {
		return nil, err
	}
	issues = parseIssue(res)

	//在tag.name中查询关键词
	q = es.NewQueryStringQuery("name:" + req.Keyword)
	res, err = l.svcCtx.ES.Search("tag").Query(q).Do(l.ctx)
	if err != nil {
		return nil, err
	}
	tags := parseTag(res)

	for i, _ := range tags {
		whereBuilde := l.svcCtx.IssueModel.RowBuilder()
		issueList, err := l.svcCtx.IssueModel.FindListByTagid(l.ctx, whereBuilde, tags[i].Id, "id DESC")
		if err != nil && err != model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "search issue db err. rowType: %s ,err : %v", "tag", err)
		}

		for j, _ := range issueList {
			issues = append(issues, *issueList[j])
		}
	}

	//在comment.content中查询关键词
	q = es.NewQueryStringQuery("content:" + req.Keyword)
	res, err = l.svcCtx.ES.Search("comment").Query(q).Do(l.ctx)
	if err != nil {
		return nil, err
	}
	comments := parseComment(res)

	for i, _ := range comments {
		whereBuilder := l.svcCtx.IssueModel.RowBuilder().Where("id = ?", strconv.Itoa(int(comments[i].Issueid)))
		issue, err := l.svcCtx.IssueModel.FindOneByQuery(l.ctx, whereBuilder)
		if err != nil && err != model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "search issue db err. rowType: %s ,err : %v", "tag", err)
		}
		issues = append(issues, *issue)
	}

	//封装响应
	var issueList []types.IssueInfo
	issueLogic := issue.NewIssueLogic(l.ctx, l.svcCtx)

	for i, _ := range issues {
		issueinfo, err := issueLogic.NewIssueInfo(&issues[i], &types.IssueReq{})
		if err != nil {
			return nil, err
		}
		issueList = append(issueList, *issueinfo)
	}

	return &types.IssueSearchResp{
		List: issueList,
	}, nil
}

//解析数据 issue
func parseIssue(res *es.SearchResult) []model.Issue {
	var issue model.Issue
	var issues []model.Issue

	for _, item := range res.Each(reflect.TypeOf(issue)) { //从搜索结果中取数据的方法
		issues = append(issues, item.(model.Issue))
	}

	return issues
}

//解析数据 tag
func parseTag(res *es.SearchResult) []model.Tag {
	var tag model.Tag
	var tags []model.Tag

	for _, item := range res.Each(reflect.TypeOf(tag)) { //从搜索结果中取数据的方法
		tags = append(tags, item.(model.Tag))
	}

	return tags
}

//解析数据 comment
func parseComment(res *es.SearchResult) []model.Comment {
	var comment model.Comment
	var comments []model.Comment

	for _, item := range res.Each(reflect.TypeOf(comment)) { //从搜索结果中取数据的方法
		comments = append(comments, item.(model.Comment))
	}

	return comments
}
