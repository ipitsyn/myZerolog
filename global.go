package myzerolog

import zl "github.com/rs/zerolog"

// AdjustCallerWidth: adjust length of field dedicated for caller
func AdjustCallerWidth(w int) {
	Logger.AdjustCallerWidth(w)
}

// Flushes the buffered IO writer
func Sync() {
	Logger.Sync()
}

/*
 *
 * Global Logger interface implementation
 *
 */

// Trace: log message with trace level.
func Trace(args ...interface{}) {
	Logger.emitMessage(zl.TraceLevel, Logger.defaultDepth, args)
}

// Tracef: format and log message with trace level.
func Tracef(format string, args ...interface{}) {
	Logger.emitFormattedMessage(zl.TraceLevel, Logger.defaultDepth, format, args)
}

// Debug: log message with debug level.
func Debug(args ...interface{}) {
	Logger.emitMessage(zl.DebugLevel, Logger.defaultDepth, args)
}

// Debugf: format and log message with debug level.
func Debugf(format string, args ...interface{}) {
	Logger.emitFormattedMessage(zl.DebugLevel, Logger.defaultDepth, format, args)
}

// Info: log message with info level.
func Info(args ...interface{}) {
	Logger.emitMessage(zl.InfoLevel, Logger.defaultDepth, args)
}

// Infof: format and log message with info level.
func Infof(format string, args ...interface{}) {
	Logger.emitFormattedMessage(zl.InfoLevel, Logger.defaultDepth, format, args)
}

// Warn: log message with warn level.
func Warn(args ...interface{}) {
	Logger.emitMessage(zl.WarnLevel, Logger.defaultDepth, args)
}

// Warnf: format and log message with warn level.
func Warnf(format string, args ...interface{}) {
	Logger.emitFormattedMessage(zl.WarnLevel, Logger.defaultDepth, format, args)
}

// Error: log message with error level.
func Error(args ...interface{}) {
	Logger.emitMessage(zl.ErrorLevel, Logger.defaultDepth, args)
}

// Errorf: format and log message with error level.
func Errorf(format string, args ...interface{}) {
	Logger.emitFormattedMessage(zl.ErrorLevel, Logger.defaultDepth, format, args)
}

// Fatal: log message with fatal level.
func Fatal(args ...interface{}) {
	Logger.emitMessage(zl.FatalLevel, Logger.defaultDepth, args)
}

// Fatalf: format and log message with fatal level.
func Fatalf(format string, args ...interface{}) {
	Logger.emitFormattedMessage(zl.FatalLevel, Logger.defaultDepth, format, args)
}

// Panic: log message with panic level.
func Panic(args ...interface{}) {
	Logger.emitMessage(zl.PanicLevel, Logger.defaultDepth, args)
}

// Panicf: format and log message with panic level.
func Panicf(format string, args ...interface{}) {
	Logger.emitFormattedMessage(zl.PanicLevel, Logger.defaultDepth, format, args)
}
