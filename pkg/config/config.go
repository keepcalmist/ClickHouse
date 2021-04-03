package config

import "github.com/spf13/viper"

func New() *viper.Viper {
	conf := viper.New()
	conf.SetDefault("GRPC_SERVER_PORT", "5300")
	conf.SetDefault("REST_SERVER_ADDRESS", "5400")
	conf.SetDefault("REDIS_PORT", "6379")
	return conf
}
