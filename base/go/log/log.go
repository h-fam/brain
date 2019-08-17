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
func (l *Level) UnmarshalText(text []byte) error {
	lvl, err := ParseLevel(string(text))
	if err != nil {
		return err
	}
	*l = Level(lvl)
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

// Logger is the base logger used in Atom apps.
type Logger struct {
	entry *logrus.Entry
}

// New creates a new logger. This logger uses a simple text
// formatter. The default output will be to os.Stderr. This
// logger will also have all entries appended with the provided
// LogType.
func New(t LogType) Logger {
	return Logger{
		entry: &logrus.Entry{
			Logger: &logrus.Logger{
				Out: os.Stderr,
				Formatter: &logrus.TextFormatter{
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
func (l Logger) Debug(args ...interface{}) {
	l.entry.Debug(args...)
}

// Debugf logs a message at level Debug on the standard logger.
func (l Logger) Debugf(format string, args ...interface{}) {
	l.entry.Debugf(format, args...)
}

// Info logs a message at level Info on the standard logger.
func (l Logger) Info(args ...interface{}) {
	l.entry.Info(args...)
}

// Infof logs a message at level Info on the standard logger.
func (l Logger) Infof(format string, args ...interface{}) {
	l.entry.Infof(format, args...)
}

// Warn logs a message at level Warn on the standard logger.
func (l Logger) Warn(args ...interface{}) {
	l.entry.Warn(args...)
}

// Warnf logs a message at level Warn on the standard logger.
func (l Logger) Warnf(format string, args ...interface{}) {
	l.entry.Warnf(format, args...)
}

// Error logs a message at level Error on the standard logger.
func (l Logger) Error(args ...interface{}) {
	l.entry.Error(args...)
}

// Errorf logs a message at level Error on the standard logger.
func (l Logger) Errorf(format string, args ...interface{}) {
	l.entry.Errorf(format, args...)
}

// WithFields adds additional fields to the log message. Field
// "type" cannot be used by users. Duplicate fields will also
// be removed. If you need a different type log use
// log.New(<LogType>).
func (l Logger) WithFields(fields Fields) Logger {
	if _, ok := fields["type"]; ok {
		l.Errorf("Field \"type\" cannot be assigned by users.  Discarding")
		delete(fields, "type")
	}
	return Logger{entry: l.entry.WithFields(fields)}
}

// SetOutput sets the output to desired io.Writer like stdout, stderr etc
func (l Logger) SetOutput(w io.Writer) {
	l.entry.Logger.Out = w
}

// SetLevel sets the logger level for emitting the log entry.
func (l Logger) SetLevel(level Level) {
	l.entry.Logger.SetLevel(logrus.Level(level))
}

func (lt LogType) String() string {
	return string(lt)
}

var defaultLogger = New(AppLog)

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

// WithFields adds additional fields to the log message. Field
// "type" cannot be used by users. Duplicate fields will also
// be removed. If you need a different type log use
// log.New(<LogType>).
func WithFields(fields Fields) Logger {
	return defaultLogger.WithFields(fields)
}

// SetOutput sets the output to desired io.Writer like stdout, stderr etc
func SetOutput(w io.Writer) {
	defaultLogger.SetOutput(w)
}

// SetLevel sets the logger level for emitting the log entry.
func SetLevel(level Level) {
	defaultLogger.SetLevel(level)
}
