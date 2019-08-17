package log

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLoggerDebug(t *testing.T) {
	var buffer bytes.Buffer
	log := fakeLogger()
	log.SetLevel(DebugLevel)
	log.SetOutput(&buffer)
	tests := []struct {
		desc string
		l    Logger
		args []interface{}
		want string
	}{{
		desc: "empty msg debug log case",
		l:    log,
		want: `level=debug type=APPLOG`,
	}, {
		desc: "debug log with msg case",
		args: []interface{}{"debug message"},
		l:    log,
		want: `level=debug msg="debug message" type=APPLOG`,
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			tt.l.Debug(tt.args...)
			output := buffer.String()
			buffer.Reset()
			if strings.TrimRight(output, "\n\r") != tt.want {
				t.Errorf("Debug(%v) failed: got %v, want %v", tt.args, output, tt.want)
			}
		})
	}
}

func TestLoggerDebugf(t *testing.T) {
	var buffer bytes.Buffer
	log := fakeLogger()
	log.SetLevel(DebugLevel)
	log.SetOutput(&buffer)
	tests := []struct {
		desc   string
		l      Logger
		format string
		args   []interface{}
		want   string
	}{{
		desc: "empty debugf log case",
		l:    log,
		want: `level=debug type=APPLOG`,
	}, {
		desc:   "debugf log with msg case",
		format: "debugf message %s",
		args:   []interface{}{"testing"},
		l:      log,
		want:   `level=debug msg="debugf message testing" type=APPLOG`,
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			tt.l.Debugf(tt.format, tt.args...)
			output := buffer.String()
			buffer.Reset()
			if strings.TrimRight(output, "\n\r") != tt.want {
				t.Errorf("Debugf(%v) failed: got %v, want %v", tt.args, output, tt.want)
			}
		})
	}
}

func TestLoggerInfo(t *testing.T) {
	var buffer bytes.Buffer
	log := fakeLogger()
	log.SetOutput(&buffer)
	tests := []struct {
		desc string
		l    Logger
		args []interface{}
		want string
	}{{
		desc: "empty info log case",
		l:    log,
		want: `level=info type=APPLOG`,
	}, {
		desc: "info log with msg case",
		args: []interface{}{"info message"},
		l:    log,
		want: `level=info msg="info message" type=APPLOG`,
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			tt.l.Info(tt.args...)
			output := buffer.String()
			buffer.Reset()
			if strings.TrimRight(output, "\n\r") != tt.want {
				t.Errorf("Info(%v) failed: got %v, %v", tt.args, output, tt.want)
			}
		})
	}
}

func TestLoggerInfof(t *testing.T) {
	var buffer bytes.Buffer
	log := fakeLogger()
	log.SetOutput(&buffer)
	tests := []struct {
		desc   string
		l      Logger
		format string
		args   []interface{}
		want   string
	}{{
		desc: "empty infof log case",
		l:    log,
		want: `level=info type=APPLOG`,
	}, {
		desc:   "infof log with msg case",
		format: "infof message %s",
		args:   []interface{}{"testing"},
		l:      log,
		want:   `level=info msg="infof message testing" type=APPLOG`,
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			tt.l.Infof(tt.format, tt.args...)
			output := buffer.String()
			buffer.Reset()
			if strings.TrimRight(output, "\n\r") != tt.want {
				t.Errorf("Infof(%v) failed: got %v, %v", tt.args, output, tt.want)
			}
		})
	}
}

func TestLoggerWarn(t *testing.T) {
	var buffer bytes.Buffer
	log := fakeLogger()
	log.SetOutput(&buffer)
	tests := []struct {
		desc string
		l    Logger
		args []interface{}
		want string
	}{{
		desc: "empty warn log case",
		l:    log,
		want: `level=warning type=APPLOG`,
	}, {
		desc: "warn log with msg case",
		args: []interface{}{"warn message"},
		l:    log,
		want: `level=warning msg="warn message" type=APPLOG`,
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			tt.l.Warn(tt.args...)
			output := buffer.String()
			buffer.Reset()
			if strings.TrimRight(output, "\n\r") != tt.want {
				t.Errorf("Warn(%v) failed: got %v, want %v", tt.args, output, tt.want)
			}
		})
	}
}

