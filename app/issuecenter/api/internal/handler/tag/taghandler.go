package tag

import (
	"net/http"

	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/logic/tag"
	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/svc"
	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func TagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TagReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := tag.NewTagLogic(r.Context(), svcCtx)
		resp, err := l.Tag(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
