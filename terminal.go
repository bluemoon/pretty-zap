package prettyzap

import "go.uber.org/zap/zapcore"

// BrightColor represents a text color.
type EscapeCodes string

// Foreground colors.
const (
	Black   EscapeCodes = "\x1b[30m"
	Red                 = "\x1b[31m"
	Green               = "\x1b[32m"
	Yellow              = "\x1b[33m"
	Blue                = "\x1b[34m"
	Magenta             = "\x1b[35m"
	Cyan                = "\x1b[36m"
	White               = "\x1b[37m"
	Clear               = "\x1b[0m"
	Bold                = "\x1b[1m"
)

var (
	levelColors = map[zapcore.Level]EscapeCodes{
		zapcore.DebugLevel:  Green,
		zapcore.InfoLevel:   Cyan,
		zapcore.WarnLevel:   Yellow,
		zapcore.ErrorLevel:  Red,
		zapcore.PanicLevel:  Red,
		zapcore.FatalLevel:  Red,
		zapcore.DPanicLevel: Red, // uber's development panic level
	}
	unknownLevelColor EscapeCodes = Red
)

// ColorKey will color the provided key at the associated log level color
func ColorKey(level zapcore.Level) string {
	c, ok := levelColors[level]
	if !ok {
		c = unknownLevelColor
	}
	return string(c)
}
