package entity

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gopkg.in/guregu/null.v4"
)

// OrderCreateResponse 订单
type OrderCreateResponse struct {
	COrderNo        string `json:"cOrderNo"`
	VerificationPin string `json:"verificationPin"`
	Type            string `json:"type"`
	WaybillNo       string `json:"waybillNo"`
}

// CreateOrderRequest 创建订单请求
type CreateOrderRequest struct {
	COrderNo              null.String     `json:"cOrderNo,omitempty"`              // 客户单号(长度 1-30)
	ReferenceNo           null.String     `json:"referenceNo,omitempty"`           // 参考单号(长度 1-30)
	Reference4            null.String     `json:"reference4,omitempty"`            // 预留字段(长度 1-255)，应用在面单下方，可存放 sku 信息
	YtReference           null.String     `json:"ytReference,omitempty"`           // 面单 Reference 栏位显示内容(长度 1-30)
	ShippingType          null.String     `json:"shippingType,omitempty"`          // 配送类型: HDN(送货上门), ZT(自提), 默认为 HDN(送货上门)
	ProductCode           null.String     `json:"productCode,omitempty"`           // 产品编码(长度 1-100), 非全境可不传
	DeclaredValue         float64         `json:"declaredValue"`                   // 包裹预报货值, 单位: 美金, 范围 0.0001-100.00
	QueryCollectStartTime null.String     `json:"queryCollectStartTime,omitempty"` // 揽收开始时间, 格式: yyyy-MM-dd HH:mm:ss
	QueryCollectEndTime   null.String     `json:"queryCollectEndTime,omitempty"`   // 揽收结束时间, 格式: yyyy-MM-dd HH:mm:ss
	OrderShipper          OrderShipper    `json:"orderShipper"`                    // 寄件信息
	OrderConsignee        OrderConsignee  `json:"orderConsignee"`                  // 收件信息
	OrderGoods            OrderGoods      `json:"orderGoods"`                      // 订单货物规格
	OrderItemList         []OrderItem     `json:"orderItemList"`                   // 订单物品信息
	OrderInsurance        *OrderInsurance `json:"orderInsurance,omitempty"`        // 订单保价
}

func (m CreateOrderRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.COrderNo, validation.When(m.COrderNo.Valid, validation.Length(1, 30).Error("客户单号长度必须在 {{.min}}-{{.max}} 之间"))),
		validation.Field(&m.ReferenceNo, validation.When(m.ReferenceNo.Valid, validation.Length(1, 30).Error("参考单号长度必须在 {{.min}}-{{.max}} 之间"))),
		validation.Field(&m.Reference4, validation.When(m.Reference4.Valid, validation.Length(1, 255).Error("预留字段长度必须在 {{.min}}-{{.max}} 之间"))),
		validation.Field(&m.YtReference, validation.When(m.YtReference.Valid, validation.Length(1, 30).Error("面单 Reference 栏位内容长度必须在 {{.min}}-{{.max}} 之间"))),
		validation.Field(&m.ProductCode, validation.When(m.ProductCode.Valid, validation.Length(1, 100).Error("产品编码长度必须在 {{.min}}-{{.max}} 之间"))),
		validation.Field(&m.DeclaredValue, validation.Required.Error("包裹预报货值不能为空"), validation.Min(0.0001).Error("包裹预报货值不能小于 {{.min}}"), validation.Max(100.00).Error("包裹预报货值不能大于 {{.max}}")),
		validation.Field(&m.OrderShipper),
		validation.Field(&m.OrderConsignee),
		validation.Field(&m.OrderGoods),
		validation.Field(&m.OrderItemList, validation.Required.Error("订单物品信息不能为空")),
	)
}

// OrderShipper 发件人信息
type OrderShipper struct {
	ShipperName    string      `json:"shipperName"`            // 发件人-姓名, 长度为 1-50
	ShipperPhone   string      `json:"shipperPhone"`           // 发件人-手机号, 长度为 10-14
	ShipperCountry string      `json:"shipperCountry"`         // 发件人-国家
	ShipperState   string      `json:"shipperState"`           // 发件人-省/州, 长度 1-35
	ShipperCity    string      `json:"shipperCity"`            // 发件人-市, 长度 1-50
	ShipperArea    null.String `json:"shipperArea,omitempty"`  // 发件人-区, 长度 1-50
	ShipperStreet  string      `json:"shipperStreet"`          // 发件人-详细地址, 长度 1-100
	ShipperCode    string      `json:"shipperCode"`            // 发件人-邮编
	ShipperEmail   null.String `json:"shipperEmail,omitempty"` // 发件人-邮箱, 长度为 1-100
}

