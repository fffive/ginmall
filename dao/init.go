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
	// 日志打印
	var ormlogger logger.Interface
	if gin.Mode() == "debug" {
		ormlogger = logger.Default.LogMode(logger.Info)
	} else {
		ormlogger = logger.Default
	}

	// 打开Database 配置 https://gorm.io/zh_CN/docs/connecting_to_the_database.html
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connRead,
		DefaultStringSize:         256,  //string 类型默认长度
		DisableDatetimePrecision:  true, // 禁止使用datetime精度，数据库5.6 之前不知道次
		DontSupportRenameIndex:    true, //重命名索引时采取删除并创建mysql 5.7 之前不支持
		DontSupportRenameColumn:   true, //用change重命名列 mysql 8.0 之前不支持
		SkipInitializeWithVersion: true, // 根据版本自动配置
	}), &gorm.Config{
		Logger: ormlogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return
	}
	// 创建sqlDb
	sqlDb, _ := db.DB()
	// 设置最大连接池 最大连接数  设置连接时间最大值
	sqlDb.SetConnMaxLifetime(time.Second * 20)
	sqlDb.SetMaxIdleConns(20)
	sqlDb.SetMaxOpenConns(200)
	_db = db

	// 主从配置  https://gorm.io/zh_CN/docs/dbresolver.html
	_ = _db.Use(dbresolver.Register(dbresolver.Config{
		Sources:           []gorm.Dialector{mysql.Open(connWrite)},
		Replicas:          []gorm.Dialector{mysql.Open(connWrite), mysql.Open(connRead)},
		Policy:            dbresolver.RandomPolicy{}, // 负载均衡
		TraceResolverMode: true,
	}))

	// 迁移 
	Migration();
}

// 确保上下文连接
func NewDBclient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
