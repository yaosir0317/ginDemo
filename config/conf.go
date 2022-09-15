package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	PoolSize int    `yaml:"pool_size"`
	DB       int    `yaml:"db"`
}

type MysqlConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Addr     string `yaml:"addr"`
	DBName   string `yaml:"db_name"`
	MaxOpen  int    `yaml:"max_open"`
	MaxIdle  int    `yaml:"max_idle"`
}

type HTTPConfig struct {
	Host string `yaml:"host"`
	Port int64  `yaml:"port"`
}

type KafkaConfig struct {
	Brokers []string `yaml:"brokers"`
	Topics  []string `yaml:"topics"`
	Version string   `yaml:"version"`
	Group   string   `yaml:"group"`
}

type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxsize"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
}

type Config struct {
	ENV         string       `yaml:"env"`
	ServiceName string       `yaml:"service_name"`
	HTTP        *HTTPConfig  `yaml:"http"`
	Mysql       *MysqlConfig `yaml:"mysql"`
	Redis       *RedisConfig `yaml:"redis"`
	Kafka       *KafkaConfig `yaml:"kafka"`
	JokeMysql   *MysqlConfig `yaml:"joke_mysql"`
	Log         *LogConfig   `yaml:"log"`
}

func NewConfig(filePath string) *Config {
	bs, err := ioutil.ReadFile(filePath)
	if err != nil {
		//log.PanicWithContext(context.Background(), "read config file error", "err", err.Error(), "path", filePath)
	}
	conf := Config{}
	err = yaml.Unmarshal(bs, &conf)
	if err != nil {
		//log.PanicWithContext(context.Background(), "unmarshal config file error", "err", err.Error(), "path", filePath)
	}
	return &conf
}