func (m OrderShipper) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.ShipperName, validation.Required.Error("发件人姓名不能为空"), validation.Length(1, 50).Error("发件人姓名长度必须在 {{.min}}-{{.max}} 之间")),
		validation.Field(&m.ShipperPhone, validation.Required.Error("发件人手机号不能为空"), validation.Length(10, 14).Error("发件人手机号长度必须在 {{.min}}-{{.max}} 之间")),
		validation.Field(&m.ShipperCountry, validation.Required.Error("发件人国家不能为空")),
		validation.Field(&m.ShipperState, validation.Required.Error("发件人省/州不能为空"), validation.Length(1, 35).Error("发件人省/州长度必须在 {{.min}}-{{.max}} 之间")),
		validation.Field(&m.ShipperCity, validation.Required.Error("发件人市不能为空"), validation.Length(1, 50).Error("发件人城市长度必须在 {{.min}}-{{.max}} 之间")),
		validation.Field(&m.ShipperArea, validation.When(m.ShipperArea.Valid, validation.Length(1, 50).Error("发件人区长度必须在 {{.min}}-{{.max}} 之间"))),
		validation.Field(&m.ShipperStreet, validation.Required.Error("发件人详细地址不能为空"), validation.Length(1, 100).Error("发件人详细地址长度必须在 {{.min}}-{{.max}} 之间")),
		validation.Field(&m.ShipperCode, validation.Required.Error("发件人邮编不能为空")),
		validation.Field(&m.ShipperEmail, validation.When(m.ShipperEmail.Valid, validation.Length(1, 100).Error("发件人邮箱长度必须在 {{.min}}-{{.max}} 之间"))),
	)
}

// OrderConsignee 收件人信息
type OrderConsignee struct {
	ConsigneeName    string      `json:"consigneeName"`             // 收件人-姓名, 长度为 1-100
	ConsigneePhone   string      `json:"consigneePhone"`            // 收件人-手机号, 长度为 10-14
	ConsigneeCountry string      `json:"consigneeCountry"`          // 收件人-国家
	ConsigneeState   string      `json:"consigneeState"`            // 收件人-州, 长度 1-35
	ConsigneeCity    string      `json:"consigneeCity"`             // 收件人-市, 长度 1-50
	ConsigneeArea    null.String `json:"consigneeArea,omitempty"`   // 收件人-区, 长度 1-50
	Address1         string      `json:"address1"`                  // 收件地址 1, 长度 1-255
	Address2         null.String `json:"address2,omitempty"`        // 收件地址 2, 长度 1-255
	Address3         null.String `json:"address3,omitempty"`        // 收件地址 3, 长度 1-255
	ConsigneeCode    string      `json:"consigneeCode"`             // 收件人-邮编, 必须为 5-6 位数字
	ConsigneeNumIn   null.String `json:"consigneeNumIn,omitempty"`  // 收件人-内门牌号, 长度 1-20
	ConsigneeNumExt  null.String `json:"consigneeNumExt,omitempty"` // 收件人-外门牌号, 长度 1-20
	Remarks          null.String `json:"remarks,omitempty"`         // 收件地址的附加信息, 长度 1-120
	ConsigneeEmail   null.String `json:"consigneeEmail,omitempty"`  // 收件人-邮箱, 长度为 1-100
}

func (m OrderConsignee) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.ConsigneeName, validation.Required.Error("收件人姓名不能为空"), validation.Length(1, 100).Error("收件人姓名长度必须在 {{.min}}-{{.max}} 之间")),
		validation.Field(&m.ConsigneePhone, validation.Required.Error("收件人手机号不能为空"), validation.Length(10, 14).Error("收件人手机号长度必须在 {{.min}}-{{.max}} 之间")),
		validation.Field(&m.ConsigneeCountry, validation.Required.Error("收件人国家不能为空")),
		validation.Field(&m.ConsigneeState, validation.Required.Error("收件人州不能为空"), validation.Length(1, 35).Error("收件人州长度必须在 {{.min}}-{{.max}} 之间")),
		validation.Field(&m.ConsigneeCity, validation.Required.Error("收件人市不能为空"), validation.Length(1, 50).Error("收件人市长度必须在 {{.min}}-{{.max}} 之间")),
		validation.Field(&m.Address1, validation.Required.Error("收件地址 1 不能为空"), validation.Length(1, 255).Error("收件地址 1 长度必须在 {{.min}}-{{.max}} 之间")),
		validation.Field(&m.ConsigneeCode, validation.Required.Error("收件人邮编不能为空"), validation.Length(5, 6).Error("收件人邮编必须为 {{.min}}-{{.max}} 位数字")),
	)
}

// OrderGoods 订单货物规格
type OrderGoods struct {
	Weight     float64     `json:"weight"`               // 包裹预报重量, 单位: kg, 值为 0.001-99.00
	Length     float64     `json:"length"`               // 包裹的长, 单位: cm, 范围 1 到 999
	Height     float64     `json:"height"`               // 包裹的高, 单位: cm, 范围 1 到 999
	Width      float64     `json:"width"`                // 包裹的宽, 单位: cm, 范围 1 到 999
	LengthUnit null.String `json:"lengthUnit,omitempty"` // 包裹长度的计量单位, 例如厘米(CM)、米(M)、英寸(INCH)。默认为 CM
	WidthUnit  null.String `json:"widthUnit,omitempty"`  // 包裹宽度的计量单位, 默认为 CM
	HeightUnit null.String `json:"heightUnit,omitempty"` // 包裹高度的计量单位, 默认为 CM
	WeightUnit null.String `json:"weightUnit,omitempty"` // 包裹预报重量的计量单位, 例如千克(KG)、磅(LB)。默认为 KG
}

