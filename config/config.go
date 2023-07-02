package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DBConfig    *DatabaseConfig
	LogConfig   *LogConfig
	AppConfig   *AppConfig
	RedisConfig *RedisConfig
}

func loadDefaultConfig[T comparable](key string, t T) T {
	val, ok := viper.Get(key).(T)
	if !ok || val == *new(T) {
		return t
	}
	return val
}

// LoadConfig sets up viper and returns config, or panics if config can't be loaded
func LoadConfig() Config {
	switch os.Getenv("ENV") {
	case "local":
		loadLocalConfig()
	default:
		loadStandardConfig()
	}
	return Config{
		DBConfig:    loadDBConfig(),
		LogConfig:   loadLogConfig(),
		AppConfig:   loadAppConfig(),
		RedisConfig: loadRedisConfig(),
	}
}

func loadStandardConfig() {
	viper.ReadInConfig()
}

func loadLocalConfig() {
	viper.SetConfigName("local")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
}
