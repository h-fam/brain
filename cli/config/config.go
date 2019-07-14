package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Config holds configuration information about the user's settings.
type Config struct {
	*viper.Viper
	filename string
}

func getHome() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get homedir: %v", err)
	}
	return homeDir, nil
}

// New returns a new configuration named name.  Path must
// be set to a valid user file path.
func New(name, path string) (*Config, error) {
	v := viper.New()

	v.SetConfigName(name)
	v.SetConfigType("yaml")
	v.AddConfigPath(path)
	if err := v.ReadInConfig(); err != nil {
		// if the error is file not found that is fine.
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}
	c := Config{
		Viper:    v,
		filename: filepath.Join(path, name+".yaml"),
	}
	return &c, nil
}

func (c *Config) Write() error {
	return c.WriteConfigAs(c.filename)
}
