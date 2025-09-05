package model

import (
	"database/sql"
	"time"
)

// Users 对应于 MySQL 表 `users`
// - Go 中 struct 字段通过 gorm 标签与表字段映射
// - ID 使用无符号 64 位整数，对应 bigint unsigned
// - Username/Password/Email 分别映射到相应列；Email 使用 sql.NullString 表示可空
// - CreatedAt/UpdatedAt 使用 time.Time，对应 timestamp 列
// - TableName() 返回固定表名，避免 GORM 的默认复数/小写推断差异
// 参考 DDL:
// CREATE TABLE `users` (
//
//	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
//	`username` varchar(50) NOT NULL UNIQUE,
//	`password` varchar(255) NOT NULL,
//	`email` varchar(100) DEFAULT NULL UNIQUE,
//	`created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
//	`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
//	PRIMARY KEY (`id`)
//
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
type Users struct {
	ID           uint64         `gorm:"primaryKey;autoIncrement;column:id"`
	Username     string         `gorm:"size:50;not null;uniqueIndex;column:username"`
	PasswordHash string         `gorm:"size:255;not null;column:password"`
	Email        sql.NullString `gorm:"size:100;uniqueIndex;column:email"`
	CreatedAt    time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt    time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`
}

// TableName 指定模型绑定的表名
func (Users) TableName() string { return "users" }
