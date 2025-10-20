package entity

// OrderCreateResponse 订单
type OrderCreateResponse struct {
	FourSegmentCode string `json:"fourSegmentCode"`
	COrderNo        string `json:"cOrderNo"`        // 客户单号
	VerificationPin string `json:"verificationPin"` // 签收 PIN 码
	Type            string `json:"type"`            // 操作类型
	WaybillNo       string `json:"waybillNo"`       // 运单号
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
