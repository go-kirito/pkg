/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2021/3/11 15:26
 */
package zdb

import (
	"context"
	"database/sql"
	"sync"

	"github.com/go-kirito/pkg/util/id"
)

type txManager struct {
	m map[int64]map[string]*DB
	sync.Mutex
}

func newTxManager() *txManager {
	return &txManager{
		m: make(map[int64]map[string]*DB),
	}
}

func (tm *txManager) getDB(name string, txId int64) *DB {
	tm.Lock()
	defer tm.Unlock()
	if dbs, ok := tm.m[txId]; ok {
		if db, ok := dbs[name]; ok {
			return db
		}
	}
	return nil
}

func (tm *txManager) addDB(name string, txId int64, db *DB) {
	tm.Lock()
	defer tm.Unlock()
	dbs, ok := tm.m[txId]
	if !ok {
		dbs = make(map[string]*DB)
	}
	dbs[name] = db
	tm.m[txId] = dbs
	return
}

func (tm *txManager) BeginTx(ctx context.Context, opts ...*sql.TxOptions) *TxContext {
	txId := id.New()
	return NewTxContext(ctx, txId, opts...)
}

func (tm *txManager) Commit(tc *TxContext) error {
	tm.Lock()
	defer func() {
		delete(tm.m, tc.getTxId())
		tm.Unlock()
	}()

	if dbs, ok := tm.m[tc.getTxId()]; ok {
		for _, db := range dbs {
			err := db.Commit().Error()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (tm *txManager) Rollback(tc *TxContext) error {
	tm.Lock()
	defer func() {
		delete(tm.m, tc.getTxId())
		tm.Unlock()
	}()

	if dbs, ok := tm.m[tc.getTxId()]; ok {
		for _, db := range dbs {
			err := db.Rollback().Error()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
