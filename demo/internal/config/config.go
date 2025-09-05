package config

import "github.com/zeromicro/go-zero/rest"

// Config 是服务配置根结构
// - RestConf 由 go-zero 提供，包含 HTTP 端口、日志、超时等
// - Database 为自定义数据库配置（驱动、DSN、日志开关）
type Config struct {
	rest.RestConf
	Database DatabaseConfig
}

// DatabaseConfig 描述数据库连接所需参数
// - Driver: 驱动标识（mysql/postgres/sqlite）
// - DSN: 数据源字符串，形如 user:pass@tcp(host:port)/db?params
// - LogMode: 是否开启 ORM SQL 日志（开发环境可开）
type DatabaseConfig struct {
	DSN     string `json:"dsn" yaml:"dsn"`
	Driver  string `json:"driver" yaml:"driver"`
	LogMode bool   `json:"logMode" yaml:"logMode"`
}
