package configs

import (
	"fmt"
	"sync"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type (
	config struct {
		TSDB TSDB `mapstructure:"tsdb"`
	}

	TSDB struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	}
)

var (
	configs    *config
	configOnce sync.Once
)

// NewConfigs 单列入口
func NewConfigs() *config {
	configOnce.Do(func() {
		configs = &config{}
	})
	return configs
}

// Init 初始化项目
func (c *config) Read(cmd *cobra.Command, cfgFile string) error {
	v := viper.New()
	_ = v.BindPFlags(cmd.Flags())
	// Find home directory.
	v.AddConfigPath(cfgFile)
	v.SetConfigName("config")
	v.SetConfigType("toml")

	// Find and read the configs file
	if err := v.ReadInConfig(); err != nil { // Handle errors reading the configs file
		return err
	}

	if err := v.Unmarshal(c); err != nil {
		return err
	}
	fmt.Println(c)
	return nil
}
