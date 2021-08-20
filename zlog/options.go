/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/9/27 17:56
 */
package zlog

import "github.com/go-kirito/pkg/zconfig"

type options struct {
	Driver     string
	Format     string
	Output     string
	Level      string
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

type OptionFunc func(*options)

func NewOptionsWithConfig() options {

	opt := options{
		Driver:     "zap",
		Output:     "console",
		Level:      "debug",
		Format:     "text",
		Filename:   "./access.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}

	if err := zconfig.UnmarshalKey("log", &opt); err != nil {
		return opt
	}

	return opt

}

func Driver(driver string) OptionFunc {
	return func(o *options) {
		o.Driver = driver
	}
}

func Format(format string) OptionFunc {
	return func(o *options) {
		o.Format = format
	}
}

func Output(output string) OptionFunc {
	return func(o *options) {
		o.Output = output
	}
}

func Level(level string) OptionFunc {
	return func(o *options) {
		o.Level = level
	}
}

func Filename(filename string) OptionFunc {
	return func(o *options) {
		o.Filename = filename
	}
}

func MaxSize(maxSize int) OptionFunc {
	return func(o *options) {
		o.MaxSize = maxSize
	}
}

func MaxBackups(maxBackups int) OptionFunc {
	return func(o *options) {
		o.MaxBackups = maxBackups
	}
}

func MaxAge(maxAge int) OptionFunc {
	return func(o *options) {
		o.MaxAge = maxAge
	}
}

func Compress(compress bool) OptionFunc {
	return func(o *options) {
		o.Compress = compress
	}
}
