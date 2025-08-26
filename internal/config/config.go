package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	v *viper.Viper
}

func New() *Config {
	return &Config{
		v: viper.GetViper(),
	}
}

func (c *Config) Load(cfgFile string) error {
	if cfgFile != "" {
		c.v.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("failed to get home directory: %w", err)
		}

		c.v.AddConfigPath(home)
		c.v.AddConfigPath(".")
		c.v.SetConfigName(".meibel")
		c.v.SetConfigType("yaml")
	}

	c.v.SetEnvPrefix("MEIBEL")
	c.v.AutomaticEnv()

	c.setDefaults()

	if err := c.v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("failed to read config file: %w", err)
		}
	}

	return nil
}

func (c *Config) setDefaults() {
	c.v.SetDefault("server", "http://api.meibel.ai")
	c.v.SetDefault("output", "json")
	c.v.SetDefault("profile", "default")
}

func (c *Config) Get(key string) interface{} {
	profile := c.v.GetString("profile")
	if profile != "" && profile != "default" {
		profileKey := fmt.Sprintf("profiles.%s.%s", profile, key)
		if c.v.IsSet(profileKey) {
			return c.v.Get(profileKey)
		}
	}
	return c.v.Get(key)
}

func (c *Config) GetString(key string) string {
	val := c.Get(key)
	if s, ok := val.(string); ok {
		return s
	}
	return ""
}

func (c *Config) Set(key string, value interface{}) {
	profile := c.v.GetString("profile")
	if profile != "" && profile != "default" {
		key = fmt.Sprintf("profiles.%s.%s", profile, key)
	}
	c.v.Set(key, value)
}

func (c *Config) Save() error {
	configFile := c.v.ConfigFileUsed()
	if configFile == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("failed to get home directory: %w", err)
		}
		configFile = filepath.Join(home, ".meibel.yaml")
	}

	if err := c.v.WriteConfigAs(configFile); err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

	return nil
}

func (c *Config) ListProfiles() []string {
	profiles := []string{"default"}

	profilesMap := c.v.GetStringMap("profiles")
	for name := range profilesMap {
		profiles = append(profiles, name)
	}

	return profiles
}
