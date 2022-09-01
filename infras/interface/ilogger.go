package interfaces

//go:generate mockgen -destination=../../test/mock/ilogger_mock_infras.go -package=mock GoWeb/infras/interface IApiLogger

// Logger methods interface
type IApiLogger interface {
	// init logger
	InitLogger()
	// debug 使用 fmt.Sprint 建構和記錄訊息
	// @param args
	Debug(args ...interface{})
	// debugf 使用 fmt.Sprintf 記錄template訊息
	// @param template
	// @param args
	Debugf(template string, args ...interface{})
	// info 使用 fmt.Sprint 建構和記錄訊息
	// @param args
	Info(args ...interface{})
	// infof 使用 fmt.Sprintf 記錄template訊息
	// @param template
	// @param args
	Infof(template string, args ...interface{})
	// infow logs a message with some additional context.
	// The variadic key-value pairs are treated as they are in With.
	// @param template
	// @param keysAndValues
	Infow(template string, keysAndValues ...interface{})
	// warn 使用 fmt.Sprint 建構和記錄訊息
	// @param args
	Warn(args ...interface{})
	// warnf 使用 fmt.Sprintf 記錄模板訊息。
	// @param template
	// @param args
	Warnf(template string, args ...interface{})
	// error 使用 fmt.Sprint 建構和記錄訊息
	// @param args
	Error(args ...interface{})
	// errorf 使用 fmt.Sprintf 記錄模板訊息
	// @param template
	// @param args
	Errorf(template string, args ...interface{})
	// dpanic 使用 fmt.Sprint 建構和記錄訊息
	// 代表"panic in development." 在development中，它會列印Panic級別的日誌：反之，它將發生在Error級別的日誌
	// @param args
	DPanic(args ...interface{})
	// dpanicf 使用 fmt.Sprintf 記錄template訊息
	// @param template
	// @param args
	DPanicf(template string, args ...interface{})
	// fatal 使用 fmt.Sprint 建構和記錄訊息
	// @param args
	Fatal(args ...interface{})
	// fatalf 使用 fmt.Sprintf 記錄template訊息
	// @param template
	// @param args
	Fatalf(template string, args ...interface{})
}
