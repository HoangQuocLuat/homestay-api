package config

import "time"

type Config struct {
	Http       HttpServer       `json:"http" yaml:"http"`
	PostgreSQL PostgreSQLConfig `json:"postgres" yaml:"postgres"`
}

type HttpServer struct {
	Path string `json:"path" yaml:"path"`
	Port string `json:"port" yaml:"port"`
}

type PostgreSQLConfig struct {
	Host            string        `mapstructure:"host" `
	Database        string        `mapstructure:"database"`
	Port            int           `mapstructure:"port"`
	Username        string        `mapstructure:"username"`
	Password        string        `mapstructure:"password" json:"-"`
	Options         string        `mapstructure:"options"`
	MaxIdleConns    int           `mapstructure:"maxIdleConns"`
	MaxOpenConns    int           `mapstructure:"maxOpenConns"`
	ConnMaxLifetime time.Duration `mapstructure:"connMaxLifetime"`
}
