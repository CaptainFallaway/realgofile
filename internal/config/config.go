package config

import (
	"github.com/CaptainFallaway/realgofile/pkg/common"
)

type Config struct {
	Addr     string
	Debug    bool
	DbString string
}

func LoadEnv(conf *Config) {
	conf.Addr = common.GetEnvVar("ADDR", conf.Addr)
	conf.Debug = common.GetEnvVar("DEBUG", conf.Debug)
	conf.DbString = common.GetEnvVar("DBSTRING", conf.DbString)
}
