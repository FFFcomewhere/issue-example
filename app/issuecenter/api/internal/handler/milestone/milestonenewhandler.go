package milestone

import (
	"net/http"

	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/logic/milestone"
	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/svc"
	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func MilestonenewHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MilestoneNewReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := milestone.NewMilestonenewLogic(r.Context(), svcCtx)
		resp, err := l.Milestonenew(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
