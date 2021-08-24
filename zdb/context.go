/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2021/4/2 13:47
 */
package zdb

import (
	"context"
	"database/sql"
)

type TxContext struct {
	context.Context
	txId    int64
	sqlOpts []*sql.TxOptions
}

func NewTxContext(ctx context.Context, txId int64, opts ...*sql.TxOptions) *TxContext {
	return &TxContext{
		Context: ctx,
		txId:    txId,
		sqlOpts: opts,
	}
}

func (tc *TxContext) getTxId() int64 {
	return tc.txId
}

func (tc *TxContext) getContext() context.Context {
	return tc.Context
}

func (tc *TxContext) getSQLOpts() []*sql.TxOptions {
	return tc.sqlOpts
}
