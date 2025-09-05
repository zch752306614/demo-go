package logic

import (
	"context"

	"demo/internal/svc"
	"demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// DemoLogic 演示逻辑层（等价于 Java 的 Service 层）
// - 承担业务处理与领域逻辑
// - 通过 svcCtx 访问共享资源（如 DB、缓存等）
type DemoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDemoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DemoLogic {
	return &DemoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Demo 返回简单问候语
func (l *DemoLogic) Demo(req *types.Request) (resp *types.Response, err error) {
	return &types.Response{Message: "hello " + req.Name}, nil
}
