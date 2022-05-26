package svc

import (
	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/config"
	"github.com/FFFcomewhere/issue-example/app/issuecenter/model"
	"github.com/FFFcomewhere/issue-example/app/usercenter/rpc/user"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	UserRpc user.User

	CommentModel   model.CommentModel
	IssueModel     model.IssueModel
	MilestoneModel model.MilestoneModel
	TagModel       model.TagModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config: c,

		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),

		CommentModel:   model.NewCommentModel(sqlConn, c.Cache),
		IssueModel:     model.NewIssueModel(sqlConn, c.Cache),
		MilestoneModel: model.NewMilestoneModel(sqlConn, c.Cache),
		TagModel:       model.NewTagModel(sqlConn, c.Cache),
	}
}
