package svc

import (
	"github.com/wangzhou-ccc/mygozero/apps/app/api/internal/config"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/wangzhou-ccc/mygozero/apps/order/orderclient"
	"github.com/wangzhou-ccc/mygozero/apps/product/productclient"
)

type ServiceContext struct {
	Config     config.Config
	OrderRPC   orderclient.Order
	ProductRPC productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRPC:   orderclient.NewOrder(zrpc.MustNewClient(c.OrderRPC)),
		ProductRPC: productclient.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
	}
}
