package issue

import (
	"net/http"

	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/logic/issue"
	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/svc"
	"github.com/FFFcomewhere/issue-example/app/issuecenter/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func IssuesignalHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IssueSignalReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := issue.NewIssuesignalLogic(r.Context(), svcCtx)
		resp, err := l.Issuesignal(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