func TestLoggerWarnf(t *testing.T) {
	var buffer bytes.Buffer
	log := fakeLogger()
	log.SetOutput(&buffer)
	tests := []struct {
		desc   string
		l      Logger
		format string
		args   []interface{}
		want   string
	}{{
		desc: "empty Warnf log case",
		l:    log,
		want: `level=warning type=APPLOG`,
	}, {
		desc:   "warnf log with msg case",
		format: "warnf message %s",
		args:   []interface{}{"testing"},
		l:      log,
		want:   `level=warning msg="warnf message testing" type=APPLOG`,
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			tt.l.Warnf(tt.format, tt.args...)
			output := buffer.String()
			buffer.Reset()
			if strings.TrimRight(output, "\n\r") != tt.want {
				t.Errorf("Warnf(%v) failed: got %v, want %v", tt.args, output, tt.want)
			}
		})
	}
}

func TestLoggerError(t *testing.T) {
	var buffer bytes.Buffer
	log := fakeLogger()
	log.SetOutput(&buffer)
	tests := []struct {
		desc string
		l    Logger
		args []interface{}
		want string
	}{{
		desc: "empty error log case",
		l:    log,
		want: `level=error type=APPLOG`,
	}, {
		desc: "error log with msg case",
		args: []interface{}{"error message"},
		l:    log,
		want: `level=error msg="error message" type=APPLOG`,
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			tt.l.Error(tt.args...)
			output := buffer.String()
			buffer.Reset()
			if strings.TrimRight(output, "\n\r") != tt.want {
				t.Errorf("Error(%v) failed: got %v, want %v", tt.args, output, tt.want)
			}
		})
	}
}

func TestLoggerErrorf(t *testing.T) {
	var buffer bytes.Buffer
	log := fakeLogger()
	log.SetOutput(&buffer)
	tests := []struct {
		desc   string
		l      Logger
		format string
		args   []interface{}
		want   string
	}{{
		desc: "empty errorf log case",
		l:    log,
		want: `level=error type=APPLOG`,
	}, {
		desc:   "errorf log with msg case",
		format: "errorf message %s",
		args:   []interface{}{"testing"},
		l:      log,
		want:   `level=error msg="errorf message testing" type=APPLOG`,
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			tt.l.Errorf(tt.format, tt.args...)
			output := buffer.String()
			buffer.Reset()
			if strings.TrimRight(output, "\n\r") != tt.want {
				t.Errorf("Errorf(%v) failed: got %v, %v", tt.args, output, tt.want)
			}
		})
	}
}

func fakeLogger() Logger {
	l := New(AppLog)
	// remove timestamp from test output
	l.entry.Logger.Formatter.(*logrus.TextFormatter).FullTimestamp = false
	l.entry.Logger.Formatter.(*logrus.TextFormatter).DisableTimestamp = true
	return l
}

func ExampleLogger() {
	test := "testing"
	l := New(AppLog)
	l.SetLevel(DebugLevel)
	// remove timestamp from test output
	l.entry.Logger.Formatter.(*logrus.TextFormatter).FullTimestamp = false
	l.entry.Logger.Formatter.(*logrus.TextFormatter).DisableTimestamp = true

	l.SetOutput(os.Stdout)

	l.Debug("debug message")
	l.WithFields(Fields{
		"field1": "value1",
	}).Debug("debug message with fields")
	l.Debugf("debugf message %s", test)
	l.WithFields(Fields{
		"field1": "value1",
	}).Debugf("debugf message %s", test)
	l.Warn("warn message")
	l.WithFields(Fields{
		"field1": "value1",
	}).Warn("warn message with fields")
	l.Warnf("warnf message %s", test)
	l.WithFields(Fields{
		"field1": "value1",
	}).Warnf("warnf message with fields")
	l.Info("info message")
	l.WithFields(Fields{
		"field1": "value1",
	}).Info("info message with fields")
	l.Infof("infof message %s", test)
	l.WithFields(Fields{
		"field1": "value1",
	}).Infof("infof message with fields %s", test)
	l.WithFields(Fields{
		"jobId": "12345",
		"type":  "job",
		"level": "warn",
	}).Infof("job log with warn log %s", test)

	// Output:
	// level=debug msg="debug message" type=APPLOG
	// level=debug msg="debug message with fields" field1=value1 type=APPLOG
	// level=debug msg="debugf message testing" type=APPLOG
	// level=debug msg="debugf message testing" field1=value1 type=APPLOG
	// level=warning msg="warn message" type=APPLOG
	// level=warning msg="warn message with fields" field1=value1 type=APPLOG
	// level=warning msg="warnf message testing" type=APPLOG
	// level=warning msg="warnf message with fields" field1=value1 type=APPLOG
	// level=info msg="info message" type=APPLOG
	// level=info msg="info message with fields" field1=value1 type=APPLOG
	// level=info msg="infof message testing" type=APPLOG
	// level=info msg="infof message with fields testing" field1=value1 type=APPLOG
	// level=error msg="Field \"type\" cannot be assigned by users.  Discarding" type=APPLOG
	// level=info msg="job log with warn log testing" fields.level=warn jobId=12345 type=APPLOG

}

