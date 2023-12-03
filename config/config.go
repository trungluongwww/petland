package config

import (
	"github.com/caarlos0/env/v10"
	"log"
)

type Env struct {
	// common
	AppPort string `env:"PORT"`
	AppEnv  string `env:"ENVIRONMENT"`

	// config admin
	SecretJwtAdmin string `env:"SECRET_JWT_ADMIN,required"`

	// config app
	SecretJwtApp string `env:"SECRET_JWT_APP,required"`

	// database
	MySqlUsername string `env:"MYSQL_USERNAME,required"`
	MySqlPassword string `env:"MYSQL_PASSWORD,required"`
	MYSqlDatabase string `env:"MYSQL_DATABASE,required"`
	MySqlProtocol string `env:"MYSQL_PROTOCOL,required"`
}

var (
	cfgEnv Env
)

func Init() Env {
	if err := env.Parse(&cfgEnv); err != nil {
		log.Panicf("Init env failed %v", err)
	}

	return cfgEnv
}

func (e Env) IsDevelop() bool {
	return e.AppEnv == "develop"
}

func GetEnv() Env {
	return cfgEnv
}
