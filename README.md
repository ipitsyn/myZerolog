# myZeroLog

## An attempt to implement Go console logger by wrapping `zerolog` module.

Implements commonly used `Logger` interface methods:
- Trace / Tracef
- Debug / Debugf
- Info / Infof
- Warn / Warnf
- Error / Errorf
- Panic / Panicf
- Fatal / Fatalf

## NB! Test results
It is actually slower than the `zap` sugared logger :-(  
Even with the buffered io.Writer. Most likely because `zerolog`'s ConsoleWriter produce the line to be written from the JSON, so assembling / disassembling JSON is what brings overhead here.

```
15:44:09.277 WRN [         test.go:090] myZerolog speed: 121321.515 operations per second
15:44:09.277 WRN [         test.go:091] ZapSugar speed: 147570.394 operations per second
```