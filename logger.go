package myzerolog

import (
	"bufio"
	"bytes"
	"os"

	zl "github.com/rs/zerolog"
)

var Logger = MyNewConsoleLogger()

// Pre-populate buffer to print caller, to increase performance
func populateCallerBuffer(w int) []byte {
	buf := bytes.Repeat([]byte{32}, w)
	buf[0] = 91
	buf[w-1] = 93
	return buf
}

type myLogger struct {
	zLogger   *zl.Logger
	bufWriter *bufio.Writer

	defaultDepth int
	callerWidth  int
	callerBuffer []byte
}

func MyNewConsoleLogger() *myLogger {
	zl.CallerMarshalFunc = myCallerMarshalFunc
	zl.TimeFieldFormat = zl.TimeFormatUnixMs
	populateLevelMaps()

	l := myLogger{
		defaultDepth: 2,
		callerWidth:  16,
		bufWriter:    bufio.NewWriter(os.Stdout),
	}
	l.callerBuffer = populateCallerBuffer(l.callerWidth + 2)

	writer := zl.ConsoleWriter{
		Out:          l.bufWriter,
		TimeFormat:   "15:04:05.000",
		FormatLevel:  myDefaultFormatLevel(false),
		FormatCaller: l.myDefaultFormatCaller(false),
	}

	zLogger := zl.New(os.Stderr).With().Timestamp().Caller().Logger().Output(writer)
	l.zLogger = &zLogger
	return &l
}

func (l *myLogger) myDefaultFormatCaller(noColor bool) zl.Formatter {
	return func(i interface{}) string {
		var c string
		if cc, ok := i.(string); ok {
			c = cc
		}
		lc := len(c)
		if lc > 0 {
			//c = colorize(c, colorBold, noColor) + colorize(" |", colorCyan, noColor)
			buf := make([]byte, l.callerWidth+2)
			copy(buf, l.callerBuffer)
			if lc > l.callerWidth {
				buf[1] = 126
				copy(buf[2:], c[lc-l.callerWidth+1:lc])
			} else {
				copy(buf[l.callerWidth-lc+1:], c)
			}
			return string(buf)
		}
		return string(l.callerBuffer)
	}
}

// AdjustCallerWidth: adjust length of field dedicated for caller
func (l *myLogger) AdjustCallerWidth(w int) {
	l.callerWidth = w
	l.callerBuffer = populateCallerBuffer(w + 2)
}

// Flushes the buffered IO writer
func (l *myLogger) Sync() {
	l.bufWriter.Flush()
}

func (l *myLogger) emitMessage(lvl zl.Level, depth int, args []interface{}) {
	e := l.zLogger.WithLevel(lvl)
	e.CallerSkipFrame(depth)
	// setting 'level' key to numeric value to optimize decoding
	e.Int("level", int(lvl)).Msg(prepareMessage(args))
}

func (l *myLogger) emitFormattedMessage(lvl zl.Level, depth int, format string, args []interface{}) {
	e := l.zLogger.WithLevel(lvl)
	e.CallerSkipFrame(depth)
	// setting 'level' key to numeric value to optimize decoding
	e.Int("level", int(lvl)).Msg(prepareFormattedMessage(format, args))
}

/*
 *
 * Logger interface implementation
 *
 */

// Trace: log message with trace level.
func (l *myLogger) Trace(args ...interface{}) {
	l.emitMessage(zl.TraceLevel, l.defaultDepth, args)
}

// Tracef: format and log message with trace level.
func (l *myLogger) Tracef(format string, args ...interface{}) {
	l.emitFormattedMessage(zl.TraceLevel, l.defaultDepth, format, args)
}

// Debug: log message with debug level.
func (l *myLogger) Debug(args ...interface{}) {
	l.emitMessage(zl.DebugLevel, l.defaultDepth, args)
}

// Debugf: format and log message with debug level.
func (l *myLogger) Debugf(format string, args ...interface{}) {
	l.emitFormattedMessage(zl.DebugLevel, l.defaultDepth, format, args)
}

// Info: log message with info level.
func (l *myLogger) Info(args ...interface{}) {
	l.emitMessage(zl.InfoLevel, l.defaultDepth, args)
}

// Infof: format and log message with info level.
func (l *myLogger) Infof(format string, args ...interface{}) {
	l.emitFormattedMessage(zl.InfoLevel, l.defaultDepth, format, args)
}

// Warn: log message with warn level.
func (l *myLogger) Warn(args ...interface{}) {
	l.emitMessage(zl.WarnLevel, l.defaultDepth, args)
}

// Warnf: format and log message with warn level.
func (l *myLogger) Warnf(format string, args ...interface{}) {
	l.emitFormattedMessage(zl.WarnLevel, l.defaultDepth, format, args)
}

// Error: log message with error level.
func (l *myLogger) Error(args ...interface{}) {
	l.emitMessage(zl.ErrorLevel, l.defaultDepth, args)
}

// Errorf: format and log message with error level.
func (l *myLogger) Errorf(format string, args ...interface{}) {
	l.emitFormattedMessage(zl.ErrorLevel, l.defaultDepth, format, args)
}

// Fatal: log message with fatal level.
func (l *myLogger) Fatal(args ...interface{}) {
	l.emitMessage(zl.FatalLevel, l.defaultDepth, args)
}

// Fatalf: format and log message with fatal level.
func (l *myLogger) Fatalf(format string, args ...interface{}) {
	l.emitFormattedMessage(zl.FatalLevel, l.defaultDepth, format, args)
}

// Panic: log message with panic level.
func (l *myLogger) Panic(args ...interface{}) {
	l.emitMessage(zl.PanicLevel, l.defaultDepth, args)
}

// Panicf: format and log message with panic level.
func (l *myLogger) Panicf(format string, args ...interface{}) {
	l.emitFormattedMessage(zl.PanicLevel, l.defaultDepth, format, args)
}
