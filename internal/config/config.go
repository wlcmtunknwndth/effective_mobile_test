package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Db     Database `yaml:"db"`
	Server Server   `yaml:"server"`
	Broker Broker   `yaml:"broker"`
}

type Server struct {
	Address     string        `yaml:"address" env-default:"localhost:7777"`
	Timeout     time.Duration `yaml:"timeout" env-default:"5s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"30s"`
}

type Database struct {
	DbUser  string `yaml:"db_user" env-required:"true"`
	DbPass  string `yaml:"db_pass" env-required:"true"`
	DbName  string `yaml:"db_name" env-required:"true"`
	SslMode string `yaml:"ssl_mode" env-default:"false"`
	Port    string `yaml:"port" env-required:"5432"`
}

type Broker struct {
	MaxReconnects int           `yaml:"max_reconnects"`
	ReconnectWait time.Duration `yaml:"reconnect_wait"`
	Address       string        `yaml:"address"`
	Retry         bool          `yaml:"retry"`
}

const configPathEnv = "config_path"

func MustLoad() *Config {
	const op = "internal.config.MustLoad"

	path, ok := os.LookupEnv(configPathEnv)
	if !ok || path == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}

//func fetchConfigPath() string{
//	var res string
//
//	flag.StringVar(&res, "config", "", "path to config file")
//	flag.Parse()
//}
