package global

import (
	"gin_template/config"
	"github.com/spf13/viper"
)

var (
	GLA_CONFIG config.Server
	GLA_VIPER  *viper.Viper
)
