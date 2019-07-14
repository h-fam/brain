package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Config holds configuration information about the user's settings.
type Config struct {
	*viper.Viper
}

// New returns a new configuration named name.
func New(name string) (*Config, error) {
	v := viper.New()
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get homedir: %v", err)
	}
	v.SetConfigName(name)
	v.AddConfigPath(homeDir)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	c := Config{v}
	return &c, nil
}
