/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2021/3/18 14:33
 */
package zlog

import (
	log2 "log"

	"github.com/go-kirito/pkg/zlog/zap"
)

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

func NewLogger(optsFunc ...OptionFunc) Logger {

	opts := NewOptionsWithConfig()

	for _, o := range optsFunc {
		o(&opts)
	}

	return zap.NewLogger(
		zap.Format(opts.Format),
		zap.Level(opts.Level),
		zap.Output(opts.Output),
		zap.Filename(opts.Filename),
		zap.MaxSize(opts.MaxSize),
		zap.MaxBackups(opts.MaxBackups),
		zap.MaxAge(opts.MaxAge),
		zap.Compress(opts.Compress),
		zap.AccessKey(opts.AccessKey),
		zap.AccessKeySecret(opts.AccessKeySecret),
		zap.EndPoint(opts.EndPoint),
		zap.Project(opts.Project),
		zap.LogStore(opts.LogStore),
		zap.Topic(opts.Topic),
	)
}

func NewLoggerWithOption(optsFunc ...OptionFunc) Logger {
	opts := options{
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
	for _, o := range optsFunc {
		o(&opts)
	}
	log2.Println(opts.Output)

	return zap.NewLogger(
		zap.Format(opts.Format),
		zap.Level(opts.Level),
		zap.Output(opts.Output),
		zap.Filename(opts.Filename),
		zap.MaxSize(opts.MaxSize),
		zap.MaxBackups(opts.MaxBackups),
		zap.MaxAge(opts.MaxAge),
		zap.Compress(opts.Compress),
		zap.AccessKey(opts.AccessKey),
		zap.AccessKeySecret(opts.AccessKeySecret),
		zap.EndPoint(opts.EndPoint),
		zap.Project(opts.Project),
		zap.LogStore(opts.LogStore),
		zap.Topic(opts.Topic),
	)
}

func Init() {
	log = NewLogger()
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