func (m OrderGoods) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Weight, validation.Required.Error("包裹预报重量不能为空"), validation.Min(0.001).Error("包裹预报重量不能小于 {{.min}}"), validation.Max(99.00).Error("包裹预报重量不能大于 {{.max}}")),
		validation.Field(&m.Length, validation.Required.Error("包裹的长不能为空"), validation.Min(1.0).Error("包裹的长不能小于 {{.min}}"), validation.Max(999.0).Error("包裹的长不能大于 {{.max}}")),
		validation.Field(&m.Height, validation.Required.Error("包裹的高不能为空"), validation.Min(1.0).Error("包裹的高不能小于 {{.min}}"), validation.Max(999.0).Error("包裹的高不能大于 {{.max}}")),
		validation.Field(&m.Width, validation.Required.Error("包裹的宽不能为空"), validation.Min(1.0).Error("包裹的宽不能小于 {{.min}}"), validation.Max(999.0).Error("包裹的宽不能大于 {{.max}}")),
	)
}

// OrderItem 订单物品信息
type OrderItem struct {
	ItemNameEn string `json:"itemNameEn"` // 物品名称, 长度 1-128
	ItemNameZh string `json:"itemNameZh"` // 物品中文名称, 长度 1-60
	ItemQty    int    `json:"itemQty"`    // 物品件数, 范围 1 到 9999
	EntryPort  string `json:"entryPort"`  // 入口岸
}

func (m OrderItem) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.ItemNameEn, validation.Required.Error("物品名称不能为空"), validation.Length(1, 128).Error("物品名称长度必须在 {{.min}}-{{.max}} 之间")),
		validation.Field(&m.ItemNameZh, validation.Required.Error("物品中文名称不能为空"), validation.Length(1, 60).Error("物品中文名称长度必须在 {{.min}}-{{.max}} 之间")),
		validation.Field(&m.ItemQty, validation.Required.Error("物品件数不能为空"), validation.Min(1).Error("物品件数不能小于 {{.min}}"), validation.Max(9999).Error("物品件数不能大于 {{.max}}")),
		validation.Field(&m.EntryPort, validation.Required.Error("入口岸不能为空")),
	)
}

// OrderInsurance 订单保价
type OrderInsurance struct {
	InsuredAmount float64 `json:"insuredAmount"` // 保价金额(是否保价为是时必填), 范围 0.0001-10000
}

func (m OrderInsurance) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.InsuredAmount, validation.Required.Error("保价金额不能为空"), validation.Min(0.0001).Error("保价金额不能小于 {{.min}}"), validation.Max(10000.0).Error("保价金额不能大于 {{.max}}")),
	)
}

// CreateOrderResponseData 创建订单响应数据
type CreateOrderResponseData struct {
	COrderNo        string `json:"cOrderNo"`        // 客户单号
	VerificationPin string `json:"verificationPin"` // 签收 PIN 码
	Type            string `json:"type"`            // 操作类型
	WaybillNo       string `json:"waybillNo"`       // 运单号
}

// CancelOrderRequest 取消订单请求
type CancelOrderRequest struct {
	OrderNo string      `json:"orderNo"`           // GOFO 的运单号
	Remarks null.String `json:"remarks,omitempty"` // 取消备注, 长度 1-100
}

func (m CancelOrderRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.OrderNo, validation.Required.Error("运单号不能为空")),
		validation.Field(&m.Remarks, validation.When(m.Remarks.Valid, validation.Length(1, 100).Error("取消备注长度必须在 {{.min}}-{{.max}} 之间"))),
	)
}

// TrackEvent 轨迹事件
type TrackEvent struct {
	PubEsContext    string `json:"pubEsContext"`    // 轨迹描述
	OperationMove   string `json:"operationMove"`   // 轨迹编码
	OrderNo         string `json:"orderNo"`         // 运单单号
	ThirdWaybillNo  string `json:"thirdWaybillNo"`  // 客户单号
	Operator        string `json:"operator"`        // 操作人姓名
	OperationTime   string `json:"operationTime"`   // 操作时间
	GroupTimeZone   string `json:"groupTimeZone"`   // 时区
	Pin             string `json:"pin"`             // 是否通过 pin 签收
	EnContext       string `json:"enContext"`       // 轨迹英文描述
	SignerType      string `json:"signerType"`      // 签收人类型
	Location        string `json:"location"`        // 轨迹发生地点
	DeptId          int    `json:"dept_id"`         //
	Signer          string `json:"signer"`          // 签收人
	ErrorCode       int    `json:"errorCode"`       // 异常类型
	ProcessCity     string `json:"processCity"`     // 轨迹发生城市
	ProcessProvince string `json:"processProvince"` // 轨迹发生州
	ProcessPostCode string `json:"processPostCode"` // 轨迹发生邮编
}
