package myzerolog

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	zl "github.com/rs/zerolog"
)

const (
	colorBlack = iota + 30
	colorRed
	colorGreen
	colorYellow
	colorBlue
	colorMagenta
	colorCyan
	colorWhite

	colorBold     = 1
	colorDarkGray = 90
)

type (
	levelNumColorMap map[zl.Level]map[bool]string
	levelStrColorMap map[string]map[bool]string
)

var numLevelColors, strLevelColors = populateLevelMaps()

// for Caller display base file name only
var myCallerMarshalFunc = func(pc uintptr, file string, line int) string {
	return fmt.Sprintf("%s:%03d",
		filepath.Base(file),
		line,
	)
}

// colorize returns the string s wrapped in ANSI code c, unless disabled is true.
func colorize(s interface{}, c int, disabled bool) string {
	if disabled {
		return fmt.Sprintf("%s", s)
	}
	return fmt.Sprintf("\x1b[%dm%v\x1b[0m", c, s)
}

// utility function to iterate over levels and apply function
func forEachLevel(f func(lvl zl.Level)) {
	for _, l := range []zl.Level{
		zl.TraceLevel,
		zl.DebugLevel,
		zl.InfoLevel,
		zl.WarnLevel,
		zl.ErrorLevel,
		zl.FatalLevel,
		zl.PanicLevel,
	} {
		f(l)
	}
}

// pre-populate logging level values to increase perofrmanse
func populateLevelMaps() (levelNumColorMap, levelStrColorMap) {
	numMap := levelNumColorMap{}
	strMap := levelStrColorMap{}

	forEachLevel(func(lvl zl.Level) {
		numMap[lvl] = map[bool]string{}
	})

	for _, noColor := range []bool{true, false} {
		numMap[zl.TraceLevel][noColor] = colorize("TRC", colorDarkGray, noColor)
		numMap[zl.DebugLevel][noColor] = colorize("DBG", colorMagenta, noColor)
		numMap[zl.InfoLevel][noColor] = colorize("INF", colorCyan, noColor)
		numMap[zl.WarnLevel][noColor] = colorize("WRN", colorYellow, noColor)
		numMap[zl.ErrorLevel][noColor] = colorize("ERR", colorRed, noColor)
		numMap[zl.FatalLevel][noColor] = colorize(colorize("FTL", colorRed, noColor), colorBold, noColor)
		numMap[zl.PanicLevel][noColor] = colorize(colorize("PNC", colorRed, noColor), colorBold, noColor)
	}

	forEachLevel(func(lvl zl.Level) {
		strMap[lvl.String()] = numMap[lvl]
	})
	return numMap, strMap
}

// Level formatter
func myDefaultFormatLevel(noColor bool) zl.Formatter {
	return func(i interface{}) string {
		var l string
		switch ll := i.(type) {
		case json.Number:
			n, _ := ll.Int64()
			if m, ok := numLevelColors[zl.Level(n)]; ok {
				l = m[noColor]
			} else {
				l = colorize("???", colorBold, noColor)
			}
		case string:
			if m, ok := strLevelColors[ll]; ok {
				l = m[noColor]
			} else {
				l = colorize("???", colorBold, noColor)
			}
		default:
			fmt.Printf("%T\n", i)
			if i == nil {
				l = colorize("???", colorBold, noColor)
			} else {
				l = strings.ToUpper(fmt.Sprintf("%v", i))
			}
		}
		return l
	}
}
