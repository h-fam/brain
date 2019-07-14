package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/h-fam/errdiff"
)

func TestNew(t *testing.T) {

	tests := []struct {
		desc       string
		configFile string
		configName string
		k          string
		v          interface{}
		wantErr    bool
	}{{
		desc:       "valid string",
		configName: "test",
		k:          "string",
		v:          "that",
	}, {
		desc:       "valid int",
		configName: "test",
		k:          "int",
		v:          1234,
	}, {
		desc:       "valid map",
		configName: "test",
		k:          "map",
		v:          map[string]string{"key": "value"},
	}, {
		desc:       "valid slice",
		configName: "test",
		k:          "array",
		v:          []string{"first", "second", "third"},
	}}

	tmpPath, err := ioutil.TempDir("", "newTest")
	if err != nil {
		t.Fatalf("Failed to create tmpdir: %v", err)
	}
	defer os.RemoveAll(tmpPath)

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			c, err := New(tt.configName, tmpPath)
			if s := errdiff.Check(err, tt.wantErr); s != "" {
				t.Fatalf("New(%v) failed: %s", tt.configName, s)
			}
			if err != nil {
				return
			}
			c.Set(tt.k, tt.v)
			got := c.Get(tt.k)
			if !reflect.DeepEqual(got, tt.v) {
				t.Fatalf("Invalid value %q: got %v, want %v", tt.k, got, tt.v)
			}
			if err := c.Write(); err != nil {
				t.Fatalf("failed to save config: %v", err)
			}
			b, err := ioutil.ReadFile(c.filename)
			if err != nil {
				t.Fatalf("failed to read config: %v", err)
			}
			fmt.Println(string(b))
		})
	}
}
