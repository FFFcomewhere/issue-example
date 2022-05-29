package svc

import (
	"github.com/FFFcomewhere/issue-example/app/usercenter/api/internal/config"
	"github.com/FFFcomewhere/issue-example/app/usercenter/rpc/user"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	UsercenterRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UsercenterRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
