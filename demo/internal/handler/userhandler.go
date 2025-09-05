package handler

import (
	"net/http"

	"demo/internal/logic"
	"demo/internal/svc"
	"demo/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// CreateUserHandler 解析创建请求并调用业务逻辑
// - httpx.Parse 会从 JSON body 解析到 CreateUserRequest
// - 统一使用 httpx.ErrorCtx 返回错误，OkJsonCtx 返回成功 JSON
func CreateUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewUserLogic(r.Context(), svcCtx)
		resp, err := l.CreateUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}

// GetUserHandler 从路径读取 id 并查询
func GetUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdPath
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewUserLogic(r.Context(), svcCtx)
		resp, err := l.GetUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}

// ListUsersHandler 返回用户列表
func ListUsersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserLogic(r.Context(), svcCtx)
		resp, err := l.ListUsers()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}

// UpdateUserHandler 组合路径 id 与 JSON body 完成更新
func UpdateUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var path types.IdPath
		if err := httpx.ParsePath(r, &path); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		var body types.UpdateUserRequest
		if err := httpx.ParseJsonBody(r, &body); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		body.ID = path.ID
		l := logic.NewUserLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUser(&body)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}

// DeleteUserHandler 删除指定 id 的用户
func DeleteUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdPath
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewUserLogic(r.Context(), svcCtx)
		if err := l.DeleteUser(&req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		httpx.Ok(w)
	}
}
