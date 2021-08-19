/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/9/27 17:58
 */
package zap

type options struct {
	format     string
	output     string
	level      string
	filename   string
	maxSize    int
	maxBackups int
	maxAge     int
	compress   bool
}

type OptionFunc func(*options)

func Format(format string) OptionFunc {
	return func(o *options) {
		o.format = format
	}
}

func Output(output string) OptionFunc {
	return func(o *options) {
		o.output = output
	}
}

func Level(level string) OptionFunc {
	return func(o *options) {
		o.level = level
	}
}

func Filename(filename string) OptionFunc {
	return func(o *options) {
		o.filename = filename
	}
}

func MaxSize(maxSize int) OptionFunc {
	return func(o *options) {
		o.maxSize = maxSize
	}
}

func MaxBackups(maxBackups int) OptionFunc {
	return func(o *options) {
		o.maxBackups = maxBackups
	}
}

func MaxAge(maxAge int) OptionFunc {
	return func(o *options) {
		o.maxAge = maxAge
	}
}

func Compress(compress bool) OptionFunc {
	return func(o *options) {
		o.compress = compress
	}
}
