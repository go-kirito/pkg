/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2021/3/18 14:03
 */
package zdb

type Options struct {
	Dns          string
	MaxIdleConns int
	MaxOpenConns int
	Mode         string
}

func newOptions() *Options {
	return &Options{}
}
