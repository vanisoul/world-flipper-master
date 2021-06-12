package main

import "github.com/spf13/viper"

func LoadSmallRoleConfig() (config ExampleConfig, err error) {
	viper.AddConfigPath("./config/")
	viper.SetConfigName("example")
	viper.SetConfigType("json")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

type ExampleConfig struct {
	Name string `mapstructure:"name"`
	Id   int    `mapstructure:"id"`
}