func ExampleDefault() {
	test := "testing"
	// remove timestamp from test output
	defaultLogger.entry.Logger.Formatter.(*logrus.TextFormatter).DisableTimestamp = true
	SetLevel(DebugLevel)
	SetOutput(os.Stdout)

	Debug("debug message")
	WithFields(Fields{
		"field1": "value1",
	}).Debug("debug message with fields")
	Debugf("debugf message %s", test)
	WithFields(Fields{
		"field1": "value1",
	}).Debugf("debugf message %s", test)
	Warn("warn message")
	WithFields(Fields{
		"field1": "value1",
	}).Warn("warn message with fields")
	Warnf("warnf message %s", test)
	WithFields(Fields{
		"field1": "value1",
	}).Warnf("warnf message with fields")
	Info("info message")
	WithFields(Fields{
		"field1": "value1",
	}).Info("info message with fields")
	Infof("infof message %s", test)
	WithFields(Fields{
		"field1": "value1",
	}).Infof("infof message with fields %s", test)
	WithFields(Fields{
		"jobId": "12345",
		"type":  "job",
		"level": "warn",
	}).Infof("job log with warn log %s", test)
	Error("Error happened")
	Errorf("Error with formatting: %s", "some data")

	// Output:
	// level=debug msg="debug message" type=APPLOG
	// level=debug msg="debug message with fields" field1=value1 type=APPLOG
	// level=debug msg="debugf message testing" type=APPLOG
	// level=debug msg="debugf message testing" field1=value1 type=APPLOG
	// level=warning msg="warn message" type=APPLOG
	// level=warning msg="warn message with fields" field1=value1 type=APPLOG
	// level=warning msg="warnf message testing" type=APPLOG
	// level=warning msg="warnf message with fields" field1=value1 type=APPLOG
	// level=info msg="info message" type=APPLOG
	// level=info msg="info message with fields" field1=value1 type=APPLOG
	// level=info msg="infof message testing" type=APPLOG
	// level=info msg="infof message with fields testing" field1=value1 type=APPLOG
	// level=error msg="Field \"type\" cannot be assigned by users.  Discarding" type=APPLOG
	// level=info msg="job log with warn log testing" fields.level=warn jobId=12345 type=APPLOG
	// level=error msg="Error happened" type=APPLOG
	// level=error msg="Error with formatting: some data" type=APPLOG

}

func TestLevels(t *testing.T) {
	l := New(AppLog)
	for _, level := range AllLevels {
		lvl, err := ParseLevel(level.String())
		if err != nil {
			t.Fatalf("ParseLevel() failed: error %v", err)
		}
		if lvl != level {
			t.Fatalf("Failed to regenerate level from string: %v", level)
		}
		l.SetLevel(level)
		if l.entry.Logger.GetLevel() != logrus.Level(level) {
			t.Fatalf("SetLevel(%q) failed: got %v, want %v", level, l.entry.Logger.GetLevel(), logrus.Level(level))
		}
	}
}
