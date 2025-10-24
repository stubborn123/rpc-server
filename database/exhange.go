package database

import "github.com/google/uuid"

type Exchange struct {
	// uuid.UUID类型（通用唯一识别码）反引号包围的结构体标签（标签作用主要是数据库映射，序列化控制，验证规则，配置设定将配置值绑定到结构体字段）
	GUID uuid.UUID `gorm:"primaryKey" json:"guid"`
	// 这里gorm对应的表的映射列，json代表影射的是name的键名
	Name      string `gorm:"name" json:"name"`
	Config    string `gorm:"config" json:"config"`
	Timestamp uint64 `json:"timestamp"`
}
