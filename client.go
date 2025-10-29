package gofo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-resty/resty/v2"
	"github.com/hiscaler/gofo-go/config"
	"github.com/hiscaler/gofo-go/entity"
)

const (
	Version   = "0.0.1"
	userAgent = "GOFO API Client-Golang/" + Version + " (https://github.com/hiscaler/gofo-go)"
)

const (
	ProdBaseUrl = "https://dmsapi.gofoexpress.com"
	TestBaseUrl = "https://uat-dbu-api.eminxing.com"
)

type Client struct {
	config      *config.Config // 配置
	httpClient  *resty.Client  // Resty Client
	accessToken string         // AccessToken
	Services    services       // API Services
}

func NewClient(ctx context.Context, cfg config.Config) *Client {
	l := createLogger()
	gofoClient := &Client{
		config: &cfg,
	}
	baseUrl := ProdBaseUrl
	if cfg.Env != entity.Prod {
		baseUrl = TestBaseUrl
	}
	httpClient := resty.New().
		SetDebug(cfg.Debug).
		SetBaseURL(baseUrl).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
			"User-Agent":   userAgent,
		})
	httpClient.SetTimeout(time.Duration(cfg.Timeout)*time.Second).
		SetBasicAuth(cfg.Account, cfg.Password).
		SetRetryCount(2).
		SetRetryWaitTime(2 * time.Second).
		SetRetryMaxWaitTime(5 * time.Second)
	gofoClient.httpClient = httpClient
	xService := service{
		config:     &cfg,
		logger:     l.l,
		httpClient: gofoClient.httpClient,
	}
	gofoClient.Services = services{
		Order: (orderService)(xService),
	}
	return gofoClient
}

type NormalResponse struct {
	Code           int    `json:"code"`
	Message        string `json:"msg"`
	EnglishMessage string `json:"msgEn"`
	Data           any    `json:"data"`
}

// errorWrap 错误包装
func errorWrap(code int, message string) error {
	if code == 200 {
		return nil
	}

	switch code {
	case 301:
		message = "参数异常"
	case 305:
		message = "数据不存在"
	case 401:
		message = "未认证（签名错误）"
	case 404:
		message = "接口不存在"
	case 500:
		if message == "" {
			message = "操作失败"
		}
	default:
		message = strings.TrimSpace(message)
		if message == "" {
			message = "Unknown error"
		}
	}
	return fmt.Errorf("%d: %s", code, message)
}

func invalidInput(e error) error {
	var errs validation.Errors
	if !errors.As(e, &errs) {
		return e
	}

	if len(errs) == 0 {
		return nil
	}

	fields := make([]string, 0)
	messages := make([]string, 0)
	for field := range errs {
		fields = append(fields, field)
	}
	sort.Strings(fields)

	for _, field := range fields {
		e1 := errs[field]
		if e1 == nil {
			continue
		}

		var errObj validation.ErrorObject
		if errors.As(e1, &errObj) {
			e1 = errObj
		} else {
			var errs1 validation.Errors
			if errors.As(e1, &errs1) {
				e1 = invalidInput(errs1)
				if e1 == nil {
					continue
				}
			}
		}

		messages = append(messages, e1.Error())
	}
	return errors.New(strings.Join(messages, "; "))
}

func recheckError(resp *resty.Response, e error) error {
	if e != nil {
		if errors.Is(e, http.ErrHandlerTimeout) {
			return errorWrap(http.StatusRequestTimeout, e.Error())
		}
		return e
	}

	if resp.IsError() {
		return fmt.Errorf("%v", resp.Error())
	} else {
		var normalResponse NormalResponse
		err := json.Unmarshal(resp.Body(), &normalResponse)
		if err != nil {
			return err
		}
		return errorWrap(normalResponse.Code, normalResponse.Message)
	}
}
