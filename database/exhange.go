package database

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// 对应的json就是指定结构体字段与JSON数据的映射关系，Go也是依赖依赖 反射（reflection） 机制
type Exchange struct {
	// uuid.UUID类型（通用唯一识别码）反引号包围的结构体标签（标签作用主要是数据库映射，序列化控制，验证规则，配置设定将配置值绑定到结构体字段）
	GUID uuid.UUID `gorm:"primaryKey" json:"guid"`
	// 这里gorm对应的表的映射列，json代表影射的是name的键名
	Name      string `gorm:"name" json:"name"`
	Config    string `gorm:"config" json:"config"`
	Timestamp uint64 `json:"timestamp"`
}

func (Exchange) TableName() string {
	return "exchange"
}

// 定义一个ExchangeDB接口，定义了2个接口方法
type ExchangeDB interface {
	//查询
	ExchangeView
	//批量新增
	StoreExchanges([]Exchange) error
}

// 查询接口（接口套接口-接口嵌入 组合的形式，对比Java的接口继承接口）
type ExchangeView interface {
	//在接口里面定义，对比用独立方法，方便后面切换ORM
	QueryExchangeGuid(string) (*Exchange, error)
}

type exchangeDB struct {
	gorm *gorm.DB
}

func NewExchangeDB(db *gorm.DB) ExchangeDB {
	return &exchangeDB{gorm: db}
}

func (exd *exchangeDB) QueryExchangeGuid(guid string) (*Exchange, error) {
	var exchangeItem Exchange

	//gorm操作：从数据库从exchange表中查询指定GUID的交易所配置项
	//Take(&exchangeItem),Take获取单条记录，把数据结果传给&exchangeItem这个指针
	result := exd.gorm.Table("exchange").Where("guid=?", guid).Take(&exchangeItem)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, result.Error
		}
	}

	return &exchangeItem, nil
}

// 批量插入操作
func (exd *exchangeDB) StoreExchanges(exchangeList []Exchange) error {
	result := exd.gorm.Table("exchange").CreateInBatches(&exchangeList, len(exchangeList))
	return result.Error
}
