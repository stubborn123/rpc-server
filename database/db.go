package database

import (
	"gorm.io/gorm"
)

type DB struct {
	//gorm是Go语言流行ORM（对象关系映射）库，用于简化数据库操作（*gorm.DB是该库核心类型，数据库链接或会话对象，实际上Java的mybatis也是ORM）
	gorm         *gorm.DB
	Exchange     ExchangeDB
	SupportToken SupportTokenDB
}
