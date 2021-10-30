/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2021/3/11 17:14
 */
package zdb

import (
	"context"
	"database/sql"
	"sync"

	"github.com/go-kirito/pkg/zlog"

	"gorm.io/gorm"
)

var dm *dbManager

var tm *txManager

var mu sync.RWMutex

type funcType func(ctx context.Context) (interface{}, error)

//初始化数据库
func InitMySQL() {
	//初始化数据库管理
	dm = newDBManager()
	//初始化事务管理
	tm = newTxManager()
}

func NewOrm(ctx context.Context, dbName ...string) (db *DB) {

	name := "default"

	if len(dbName) > 0 {
		name = dbName[0]
	}

	tc, ok := ctx.(*TxContext)

	if !ok {
		db = dm.getDB(name)

		if db == nil {
			zlog.Panicf("数据库 `%s` 连接不存在", name)
		}
		return
	}

	txId := tc.getTxId()

	//TODO 并发问题
	mu.Lock()
	defer mu.Unlock()
	db = tm.getDB(name, txId)

	if db != nil {
		db.orm = db.orm.Session(&gorm.Session{NewDB: true})
		db.orm.Error = nil
		return db
	}

	//有事务id，但是没有绑定的db，说明还未开启事务
	db = dm.getDB(name)
	if db == nil {
		zlog.Panicf("数据库 `%s` 连接不存在", name)
		return
	}

	opts := tc.getSQLOpts()

	if err := db.Begin(opts...).Error(); err != nil {
		zlog.Panic("开启事务失败:" + err.Error())
	}

	tm.addDB(name, txId, db)

	return
}

func Transaction(ctx context.Context, fn funcType, opts ...*sql.TxOptions) (resp interface{}, err error) {
	panicked := true

	//处理嵌套事务的问题
	tc, ok := ctx.(*TxContext)
	if !ok {
		//生成事务的context
		tc = tm.BeginTx(ctx, opts...)
	}

	defer func() {
		if panicked || err != nil {
			//回滚当前context下的事务
			tm.Rollback(tc)
		}
	}()

	resp, err = fn(tc)

	if err == nil {
		//提交当前context的事务
		err = tm.Commit(tc)
	}

	panicked = false

	return
}
