package config

import (
	"golangtest/src/constanta"

	"github.com/spf13/viper"
)

var (
	CfgData = ConfigApps("./resources/")
)

func ConfigApps(path string) *DefaultConfig {
	//viper.SetConfigFile(path)
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	//viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(constanta.SysConfigFailedRead)
	}

	var conf DefaultConfig

	if err := viper.Unmarshal(&conf); err != nil {
		panic(constanta.SysConfigUnmarshall)
	}

	return &conf
}
