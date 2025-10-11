package gofo

import (
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/hiscaler/gofo-go/config"
)

type service struct {
	config     *config.Config // Config
	logger     *log.Logger    // Logger
	httpClient *resty.Client  // HTTP client
}

// API Services
type services struct {
	Order orderService // 订单服务
	//Product         productService         // 产品服务
	//ScanForm        scanFormService        // 交运单服务
	//ShippingAddress shippingAddressService // 发货地址服务
}
