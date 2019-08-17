package log

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

// LogType is type used for log type.
type LogType string

const (
	// AppLog is constant used for application debugging logs.
	AppLog LogType = "APPLOG"
)

// Fields type is used to pass to `WithFields`.
type Fields = logrus.Fields

type Level uint32

// Convert the Level to a string. E.g. PanicLevel becomes "panic".
func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warning"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	case PanicLevel:
		return "panic"
	}

	return "unknown"
}

// ParseLevel takes a string level and returns the log level constant.
func ParseLevel(lvl string) (Level, error) {
	switch strings.ToLower(lvl) {
	case "panic":
		return PanicLevel, nil
	case "fatal":
		return FatalLevel, nil
	case "error":
		return ErrorLevel, nil
	case "warn", "warning":
		return WarnLevel, nil
	case "info":
		return InfoLevel, nil
	case "debug":
		return DebugLevel, nil
	}

	var l Level
	return l, fmt.Errorf("not a valid Level: %q", lvl)
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (level *Level) UnmarshalText(text []byte) error {
	l, err := ParseLevel(string(text))
	if err != nil {
		return err
	}

	*level = Level(l)

	return nil
}

// A constant exposing all logging levels
var AllLevels = []Level{
	PanicLevel,
	FatalLevel,
	ErrorLevel,
	WarnLevel,
	InfoLevel,
	DebugLevel,
}

// These are the different logging levels. You can set the logging level to log
// on your instance of logger, obtained with `logrus.New()`.
const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
)

// Logger is the interface for loggers used in atom apps.
type Logger interface {
	Debug(...interface{})
	Debugf(string, ...interface{})
	Info(...interface{})
	Infof(string, ...interface{})
	Warn(...interface{})
	Warnf(string, ...interface{})
	Error(...interface{})
	Errorf(string, ...interface{})
	WithFields(fields Fields) Logger
	SetOutput(w io.Writer)
	SetLevel(Level)
}

type logger struct {
	entry *logrus.Entry
}

// New creates and return logger. This logger use a simple text
// formatter. The default output will be to os.Stderr. This
// logger will also have all entries appended with the provided
// LogType.
func New(t LogType) Logger {
	return logger{
		entry: &logrus.Entry{
			Logger: &logrus.Logger{
				Out: os.Stderr,
				Formatter: &logrus.TextFormatter{
					DisableColors: true,
					FullTimestamp: true,
				},
				Hooks:        make(logrus.LevelHooks),
				Level:        logrus.InfoLevel,
				ExitFunc:     os.Exit,
				ReportCaller: false,
			},
			Data: logrus.Fields{"type": t},
		},
	}
}

// Debug logs a message at level Debug on the standard logger.
func (l logger) Debug(args ...interface{}) {
	l.entry.Debug(args...)
}

// Debugf logs a message at level Debug on the standard logger.
func (l logger) Debugf(format string, args ...interface{}) {
	l.entry.Debugf(format, args...)
}

// Info logs a message at level Info on the standard logger.
func (l logger) Info(args ...interface{}) {
	l.entry.Info(args...)
}

// Infof logs a message at level Info on the standard logger.
func (l logger) Infof(format string, args ...interface{}) {
	l.entry.Infof(format, args...)
}

// Warn logs a message at level Warn on the standard logger.
func (l logger) Warn(args ...interface{}) {
	l.entry.Warn(args...)
}

// Warnf logs a message at level Warn on the standard logger.
func (l logger) Warnf(format string, args ...interface{}) {
	l.entry.Warnf(format, args...)
}

// Error logs a message at level Error on the standard logger.
func (l logger) Error(args ...interface{}) {
	l.entry.Error(args...)
}

// Errorf logs a message at level Error on the standard logger.
func (l logger) Errorf(format string, args ...interface{}) {
	l.entry.Errorf(format, args...)
}

// WithFields adds additional fields to the log message. Field
// "type" cannot be used by users. Duplicate fields will also
// be removed. If you need a different type log use
// log.New(<LogType>).
func (l logger) WithFields(fields Fields) Logger {
	if _, ok := fields["type"]; ok {
		l.Errorf("Field \"type\" cannot be assigned by users.  Discarding")
		delete(fields, "type")
	}
	return logger{entry: l.entry.WithFields(fields)}
}

// SetOutput sets the output to desired io.Writer like stdout, stderr etc
func (l logger) SetOutput(w io.Writer) {
	l.entry.Logger.Out = w
}

// SetLevel sets the logger level for emitting the log entry.
func (l logger) SetLevel(level Level) {
	l.entry.Logger.SetLevel(logrus.Level(level))
}

func (lt LogType) String() string {
	return string(lt)
}
