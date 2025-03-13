package config

import (
	"github.com/CaptainFallaway/realgofile/pkg/helpers"
)

type Config struct {
	Addr     string
	Debug    bool
	DbString string
}

func LoadEnv(conf *Config) {
	conf.Addr = helpers.GetEnvVar("ADDR", conf.Addr)
	conf.Debug = helpers.GetEnvVar("DEBUG", conf.Debug)
	conf.DbString = helpers.GetEnvVar("DBSTRING", conf.DbString)
}
