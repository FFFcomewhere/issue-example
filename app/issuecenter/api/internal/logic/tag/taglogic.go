package tag

import (
	"context"
	"github.com/FFFcomewhere/issue-example/app/issuecenter/model"
	"github.com/FFFcomewhere/issue-example/common/xerr"
	"github.com/pkg/errors"

	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/svc"
	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TagLogic {
	return &TagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TagLogic) Tag(req *types.TagReq) (resp *types.TagResp, err error) {

	page := req.Page
	if page < 1 {
		page = 1
	}

	whereBuilder := l.svcCtx.TagModel.RowBuilder()
	TagList, err := l.svcCtx.TagModel.FindPageListByIdASC(l.ctx, whereBuilder, (req.Page-1)*req.PageSize, req.PageSize)

	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get TagList db err. rowType: %s ,err : %v", "tag", err)
	}

	var TagInfoList []types.TagInfo
	if len(TagList) > 0 {
		for _, e := range TagList {
			Tag, err := l.newTagInfo(e, req)
			if err != nil {
				return nil, err
			}

			TagInfoList = append(TagInfoList, *Tag)
		}
	}

	return &types.TagResp{
		List: TagInfoList,
	}, nil

}

func (l *TagLogic) newTagInfo(Tag *model.Tag, req *types.TagReq) (*types.TagInfo, error) {
	var tempTag types.TagInfo
	tempTag.Name = Tag.Name
	tempTag.Tagid = Tag.Id

	return &tempTag, nil
}
