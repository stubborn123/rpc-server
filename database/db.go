package database

import (
	"context"
	"fmt"
	"rpc-server/config"

	"gorm.io/gorm"
)

type DB struct {
	//gorm是Go语言流行ORM（对象关系映射）库，用于简化数据库操作（*gorm.DB是该库核心类型，数据库链接或会话对象，实际上Java的mybatis也是ORM）
	gorm *gorm.DB
	//Exchange的实例
	Exchange             ExchangeDB
	SupportToken         SupportTokenDB
	SupportTokenExchange SupportTokenExchangeDB
}

func NewDB(ctx context.Context, dbConfig config.DBConfig) (*DB, error) {
	dsn := fmt.Sprintf("host =%s dbName=%s sslMode=disable", dbConfig.Host, dbConfig.Name)

	if dbConfig.Port != 0 {
		dsn += fmt.Sprintf("port=%d", dbConfig.Port)
	}

	if dbConfig.User != "" {
		dsn += fmt.Sprintf("port=%s", dbConfig.User)
	}

	if dbConfig.Password != "" {
		dsn += fmt.Sprintf("port=%s", dbConfig.Password)
	}

	gormConfig := gorm.Config{
		SkipDefaultTransaction: true,
		CreateBatchSize:        3_000,
	}

	retryStrategy := &retry.ExponentiaSta

}
