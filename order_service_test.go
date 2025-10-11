package gofo

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hiscaler/gofo-go/config"
	"github.com/hiscaler/gofo-go/entity"
	"gopkg.in/guregu/null.v4"
)

func TestOrderService_Create(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/open-api/v2/order/create" {
			t.Errorf("Expected to request '/open-api/v2/order/create', got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		response := entity.CreateOrderResponse{
			Code: 200,
			Msg:  "操作成功",
			Data: entity.CreateOrderResponseData{
				COrderNo:  "TEST_ORDER_001",
				WaybillNo: "WAYBILL_001",
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient(context.Background(), config.Config{Env: entity.Test})
	client.httpClient.SetBaseURL(server.URL)

	req := entity.CreateOrderRequest{
		COrderNo:      null.StringFrom("TEST_ORDER_001"),
		DeclaredValue: 12,
		OrderShipper: entity.OrderShipper{
			ShipperName:    "test",
			ShipperPhone:   "13000000000",
			ShipperCountry: "CN",
			ShipperState:   "Guangdong",
			ShipperCity:    "Shenzhen",
			ShipperStreet:  "test street",
			ShipperCode:    "518000",
		},
		OrderConsignee: entity.OrderConsignee{
			ConsigneeName:    "test",
			ConsigneePhone:   "13000000000",
			ConsigneeCountry: "US",
			ConsigneeState:   "California",
			ConsigneeCity:    "Los Angeles",
			Address1:         "test address",
			ConsigneeCode:    "90001",
		},
		OrderGoods: entity.OrderGoods{
			Weight: 1,
			Length: 1,
			Height: 1,
			Width:  1,
		},
		OrderItemList: []entity.OrderItem{
			{
				ItemNameEn: "test",
				ItemNameZh: "测试",
				ItemQty:    1,
				EntryPort:  "LAX",
			},
		},
	}
	_, err := client.Services.Order.Create(context.Background(), req)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestOrderService_Cancel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/open-api/v2/order/cancel" {
			t.Errorf("Expected to request '/open-api/v2/order/cancel', got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		response := entity.CancelOrderResponse{
			Code: 200,
			Msg:  "操作成功",
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient(context.Background(), config.Config{Env: entity.Test})
	client.httpClient.SetBaseURL(server.URL)

	req := entity.CancelOrderRequest{
		OrderNo: "WAYBILL_001",
	}
	_, err := client.Services.Order.Cancel(context.Background(), req)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestOrderService_GetOrderLabel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/open-api/v2/order/getOrderLabelUrlV2" {
			t.Errorf("Expected to request '/open-api/v2/order/getOrderLabelUrlV2', got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		response := entity.GetOrderLabelResponse{
			Code: 200,
			Data: entity.GetOrderLabelResponseData{
				Base64Code: "JVBERi0xLjUNCiXi48/TDQo...",
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient(context.Background(), config.Config{Env: entity.Test})
	client.httpClient.SetBaseURL(server.URL)

	_, err := client.Services.Order.GetOrderLabel(context.Background(), "WAYBILL_001")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestOrderService_Track(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/open-api/v2/order/track/WAYBILL_001" {
			t.Errorf("Expected to request '/open-api/v2/order/track/WAYBILL_001', got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		response := entity.TrackOrderResponse{
			Code: 200,
			Data: []entity.TrackEvent{},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient(context.Background(), config.Config{Env: entity.Test})
	client.httpClient.SetBaseURL(server.URL)

	_, err := client.Services.Order.Track(context.Background(), "WAYBILL_001")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
