package servers

import "github.com/spf13/viper"

type Server interface {
	Run(conf *viper.Viper, quit chan bool)
}
