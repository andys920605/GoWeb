package logger

import (
	"os"

	"GoWeb/infras/configs"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// region members
// For mapping config logger to app logger levels
var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,  // Debug 應是大量的，且通常在生產狀態禁用.
	"info":   zapcore.InfoLevel,   // Info 是預設的記錄優先順序
	"warn":   zapcore.WarnLevel,   // Warn 比 info 更重要
	"error":  zapcore.ErrorLevel,  // Error 是高優先順序的,如果程式順利不應該產生任何 err 級別日誌
	"dpanic": zapcore.DPanicLevel, // DPanic 特別重大的錯誤，在開發模式下引起 panic.
	"panic":  zapcore.PanicLevel,  // Panic 記錄訊息後呼叫 panic.
	"fatal":  zapcore.FatalLevel,  // Fatal 記錄訊息後呼叫 os.Exit(1)
}

// endregion

// region constructor
// api Logger
type ApiLogger struct {
	cfg         *configs.Config
	sugarLogger *zap.SugaredLogger
}

// App Logger constructor
func NewApiLogger(cfg *configs.Config) *ApiLogger {
	return &ApiLogger{cfg: cfg}
}

// endregion

// Init logger
func (log *ApiLogger) InitLogger() {
	logLevel := log.getLoggerLevel(log.cfg)
	var syncWriter zapcore.WriteSyncer
	if log.cfg.Logger.EnableFile {
		syncWriter = zapcore.AddSync(&lumberjack.Logger{
			Filename:   log.cfg.Logger.InfoFileName, // ⽇誌檔案路徑
			MaxSize:    log.cfg.Logger.MaxSize,      // 1M=1024KB=1024000byte; 1G:1<<30
			MaxAge:     log.cfg.Logger.MaxAge,       // days
			MaxBackups: log.cfg.Logger.MaxBackups,   // 最多保留5個備份
			LocalTime:  true,                        // The default is to use UTC time.
			Compress:   true,                        // 是否壓縮 disabled by default
		})
	} else {
		syncWriter = zapcore.AddSync(os.Stderr)
	}

	var encoderCfg zapcore.EncoderConfig
	if log.cfg.Logger.Development {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}
	encoderCfg.LevelKey = "LEVEL"
	encoderCfg.CallerKey = "CALLER"
	encoderCfg.TimeKey = "TIME"
	encoderCfg.NameKey = "NAME"
	encoderCfg.MessageKey = "MESSAGE"

	var encoder zapcore.Encoder
	if log.cfg.Logger.Encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(encoder, syncWriter, zap.NewAtomicLevelAt(logLevel))
	logger := zap.New(core, zap.WithCaller(!log.cfg.Logger.DisableCaller), zap.AddCallerSkip(1))

	log.sugarLogger = logger.Sugar()
	if err := log.sugarLogger.Sync(); err != nil {
		log.sugarLogger.Error(err)
	}
}

// region public methods
// debug 使用 fmt.Sprint 建構和記錄訊息
// @param args
func (log *ApiLogger) Debug(args ...interface{}) {
	log.sugarLogger.Debug(args...)
}

// debugf 使用 fmt.Sprintf 記錄template訊息
// @param template
// @param args
func (log *ApiLogger) Debugf(template string, args ...interface{}) {
	log.sugarLogger.Debugf(template, args...)
}

// info 使用 fmt.Sprint 建構和記錄訊息
// @param args
func (log *ApiLogger) Info(args ...interface{}) {
	log.sugarLogger.Info(args...)
}

// infof 使用 fmt.Sprintf 記錄template訊息
// @param template
// @param args
func (log *ApiLogger) Infof(template string, args ...interface{}) {
	log.sugarLogger.Infof(template, args...)
}

// infow logs a message with some additional context.
// The variadic key-value pairs are treated as they are in With.
// @param template
// @param keysAndValues
func (log *ApiLogger) Infow(template string, keysAndValues ...interface{}) {
	log.sugarLogger.Infow(template, keysAndValues...)
} // Austin 20220614

// warn 使用 fmt.Sprint 建構和記錄訊息
// @param args
func (log *ApiLogger) Warn(args ...interface{}) {
	log.sugarLogger.Warn(args...)
}

// warnf 使用 fmt.Sprintf 記錄template訊息
// @param template
// @param args
func (log *ApiLogger) Warnf(template string, args ...interface{}) {
	log.sugarLogger.Warnf(template, args...)
}

// error 使用 fmt.Sprint 建構和記錄訊息
// @param args
func (log *ApiLogger) Error(args ...interface{}) {
	log.sugarLogger.Error(args...)
}

// errorf 使用 fmt.Sprintf 記錄template訊息
// @param template
// @param args
func (log *ApiLogger) Errorf(template string, args ...interface{}) {
	log.sugarLogger.Errorf(template, args...)
}

// dpanic 使用 fmt.Sprint 建構和記錄訊息
// @param args
func (log *ApiLogger) DPanic(args ...interface{}) {
	log.sugarLogger.DPanic(args...)
}

// dpanicf 使用 fmt.Sprintf 記錄template訊息
// @param template
// @param args
func (log *ApiLogger) DPanicf(template string, args ...interface{}) {
	log.sugarLogger.DPanicf(template, args...)
}

// panic 使用 fmt.Sprint 建構和記錄訊息
// @param args
func (log *ApiLogger) Panic(args ...interface{}) {
	log.sugarLogger.Panic(args...)
}

// panicf 使用 fmt.Sprintf 記錄template訊息
// @param template
// @param args
func (log *ApiLogger) Panicf(template string, args ...interface{}) {
	log.sugarLogger.Panicf(template, args...)
}

// fatal 使用 fmt.Sprint 建構和記錄訊息
// @param args
func (log *ApiLogger) Fatal(args ...interface{}) {
	log.sugarLogger.Fatal(args...)
}

// fatalf 使用 fmt.Sprintf 記錄template訊息
// @param template
// @param args
func (log *ApiLogger) Fatalf(template string, args ...interface{}) {
	log.sugarLogger.Fatalf(template, args...)
}

// endregion

// region private methods
// get logger level
// @param config
// @result zapcore.Level
func (log *ApiLogger) getLoggerLevel(cfg *configs.Config) zapcore.Level {
	level, exist := loggerLevelMap[cfg.Logger.Level]
	if !exist {
		return zapcore.DebugLevel
	}
	return level
}

// endregion
