package main

import "github.com/spf13/viper"

func LoadCheckpointConfig() (config CheckpointConfig, err error) {
	viper.AddConfigPath("./config/")
	viper.SetConfigName("Checkpoint")
	viper.SetConfigType("json")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

/*
Type : boss , activity
number : 從上而下第幾個關卡
Difficulty : 難度由上而下第幾個
*/
type CheckpointConfig struct {
	Type       string `mapstructure:"type"`
	Number     int    `mapstructure:"number"`
	Difficulty int    `mapstructure:"difficulty"`
}
