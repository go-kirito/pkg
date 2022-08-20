/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2021/3/11 15:25
 */
package zdb

import (
	"database/sql"

	"gorm.io/gorm"
)

type DB struct {
	name string
	orm  *gorm.DB
}

func newDB(name string, orm *gorm.DB) *DB {
	return &DB{
		name: name,
		orm:  orm,
	}
}

func (db *DB) AutoMigrate(value ...interface{}) error {
	return db.orm.AutoMigrate(value...)
}

func (db *DB) Model(value interface{}) *DB {
	db.orm = db.orm.Model(value)
	return db
}

func (db *DB) Table(name string, args ...interface{}) *DB {
	db.orm = db.orm.Table(name, args...)
	return db
}

func (db *DB) Select(query interface{}, args ...interface{}) *DB {
	db.orm = db.orm.Select(query, args...)
	return db
}

func (db *DB) Distinct(args ...interface{}) *DB {
	db.orm = db.orm.Distinct(args...)
	return db
}

func (db *DB) Omit(columns ...string) *DB {
	db.orm = db.orm.Omit(columns...)
	return db
}

func (db *DB) Where(query interface{}, args ...interface{}) *DB {
	db.orm = db.orm.Where(query, args...)
	return db
}

func (db *DB) Not(query interface{}, args ...interface{}) *DB {
	db.orm = db.orm.Not(query, args...)
	return db
}

func (db *DB) Or(query interface{}, args ...interface{}) *DB {
	db.orm = db.orm.Or(query, args...)
	return db
}

func (db *DB) Joins(query string, args ...interface{}) *DB {
	db.orm = db.orm.Joins(query, args...)
	return db
}

func (db *DB) GroupBy(name string) *DB {
	db.orm = db.orm.Group(name)
	return db
}

func (db *DB) Having(query interface{}, args ...interface{}) *DB {
	db.orm = db.orm.Having(query, args...)
	return db
}

func (db *DB) OrderBy(value interface{}) *DB {
	db.orm = db.orm.Order(value)
	return db
}

func (db *DB) Limit(limit int) *DB {
	db.orm = db.orm.Limit(limit)
	return db
}

func (db *DB) Offset(offset int) *DB {
	db.orm = db.orm.Offset(offset)
	return db
}

func (db *DB) Scopes(funcs ...func(*DB) *DB) *DB {
	for _, f := range funcs {
		db = f(db)
	}
	return db
}

func (db *DB) Raw(sql string, values ...interface{}) *DB {
	db.orm = db.orm.Raw(sql, values...)
	return db
}

func (db *DB) Create(value interface{}) (tx *DB) {
	db.orm = db.orm.Create(value)
	return db
}

func (db *DB) Save(value interface{}) (tx *DB) {
	db.orm = db.orm.Save(value)
	return db
}

func (db *DB) First(dest interface{}, conds ...interface{}) *DB {
	db.orm = db.orm.First(dest, conds...)
	return db
}

func (db *DB) Take(dest interface{}, conds ...interface{}) *DB {
	db.orm = db.orm.Take(dest, conds...)
	return db
}

func (db *DB) Last(dest interface{}, conds ...interface{}) *DB {
	db.orm = db.orm.Last(dest, conds...)
	return db
}

func (db *DB) Find(dest interface{}, conds ...interface{}) *DB {
	db.orm = db.orm.Find(dest, conds...)
	return db
}

func (db *DB) FirstOrCreate(dest interface{}, conds ...interface{}) *DB {
	db.orm = db.orm.FirstOrCreate(dest, conds...)
	return db
}

func (db *DB) Update(column string, value interface{}) *DB {
	db.orm = db.orm.Update(column, value)
	return db
}

func (db *DB) Updates(values interface{}) *DB {
	db.orm = db.orm.Updates(values)
	return db
}

func (db *DB) UpdateColumn(column string, value interface{}) *DB {
	db.orm = db.orm.UpdateColumn(column, value)
	return db
}

func (db *DB) UpdateColumns(values interface{}) *DB {
	db.orm = db.orm.UpdateColumns(values)
	return db
}

func (db *DB) Delete(value interface{}, conds ...interface{}) *DB {
	db.orm = db.orm.Delete(value, conds...)
	return db
}

func (db *DB) Count(count *int64) *DB {
	db.orm = db.orm.Count(count)
	return db
}

func (db *DB) Pluck(column string, dest interface{}) *DB {
	db.orm = db.orm.Pluck(column, dest)
	return db
}

func (db *DB) Exec(sql string, values ...interface{}) *DB {
	db.orm = db.orm.Exec(sql, values...)
	return db
}

func (db *DB) Begin(opts ...*sql.TxOptions) *DB {
	db.orm = db.orm.Begin(opts...)
	return db
}

func (db *DB) Commit() *DB {
	db.orm = db.orm.Commit()
	return db
}

func (db *DB) Rollback() *DB {
	db.orm = db.orm.Rollback()
	return db
}

func (db *DB) Error() error {
	if db.orm.Error == gorm.ErrRecordNotFound {
		return ErrRecordNotFound
	}
	return db.orm.Error
}

func (db *DB) RowsAffected() int64 {
	return db.orm.RowsAffected
}
