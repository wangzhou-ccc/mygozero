package logic

import (
	"context"

	"github.com/wangzhou-ccc/mygozero/apps/product/internal/svc"
	"github.com/wangzhou-ccc/mygozero/apps/product/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type RollbackProductStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRollbackProductStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RollbackProductStockLogic {
	return &RollbackProductStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RollbackProductStockLogic) RollbackProductStock(in *product.UpdateProductStockRequest) (*product.UpdateProductStockResponse, error) {
	// todo: add your logic here and delete this line

	return &product.UpdateProductStockResponse{}, nil
}
