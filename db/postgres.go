package db

import (
	"context"
	"errors"
	"fmt"
	gormlog "log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"tiktok/pkg/config"
	"tiktok/pkg/log"
)

var (
	tx *gorm.DB

	schemas = []interface{}{
		&User{},
		// TODO:
	}
)

func init() {
	var err error
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.C.Postgresql.Host,
		config.C.Postgresql.User,
		config.C.Postgresql.Password,
		config.C.Postgresql.Dbname,
		config.C.Postgresql.Port,
	)
	logLevel := logger.Warn
	if config.C.SQLDebug {
		logLevel = logger.Info
	}
	newLogger := logger.New(
		gormlog.New(os.Stdout, "\r\n", gormlog.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logLevel,    // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)

	tx, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Logger.Error("postgres connection failed: ", err)
		return
	}
	err = tx.AutoMigrate(schemas...)
	if err != nil {
		log.Logger.Error("postgres migrate failed: ", err)
		return
	}

	log.Logger.Info("PostgreSQL server connected :D")
}

// --- Transaction ---

func beginTransactionWithCtx(ctx context.Context) *gorm.DB {
	tx := tx.Begin()
	if err := tx.Error; err != nil {
		log.Logger.Error("begin transaction error: ", err)
	}
	return tx.WithContext(ctx)
}

func (m *Context) Close() {
	// if !m.abort {
	// 	m.tx.Commit()
	// }
	if r := recover(); r != nil {
		m.Transaction.Rollback()
	}
	m.cancel()
}

func (m *Context) Abort() {
	m.Transaction.Rollback()
	m.cancel()
	// m.abort = true
}

func (m *Context) Commit() error {
	m.commit = true
	return m.Transaction.Commit().Error
}

func (m *Context) IsExisted(query interface{}) bool {
	return !errors.Is(m.Transaction.Where(query).Take(query).Error, gorm.ErrRecordNotFound) 
}
