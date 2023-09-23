package user

import (
	"context"

	"github.com/wangzhou-ccc/mygozero/apps/app/api/internal/svc"
	"github.com/wangzhou-ccc/mygozero/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCollectionAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCollectionAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCollectionAddLogic {
	return &UserCollectionAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCollectionAddLogic) UserCollectionAdd(req *types.UserCollectionAddReq) (resp *types.UserCollectionAddRes, err error) {
	// todo: add your logic here and delete this line

	return
}
