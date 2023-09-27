package logic

import (
	"context"

	"github.com/wangzhou-ccc/mygozero/apps/product/internal/svc"
	"github.com/wangzhou-ccc/mygozero/apps/product/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductsLogic {
	return &ProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductsLogic) Products(in *product.ProductRequest) (*product.ProductResponse, error) {
	// todo: add your logic here and delete this line

	return &product.ProductResponse{}, nil
}
