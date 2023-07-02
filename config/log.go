package config

import "github.com/spf13/viper"

type LogConfig struct {
	Level string
}

func loadLogConfig() *LogConfig {
	return &LogConfig{
		Level: viper.GetString("log.level"),
	}
}
