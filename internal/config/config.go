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
	Redis struct {
		Uri      string `mapstructure:"uri"`
		Password string `mapstructure:"pasword"`
		Database int    `mapstructure:"database"`
	} `mapstructure:"redis"`
	App struct {
		RateLimit struct {
			Requests int   `mapstructure:"requests"`
			Seconds  int64 `mapstructure:"seconds"`
		} `mapstructure:"rateLimit"`
	} `mapstructure:"app"`
}

func NewConfig() *Config {
	viper.SetConfigFile("config/local.yaml")
	viper.AutomaticEnv()

	viper.BindEnv("abios.secret", "ABIOS_SECRET")
	viper.BindEnv("redis.password", "REDIS_PASSWORD")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Failed to read config: ", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal("Failed to parse config values: ", err)
	}

	log.Printf("Config loaded: %+v\n", cfg)

	return &cfg
}
