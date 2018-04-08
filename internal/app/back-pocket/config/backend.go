package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/bobrofon/back-pocket/internal/app/back-pocket/constant"
)

func initDefaults() {
	viper.SetDefault(constant.HTTPProxy, constant.DefaultHTTPProxy)
	viper.SetDefault(constant.BindAddress, constant.DefaultBindAddress)
}

func initENV() {
	viper.BindEnv(constant.HTTPProxy)
	viper.BindEnv(constant.BindAddress)
}

func initFlags() {
	pflag.String(constant.HTTPProxy, constant.DefaultHTTPProxy,
		"http proxy for storing data")
	pflag.String(constant.BindAddress, constant.DefaultBindAddress,
		"listen address for communication with proxy")
}

func init() {
	initDefaults()
	initENV()
	initFlags()
}
