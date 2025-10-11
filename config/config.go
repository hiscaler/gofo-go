package config

type Config struct {
	Debug    bool   `json:"debug"`    // 是否启用调试模式
	Env      string `json:"env"`      // 环境
	Timeout  int    `json:"timeout"`  // HTTP 超时设定（单位：秒）
	Account  string `json:"account"`  // 用户账号
	Password string `json:"password"` // 用户密码
}
