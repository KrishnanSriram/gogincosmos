package config

import "os"

type EnvConfig struct{}

func (ec *EnvConfig) GetParam(param string) string {
	return os.Getenv(param)
}

func (ec *EnvConfig) SetParam(param, value string) {
	os.Setenv(param, value)
}

func NewEnvConfig() *EnvConfig {
	return &EnvConfig{}
}
