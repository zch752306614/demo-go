package svc

import (
	"demo/internal/config"
	"demo/internal/model"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ServiceContext 相当于 Java 项目中的 "应用上下文/容器"
// - 保存全局可复用的依赖（如配置与数据库连接）
// - 由 main 在启动时构建，并注入到各处（handler/logic）
type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

// NewServiceContext 根据配置初始化依赖
// - 初始化 GORM（MySQL 驱动）
// - 执行 AutoMigrate（开发环境友好；生产建议使用迁移脚本）
func NewServiceContext(c config.Config) *ServiceContext {
	db, err := openGorm(c)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	if err := db.AutoMigrate(&model.Users{}); err != nil {
		log.Fatalf("auto migrate failed: %v", err)
	}
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}

// openGorm 按配置创建 *gorm.DB
func openGorm(c config.Config) (*gorm.DB, error) {
	cfg := &gorm.Config{}
	if c.Database.LogMode {
		cfg.Logger = logger.Default.LogMode(logger.Info)
	}
	switch c.Database.Driver {
	case "mysql":
		return gorm.Open(mysql.Open(c.Database.DSN), cfg)
	default:
		return nil, fmt.Errorf("unsupported driver: %s", c.Database.Driver)
	}
}
