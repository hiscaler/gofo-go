package gofo

import (
	"context"
	"fmt"

	"github.com/hiscaler/gofo-go/entity"
)

// 订单服务
type orderService service

// Create 创建订单
// doc: https://www.showdoc.com.cn/gofo/9294126170337714
func (s orderService) Create(ctx context.Context, req entity.CreateOrderRequest) (*entity.OrderCreateResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, invalidInput(err)
	}

	var res struct {
		NormalResponse
		Data entity.OrderCreateResponse `json:"data"`
	}
	resp, err := s.httpClient.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(&res).
		Post("/open-api/v2/order/create")
	if err = recheckError(resp, err); err != nil {
		return nil, err
	}
	return &res.Data, nil
}

// Cancel 取消订单
// doc: https://www.showdoc.com.cn/gofo/9294126170337715
func (s orderService) Cancel(ctx context.Context, req entity.CancelOrderRequest) (bool, error) {
	if err := req.Validate(); err != nil {
		return false, invalidInput(err)
	}

	var res NormalResponse
	resp, err := s.httpClient.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(&res).
		Post("/open-api/v2/order/cancel")
	if err = recheckError(resp, err); err != nil {
		return false, err
	}
	return true, nil
}

// ShippingLabel 获取面单
// @param orderNo 订单号/运单号/客户单号
func (s orderService) ShippingLabel(ctx context.Context, orderNo string) (string, error) {
	var res struct {
		NormalResponse
		Data struct {
			Base64code string `json:"base64code"`
		} `json:"data"`
	}
	resp, err := s.httpClient.R().
		SetContext(ctx).
		SetQueryParam("orderNo", orderNo).
		SetResult(&res).
		Get("/open-api/v2/order/getOrderLabelUrlV2")
	if err = recheckError(resp, err); err != nil {
		return "", err
	}
	return res.Data.Base64code, nil
}

// Track 轨迹查询
// @param orderNo 订单号/运单号/客户单号
func (s orderService) Track(ctx context.Context, orderNo string) ([]entity.TrackEvent, error) {
	var res struct {
		NormalResponse
		Data []entity.TrackEvent `json:"data"`
	}
	resp, err := s.httpClient.R().
		SetContext(ctx).
		SetResult(&res).
		Get(fmt.Sprintf("/open-api/v2/order/track/%s", orderNo))
	if err = recheckError(resp, err); err != nil {
		return nil, err
	}
	return res.Data, nil
}
