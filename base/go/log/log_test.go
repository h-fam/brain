package log

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLoggerLog(t *testing.T) {
	tests := []struct {
		desc        string
		fields      Fields
		args        []interface{}
		wantOut     io.Writer
		wantLogType string
	}{{
		desc:        "empty default log case",
		fields:      Fields{},
		wantOut:     os.Stderr,
		wantLogType: "APPLOG",
	}, {
		desc:        "wrong level name log case",
		fields:      Fields{"level": "abc"},
		args:        []interface{}{""},
		wantOut:     os.Stderr,
		wantLogType: "APPLOG",
	}, {
		desc:        "wrong level name value log case",
		fields:      Fields{"level": 10},
		args:        []interface{}{""},
		wantOut:     os.Stderr,
		wantLogType: "APPLOG",
	}, {
		desc:        "wrong log type name value log case",
		fields:      Fields{"type": 10},
		args:        []interface{}{""},
		wantOut:     os.Stderr,
		wantLogType: "APPLOG",
	}, {
		desc:        "with log level case",
		fields:      Fields{"level": "fatal"},
		args:        []interface{}{"testing logs"},
		wantOut:     os.Stderr,
		wantLogType: "APPLOG",
	}, {
		desc: "job log case",
		fields: Fields{
			"type":  "job",
			"level": "trace",
			"jobId": "1234",
		},
		args:        []interface{}{"got job id"},
		wantOut:     os.Stderr,
		wantLogType: "JOB",
	}, {
		desc: "workflow log case",
		fields: Fields{
			"type":  "workflow",
			"level": "warn",
		},
		args:        []interface{}{"workflow log message"},
		wantOut:     os.Stderr,
		wantLogType: "WORKFLOW",
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			lgr := logObj().WithFields(tt.fields)
			loggerObj := lgr.(logger)
			lgr.Log(tt.args...)
			if loggerObj.entry.Logger.Out != tt.wantOut {
				t.Fatalf("Log(%v) failed: got %v, %v", tt.args, loggerObj.entry.Logger.Out, tt.wantOut)
			}
			if loggerObj.entry.Data["type"].(string) != tt.wantLogType {
				t.Fatalf("Log(%v) failed: got %v, %v", tt.args, loggerObj.entry.Logger.Out, tt.wantOut)
			}
		})
	}
}

func TestLoggerLogf(t *testing.T) {
	tests := []struct {
		desc        string
		fields      Fields
		format      string
		args        []interface{}
		wantOut     io.Writer
		wantLogType string
	}{
		{
			desc:        "empty default log case",
			fields:      Fields{},
			wantOut:     os.Stderr,
			wantLogType: "APPLOG",
		},
		{
			desc:        "with log level case",
			fields:      Fields{"level": "fatal"},
			format:      "%s",
			args:        []interface{}{"testing logs"},
			wantOut:     os.Stderr,
			wantLogType: "APPLOG",
		},
		{
			desc: "job log case",
			fields: Fields{
				"type":  "job",
				"level": "warn",
				"jobId": "1234",
			},
			format:      "%s",
			args:        []interface{}{"got job id"},
			wantOut:     os.Stderr,
			wantLogType: "JOB",
		},
		{
			desc: "workflow log case",
			fields: Fields{
				"type":  "workflow",
				"level": "warn",
			},
			format:      "%s",
			args:        []interface{}{"workflow log message"},
			wantOut:     os.Stderr,
			wantLogType: "WORKFLOW",
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			lgr := logObj().WithFields(tt.fields)
			loggerObj := lgr.(logger)
			lgr.Logf(tt.format, tt.args...)
			if loggerObj.entry.Logger.Out != tt.wantOut {
				t.Fatalf("Logf(%v) failed: got %v, %v", tt.args, loggerObj.entry.Logger.Out, tt.wantOut)
			}
			if loggerObj.entry.Data["type"].(string) != tt.wantLogType {
				t.Fatalf("Logf(%v) failed: got %v, %v", tt.args, loggerObj.entry.Logger.Out, tt.wantOut)
			}
		})
	}
}

func TestLoggerDebug(t *testing.T) {
	var buffer bytes.Buffer
	log := logObj()
	log.SetOutput(&buffer)
	tests := []struct {
		desc string
		l    logger
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
				t.Errorf("Debug(%v) failed: got %v, %v", tt.args, output, tt.want)
			}
		})
	}
}

func TestLoggerDebugf(t *testing.T) {
	var buffer bytes.Buffer
	log := logObj()
	log.SetOutput(&buffer)
	tests := []struct {
		desc   string
		l      logger
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
				t.Errorf("Debugf(%v) failed: got %v, %v", tt.args, output, tt.want)
			}
		})
	}
}

func TestLoggerInfo(t *testing.T) {
	var buffer bytes.Buffer
	log := logObj()
	log.SetOutput(&buffer)
	tests := []struct {
		desc string
		l    logger
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
	log := logObj()
	log.SetOutput(&buffer)
	tests := []struct {
		desc   string
		l      logger
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
	log := logObj()
	log.SetOutput(&buffer)
	tests := []struct {
		desc string
		l    logger
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
				t.Errorf("Warn(%v) failed: got %v, %v", tt.args, output, tt.want)
			}
		})
	}
}

func TestLoggerWarnf(t *testing.T) {
	var buffer bytes.Buffer
	log := logObj()
	log.SetOutput(&buffer)
	tests := []struct {
		desc   string
		l      logger
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
				t.Errorf("Warnf(%v) failed: got %v, %v", tt.args, output, tt.want)
			}
		})
	}
}

func TestLoggerError(t *testing.T) {
	var buffer bytes.Buffer
	log := logObj()
	log.SetOutput(&buffer)
	tests := []struct {
		desc string
		l    logger
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
				t.Errorf("Error(%v) failed: got %v, %v", tt.args, output, tt.want)
			}
		})
	}
}

func TestLoggerErrorf(t *testing.T) {
	var buffer bytes.Buffer
	log := logObj()
	log.SetOutput(&buffer)
	tests := []struct {
		desc   string
		l      logger
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

func logObj() logger {
	l := New().(logger)
	// remove timestamp from test output
	l.entry.Logger.Formatter.(*logrus.TextFormatter).FullTimestamp = false
	l.entry.Logger.Formatter.(*logrus.TextFormatter).DisableTimestamp = true
	return l
}

func ExampleLogger() {
	test := "testing"
	l := New()
	// remove timestamp from test output
	l.(logger).entry.Logger.Formatter.(*logrus.TextFormatter).FullTimestamp = false
	l.(logger).entry.Logger.Formatter.(*logrus.TextFormatter).DisableTimestamp = true

	l.(logger).SetOutput(os.Stdout)

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
		"type": "workflow",
	}).Log("workflow info log")
	l.WithFields(Fields{
		"type":  "workflow",
		"level": "debug",
	}).Log("workflow debug log")
	l.WithFields(Fields{
		"jobId": "12345",
		"type":  "job",
		"level": "warn",
	}).Logf("job log with warn log %s", test)

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
	// level=info msg="workflow info log" type=WORKFLOW
	// level=debug msg="workflow debug log" fields.level=debug type=WORKFLOW
	// level=warning msg="job log with warn log testing" fields.level=warn jobId=12345 type=JOB

}
