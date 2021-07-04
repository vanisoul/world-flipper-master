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
fNumber : 從上而下第幾個關卡
fDifficulty : 難度由上而下第幾個
rNumber: 從上而下第幾個關卡
rDifficulty : 難度由上而下第幾個
rFrequency : 滿體後打幾次r關卡
rFeatures : 是否開啟滿體打關卡功能
*/
type CheckpointConfig struct {
	Nox         int    `mapstructure:"nox"`
	Type        string `mapstructure:"type"`
	FNumber     int    `mapstructure:"fNumber"`
	FDifficulty int    `mapstructure:"fDifficulty"`
	RNumber     int    `mapstructure:"rNumber"`
	RDifficulty int    `mapstructure:"rDifficulty"`
	RFrequency  int    `mapstructure:"rFrequency"`
	RFeatures   bool   `mapstructure:"rFeatures"`
	Notthink    int    `mapstructure:"notthink"`
}
