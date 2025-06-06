package logic

import (
	"context"

	"gozero/app/rbac/api/internal/svc"
	"gozero/app/rbac/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RbacLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRbacLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RbacLogic {
	return &RbacLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RbacLogic) Rbac(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
