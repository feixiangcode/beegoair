package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path/filepath"
	"strings"
	"time"
)

var logger *Logger

type Logger struct {
	logPath     string
	logFile     string
	log         *zap.Logger
	currentDate string
	isFileLog   bool
	level       zapcore.Level
}

func formatEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("%d%02d%02d_%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()))
}

func (this *Logger) InitZapLog() {
	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(this.level),
		Development: true,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "t",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "trace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     formatEncodeTime,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{this.logFile},
		ErrorOutputPaths: []string{this.logFile},
		InitialFields: map[string]interface{}{
			"lang": "go",
		},
	}
	var err error
	this.log, err = cfg.Build()
	if err != nil {
		panic("log init fail:" + err.Error())
	}
}

func (this *Logger) GetLevel(l string) zapcore.Level {
	switch strings.ToLower(l) {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	}
	return zap.WarnLevel
}

func (this *Logger) UpdateLogFile() {
	t := time.Now()
	if this.isFileLog && t.Format("20060102") != this.currentDate {
		this.currentDate = t.Format("20060102")
		this.logFile = filepath.Join(this.logPath, this.currentDate+".log")
	}
}

func (this *Logger) FormateLog(args []interface{}) *zap.Logger {
	this.UpdateLogFile()
	log := this.log.With(this.ToJsonData(args))
	return log
}

func (this *Logger) ToJsonData(args []interface{}) zap.Field {
	det := make([]string, 0)
	if len(args) > 0 {
		for _, v := range args {
			det = append(det, fmt.Sprintf("%+v", v))
		}
	}
	zap := zap.Any("detail", det)
	return zap
}

func Debug(msg string, args ...interface{}) {
	logger.FormateLog(args).Sugar().Debugf(msg)
}

func Info(msg string, args ...interface{}) {
	logger.FormateLog(args).Sugar().Infof(msg)
}

func Warn(msg string, args ...interface{}) {
	logger.FormateLog(args).Sugar().Warnf(msg)
}

func Error(msg string, args ...interface{}) {
	logger.FormateLog(args).Sugar().Errorf(msg)
}

func Fatal(msg string, args ...interface{}) {
	logger.FormateLog(args).Sugar().Fatalf(msg)
}

/*func init() {
	logger = new(Logger)
	logger.isFileLog,_ = beego.AppConfig.Bool("isfile")
	logger.logPath = beego.AppConfig.String("logpath")
	logger.UpdateLogFile()
	if !logger.isFileLog {
		logger.logFile = "stdout"
	}
	logger.level = logger.GetLevel(beego.AppConfig.String("loglevel"))
	logger.InitZapLog()
}*/

type Cfg struct {
	IsFileLog bool
	LogPath   string
	Level     string
}

func InitLog(cfg *Cfg) {
	logger = new(Logger)
	logger.isFileLog = cfg.IsFileLog
	logger.logPath = cfg.LogPath
	if logger.isFileLog {
		logger.UpdateLogFile()
	} else {
		logger.logFile = "stdout"
	}
	logger.level = logger.GetLevel(cfg.Level)
	logger.InitZapLog()
}
