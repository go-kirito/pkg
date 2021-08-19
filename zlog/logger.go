/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2021/3/18 14:33
 */
package zlog

import "github.com/go-kirito/pkg/zlog/zap"

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	SetLevel(level string)

	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Panic(args ...interface{})
	Fatal(args ...interface{})
}

var log Logger

func init() {
	log = NewLogger()
}

func NewLogger(optsFunc ...OptionFunc) Logger {
	opts := newOptions(optsFunc...)
	return zap.NewLogger(
		zap.Format(opts.format),
		zap.Level(opts.level),
		zap.Output(opts.output),
		zap.Filename(opts.filename),
		zap.MaxSize(opts.maxSize),
		zap.MaxBackups(opts.maxBackups),
		zap.MaxAge(opts.maxAge),
		zap.Compress(opts.compress),
	)
}

func Instance() Logger {
	return log
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Panic(args ...interface{}) {
	log.Panic(args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func SetLevel(level string) {
	log.SetLevel(level)
}
