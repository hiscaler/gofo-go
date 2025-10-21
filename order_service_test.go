package gofo

import (
	"fmt"
	"testing"

	"gopkg.in/guregu/null.v4"
)

func TestOrderService_Create(t *testing.T) {
	req := CreateOrderRequest{
		COrderNo:      null.StringFrom("TEST_ORDER_001"),
		DeclaredValue: 12,
		OrderShipper: OrderShipper{
			ShipperName:    "test",
			ShipperPhone:   "13000000000",
			ShipperCountry: "CN",
			ShipperState:   "Guangdong",
			ShipperCity:    "Shenzhen",
			ShipperStreet:  "test street",
			ShipperCode:    "90058",
		},
		ProductCode: null.StringFrom("GOFO Parcel Pickup"),
		OrderConsignee: OrderConsignee{
			ConsigneeName:    "test",
			ConsigneePhone:   "13000000000",
			ConsigneeCountry: "US",
			ConsigneeState:   "California",
			ConsigneeCity:    "Los Angeles",
			Address1:         "test address",
			ConsigneeCode:    "90001",
		},
		OrderGoods: OrderGoods{
			Weight: 1,
			Length: 1,
			Height: 1,
			Width:  1,
		},
		OrderItemList: []OrderItem{
			{
				ItemNameEn: "test",
				ItemNameZh: "测试",
				ItemQty:    1,
			},
		},
	}
	resp, err := client.Services.Order.Create(ctx, req)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	fmt.Println(resp)
}

func TestOrderService_Cancel(t *testing.T) {
	req := CancelOrderRequest{
		OrderNo: "GFUS01014625997824",
	}
	_, err := client.Services.Order.Cancel(ctx, req)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestOrderService_ShippingLabel(t *testing.T) {
	_, err := client.Services.Order.ShippingLabel(ctx, "GFUS01014625997824")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestOrderService_Track(t *testing.T) {
	_, err := client.Services.Order.Track(ctx, "GFUS01014625997824")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
