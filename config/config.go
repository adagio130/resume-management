package config

import "time"

type Config struct {
	DB     DB     `mapstructure:"db" yaml:"db"`
	Server Server `mapstructure:"server" yaml:"server"`
	Log    Log    `mapstructure:"log" yaml:"log"`
}

type DB struct {
	Dsn             string        `mapstructure:"dsn" yaml:"dsn" default:"file:test.db?cache=shared&mode=memory"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns" yaml:"max_idle_conns" default:"10"`
	MaxOpenConns    int           `mapstructure:"max_open_conns" yaml:"max_open_conns" default:"10"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime" yaml:"conn_max_lifetime" default:"10"`
}

type Server struct {
	Port string `mapstructure:"port" yaml:"port" default:"8080"`
}

type Log struct {
	Env   string `mapstructure:"env" yaml:"env" default:"development"`
	Level string `mapstructure:"level" yaml:"level" default:"info"`
}
