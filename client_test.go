package gofo

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/hiscaler/gofo-go/config"
)

var client *Client
var ctx context.Context

func TestMain(m *testing.M) {
	b, err := os.ReadFile("./config/config.json")
	if err != nil {
		panic(fmt.Sprintf("Read config error: %s", err.Error()))
	}
	var cfg config.Config
	err = json.Unmarshal(b, &cfg)
	if err != nil {
		panic(fmt.Sprintf("Parse config file error: %s", err.Error()))
	}

	client = NewClient(ctx, cfg)
	ctx = context.Background()
	m.Run()
}
