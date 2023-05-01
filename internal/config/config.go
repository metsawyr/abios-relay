package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Abios struct {
		ApiUri string `mapstructure:"apiUri"`
		Secret string `mapstructure:"secret"`
	} `mapstructure:"abios"`
}

func NewConfig() Config {
	viper.SetConfigFile("config/default.yaml")
	viper.AutomaticEnv()

	viper.BindEnv("abios.secret", "ABIOS_SECRET")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Failed to read config: ", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal("Failed to parse config values: ", err)
	}

	log.Printf("Config loaded: %+v\n", cfg)

	return cfg
}
