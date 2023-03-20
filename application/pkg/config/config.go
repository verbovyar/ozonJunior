package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("config")

	err = viper.ReadInConfig()
	if err != nil {
		_ = fmt.Errorf("do not parse config file:%v", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		_ = fmt.Errorf("do not unmarshal config file:%v", err)
	}

	return
}
