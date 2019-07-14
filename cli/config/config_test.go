package config

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		desc string
		configFile string
		configName string
		k string
		v interface{}
	}{{
		desc: "invalid file",
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T){
			c, err := New(tt.configName)
			if s := errdiff.Check(err, tt.wantErr); s != ""{
				t.Fatalf("New(%v) failed: %s", tt.configName, s)
			}
			got := c.Get(tt.k)
			if !reflect.DeepEqual(got, tt.v) {
				t.Fatalf("Invalid value %q: got %v, want %v", tt.k, got, tt.want)
			}
		})
	}
}
