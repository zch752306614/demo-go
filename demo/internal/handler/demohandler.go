package handler

import (
	"net/http"

	"demo/internal/logic"
	"demo/internal/svc"
	"demo/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// DemoHandler 演示：从路径中解析参数，并返回 JSON
// - httpx.Parse 会根据 types.Request 的 tag 自动从 path/query/body 解析
// - OkJsonCtx 以标准 JSON 写回（带 context）
func DemoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDemoLogic(r.Context(), svcCtx)
		resp, err := l.Demo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
