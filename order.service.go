package gofo

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hiscaler/gofo-go/entity"
	"gopkg.in/guregu/null.v4"
)

// 订单服务
type orderService service

// OrderBox 订单箱子
type OrderBox struct {
	No     int           `json:"no"`     // 箱子编号
	Length float64       `json:"length"` // 箱子长度
	Width  float64       `json:"width"`  // 箱子宽度
	Height float64       `json:"height"` // 箱子高度
	Weight float64       `json:"weight"` // 箱子重量
	Skus   []OrderBoxSku `json:"skus"`   // SKU 列表
}

type OrderBoxSku struct {
	SKU         string `json:"sku"`         // SKU 编码
	ChineseName string `json:"chineseName"` // 中文品名
	EnglishName string `json:"englishName"` // 英文品名
	Quantity    int    `json:"quantity"`    // 数量
}

func (m OrderBox) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Length,
			validation.Required.Error("长不能为空"),
			validation.Min(0.01).Error("长不能小于 {{.min}}"),
			validation.Max(999999.99).Error("长不能大于 {{.max}}"),
		),
		validation.Field(&m.Width,
			validation.Required.Error("宽不能为空"),
			validation.Min(0.01).Error("宽不能小于 {{.min}}"),
			validation.Max(999999.99).Error("宽不能大于 {{.max}}"),
		),
		validation.Field(&m.Height,
			validation.Required.Error("高不能为空"),
			validation.Min(0.01).Error("高不能小于 {{.min}}"),
			validation.Max(999999.99).Error("高不能大于 {{.max}}"),
		),
		validation.Field(&m.Weight,
			validation.Required.Error("重量不能为空"),
			validation.Min(0.01).Error("重量不能小于 {{.min}}"),
			validation.Max(999999.99).Error("重量不能大于 {{.max}}"),
		),
	)
}

type CreateOrderRequest struct {
	COrderNo    null.String `json:"cOrderNo,omitempty"`    // 客户单号(长度 1-30)
	ReferenceNo null.String `json:"referenceNo,omitempty"` // 参考单号(长度 1-30)
	Reference   null.String `json:"reference"`             // 预留字段（长度 1-255），应用在面单下方，可存放 sku 信息
}

func (m CreateOrderRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.COrderNo, validation.When(m.COrderNo.Valid, validation.Length(1, 30).Error("客户单号不能超过 {{.max}} 个字符"))),
		validation.Field(&m.ReferenceNo, validation.When(m.ReferenceNo.Valid, validation.Length(1, 30).Error("参考单号不能超过 {{.max}} 个字符"))),
		validation.Field(&m.Reference, validation.When(m.Reference.Valid, validation.Length(1, 255).Error("预留字段 {{.max}} 个字符"))),
	)
}

type CreateOrderResult struct {
	CustomerNo string         `json:"customerNo"` // 客户订单号
	Orders     []entity.Order `json:"orders"`     // 订单列表
}

// Create 创建订单
func (s orderService) Create(ctx context.Context, requests []CreateOrderRequest) ([]CreateOrderResult, error) {
	for _, req := range requests {
		if err := req.Validate(); err != nil {
			return nil, invalidInput(err)
		}
	}

	var res []CreateOrderResult
	resp, err := s.httpClient.R().
		SetContext(ctx).
		SetBody(requests).
		SetResult(&res).
		Post("/open-api/v2/order/create")
	if err = recheckError(resp, err); err != nil {
		return nil, err
	}
	return res, nil
}

type OrderQueryRequest struct {
	CustomerNos string `json:"customerNos"` // 订单号
}

func (m OrderQueryRequest) validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.CustomerNos, validation.Required.Error("订单号不能为空")),
	)
}

// Query 根据查询条件筛选符合条件的订单列表数据
func (s orderService) Query(ctx context.Context, req OrderQueryRequest) ([]entity.Order, error) {
	if err := req.validate(); err != nil {
		return nil, invalidInput(err)
	}

	var res []entity.Order
	resp, err := s.httpClient.R().
		SetContext(ctx).
		SetQueryParams(map[string]string{
			"customerNos": req.CustomerNos,
		}).
		SetResult(&res).
		Get("/external/orders")
	if err != nil {
		return nil, err
	}

	if err = recheckError(resp, err); err != nil {
		return nil, err
	}
	return res, nil
}

type CancelOrderRequest struct {
	OrderNos string `json:"orderNos"` // 订单号
}

func (m CancelOrderRequest) validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.OrderNos, validation.Required.Error("订单号不能为空")),
	)
}

type OrderCancelResult struct {
	OrderNo    string      `json:"orderNo"`    // 订单号
	FailReason null.String `json:"failReason"` // 失败原因
}

// Cancel 取消订单
func (s orderService) Cancel(ctx context.Context, req CancelOrderRequest) ([]OrderCancelResult, error) {
	if err := req.validate(); err != nil {
		return nil, invalidInput(err)
	}

	var res []OrderCancelResult
	resp, err := s.httpClient.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(&res).
		Delete("/external/orders")
	if err = recheckError(resp, err); err != nil {
		return nil, err
	}
	return res, nil
}
