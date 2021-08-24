/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2021/3/11 15:26
 */
package zdb

import (
	"github.com/go-kirito/pkg/zlog"

	"github.com/go-kirito/pkg/zconfig"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type dbManager struct {
	opts map[string]*Options
	dbs  map[string]*gorm.DB
}

func newDBManager() *dbManager {
	dm := &dbManager{
		opts: make(map[string]*Options),
		dbs:  make(map[string]*gorm.DB),
	}

	//读取配置文件
	if err := zconfig.UnmarshalKey("database", &dm.opts); err != nil {
		zlog.Panic("读取数据库配置文件失败，请确认是否有`database`的配置")
	}

	//db初始化连接
	for name, opt := range dm.opts {
		engine, err := gorm.Open(mysql.Open(opt.Dns), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})

		if err != nil {
			zlog.Panic("初始化数据库失败:", err)
		}

		zlog.Infof("db:%s mode:%s 连接成功", name, opt.Mode)

		if opt.Mode == "debug" {
			engine = engine.Debug()
		}

		db, err := engine.DB()
		if err != nil {
			zlog.Panic("数据库错误:" + err.Error())
		}

		db.SetMaxOpenConns(opt.MaxOpenConns)
		db.SetMaxIdleConns(opt.MaxIdleConns)

		engine.Debug()

		dm.dbs[name] = engine
	}

	return dm
}

func (dm *dbManager) getDB(name string) *DB {
	if engine, ok := dm.dbs[name]; ok {
		return newDB(name, engine)
	}

	return nil
}
