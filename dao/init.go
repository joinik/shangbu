package dao

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

var _db *gorm.DB

func Database(connRead, connWrite string) {
	var ormLogger logger.Interface

	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connRead, // DSN data source name
		DefaultStringSize:         256,      // string 类型字段的默认长度
		DisableDatetimePrecision:  true,     // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,     // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,     // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,    // 根据版本自动配置
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()
	// 设置连接池 空闲
	sqlDB.SetMaxIdleConns(20)
	// 打开
	sqlDB.SetMaxOpenConns(100)
	// 超时
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	_db = db

	// 读写 分离设置
	_db.Use(dbresolver.
		Register(dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(connRead)},                         //写操作
			Replicas: []gorm.Dialector{mysql.Open(connWrite), mysql.Open(connWrite)}, // 读操作
			Policy:   dbresolver.RandomPolicy{},                                      // sources/replicas 负载均很策略
		}))

	Migration()

}

// 创建会话
func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}

type NewDBDao struct {
	*gorm.DB
}

func NewDBDaoByDB(db *gorm.DB) *NewDBDao {
	return &NewDBDao{db}
}
