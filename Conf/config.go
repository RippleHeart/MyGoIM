package Conf

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	M MySQLConfig `mapstructure:"mysql"`
	R RedisConfig `mapstructure:"redis"`
}
type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"db"`
	MaxIdleConns int    `mapstructure:"MaxIdleConns"`
	MaxOpenConns int    `mapstructure:"MaxOpenConns"`
}
type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"PoolSize"`
}

var Conf Config
var DSN string

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./Conf")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("ReadInConfig error:", err)
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		log.Fatal("Config Unmarshal error:", err)
	}

	DSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", Conf.M.User, Conf.M.Password, Conf.M.Host, Conf.M.Port, Conf.M.DB)
	//fmt.Println("DSN:", DSN)
}
