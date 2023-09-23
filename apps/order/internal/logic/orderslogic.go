package logic

import (
	"context"

	"github.com/wangzhou-ccc/mygozero/apps/order/internal/svc"
	"github.com/wangzhou-ccc/mygozero/apps/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrdersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrdersLogic {
	return &OrdersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrdersLogic) Orders(in *order.OrdersRequest) (*order.OrdersResponse, error) {
	// todo: add your logic here and delete this line

	return &order.OrdersResponse{}, nil
}
