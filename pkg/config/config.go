package config

import "github.com/spf13/viper"

func New() *viper.Viper {
	conf := viper.New()
	conf.SetDefault("SERVER_PORT", "5300")
	conf.SetDefault("REDIS_PORT", "6379")
	return conf
}
