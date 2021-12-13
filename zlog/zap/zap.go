/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/9/27 17:57
 */
package zap

import (
	"os"
	"strings"

	"github.com/go-kirito/pkg/zlog/writer"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logger struct {
	log   *zap.SugaredLogger
	level zap.AtomicLevel
}

func NewLogger(optsFunc ...OptionFunc) *logger {
	opts := options{
		format: "json",
		output: "console",
		level:  "debug",
	}

	for _, o := range optsFunc {
		o(&opts)
	}

	encoder := getEncoder(opts)
	writeSyncer := getWriteSyncer(opts)
	level := zap.NewAtomicLevelAt(getLevel(opts.level))

	core := zapcore.NewCore(encoder, writeSyncer, level)
	zn := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(2), zap.AddStacktrace(zapcore.WarnLevel))
	log := zn.Sugar()
	return &logger{log: log, level: level}
}

func getEncoder(opts options) zapcore.Encoder {

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "date"
	if opts.format == "json" {
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getWriteSyncer(opts options) zapcore.WriteSyncer {
	if opts.output == "console" {
		return zapcore.AddSync(os.Stderr)
	}

	if opts.output == "aliyun" {
		aliyunOpts := &writer.AliyunOption{
			AccessKey:       opts.accessKey,
			AccessKeySecret: opts.accessKeySecret,
			EndPoint:        opts.endPoint,
			Project:         opts.project,
			LogStore:        opts.logStore,
			Topic:           opts.topic,
		}
		aliyunWriter := writer.NewAliyunWriter(aliyunOpts)
		return zapcore.AddSync(aliyunWriter)
	}

	l := &lumberjack.Logger{
		Filename:   opts.filename,
		MaxSize:    opts.maxSize,
		MaxBackups: opts.maxBackups,
		MaxAge:     opts.maxAge,
		Compress:   opts.compress,
	}

	return zapcore.AddSync(l)
}

func getLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "panic":
		return zap.PanicLevel
	case "fatal":
		return zap.FatalLevel
	}
	return zap.DebugLevel
}

func (l *logger) SetLevel(level string) {
	lvl := getLevel(level)
	l.level.SetLevel(lvl)

}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.log.Debugf(format, args...)
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.log.Infof(format, args...)
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.log.Warnf(format, args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.log.Errorf(format, args...)
}

func (l *logger) Panicf(format string, args ...interface{}) {
	l.log.Panicf(format, args...)
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	l.log.Fatalf(format, args...)
}

func (l *logger) Debug(args ...interface{}) {
	l.log.Debug(args...)
}

func (l *logger) Info(args ...interface{}) {
	l.log.Info(args...)
}

func (l *logger) Warn(args ...interface{}) {
	l.log.Warn(args...)
}

func (l *logger) Error(args ...interface{}) {
	l.log.Error(args...)
}

func (l *logger) Panic(args ...interface{}) {
	l.log.Panic(args...)
}

func (l *logger) Fatal(args ...interface{}) {
	l.log.Fatal(args...)
}
