package log

import (
	"io"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

// LogType is type used for log type.
type LogType string

const (
	// AppLog is constant used for app debugging logs.
	AppLog LogType = "APPLOG"
)

// Fields type is used to pass to `WithFields`.
type Fields = logrus.Fields

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
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	WithFields(fields Fields) Logger
	SetOutput(w io.Writer)
}

type logger struct {
	entry *logrus.Entry
}

// New creates and return logger.
func New() Logger {
	var baseLogger = logger{entry: logrus.NewEntry(logrus.New())}
	// Setting default formatter as TextFormatter.
	baseLogger.entry.Logger.Formatter = &logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	}
	// Setting default level as debug level.
	baseLogger.entry.Logger.Level = logrus.DebugLevel
	// Setting default out as os.Stderr.
	baseLogger.entry.Logger.Out = os.Stderr
	return baseLogger
}

// Debug logs a message at level Debug on the standard logger.
func (l logger) Debug(args ...interface{}) {
	l.sourced(AppLog.toString()).Debug(args...)
}

// Debugf logs a message at level Debug on the standard logger.
func (l logger) Debugf(format string, args ...interface{}) {
	l.sourced(AppLog.toString()).Debugf(format, args...)
}

// Info logs a message at level Info on the standard logger.
func (l logger) Info(args ...interface{}) {
	l.sourced(AppLog.toString()).Info(args...)
}

// Infof logs a message at level Info on the standard logger.
func (l logger) Infof(format string, args ...interface{}) {
	l.sourced(AppLog.toString()).Infof(format, args...)
}

// Warn logs a message at level Warn on the standard logger.
func (l logger) Warn(args ...interface{}) {
	l.sourced(AppLog.toString()).Warn(args...)
}

// Warnf logs a message at level Warn on the standard logger.
func (l logger) Warnf(format string, args ...interface{}) {
	l.sourced(AppLog.toString()).Warnf(format, args...)
}

// Error logs a message at level Error on the standard logger.
func (l logger) Error(args ...interface{}) {
	l.sourced(AppLog.toString()).Error(args...)
}

// Errorf logs a message at level Error on the standard logger.
func (l logger) Errorf(format string, args ...interface{}) {
	l.sourced(AppLog.toString()).Errorf(format, args...)
}

// Log logs after setting approriate logger output and level.
func (l logger) Log(args ...interface{}) {
	level := l.level()
	l.entry.Log(level, args...)
}

// Logf logs after setting approriate logger output and level.
func (l logger) Logf(format string, args ...interface{}) {
	level := l.level()
	l.entry.Logf(level, format, args...)
}

// WithFields takes into consideration the fields
func (l logger) WithFields(fields Fields) Logger {
	ltype := AppLog.toString()
	if _, ok := fields["type"]; ok {
		if ltype, ok = fields["type"].(string); !ok {
			l.Errorf("wrong log type value passed %v, it should be string , changing the log type to APPLOG", l.entry.Data["type"])
			ltype = AppLog.toString()
		}
	}
	fields["type"] = strings.ToUpper(ltype)
	return logger{l.sourced(ltype).WithFields(fields)}
}

// SetOutput sets the output to desired io.Writer like stdout, stderr etc
func (l logger) SetOutput(w io.Writer) {
	l.entry.Logger.Out = w
}

// sourced adds a type field to the entry data.
func (l logger) sourced(logType string) *logrus.Entry {
	return l.entry.WithField("type", logType)
}

func (lt Logtype) toString() string {
	return string(lt)
}

func (l logger) level() logrus.Level {
	ll := "INFO"
	if _, ok := l.entry.Data["level"]; ok {
		if ll, ok = l.entry.Data["level"].(string); !ok {
			l.Errorf("wrong level value passed %v , it should be string changing the log level to info", l.entry.Data["level"])
			ll = "INFO"
		}
	}
	level, err := logrus.ParseLevel(ll)
	if err != nil {
		level = logrus.InfoLevel
	}
	return level
}

