package main

import (
	"github.com/spf13/viper"
)

func LoadBigRoleConfig() (config RoleConfig, err error) {
	viper.AddConfigPath("./config/")
	viper.SetConfigName("big-role")
	viper.SetConfigType("json")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func LoadSmallRoleConfig() (config RoleConfig, err error) {
	viper.AddConfigPath("./config/")
	viper.SetConfigName("small-role")
	viper.SetConfigType("json")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

type RoleConfig struct {
	RoleImgPair []RoleImgPair `mapstructure:"roleImgPair"`
}

type RoleImgPair struct {
	ImgName  string `mapstructure:"imgName"`
	RoleType int    `mapstructure:"roleType"`
	RoleName string `mapstructure:"roleName"`
}
