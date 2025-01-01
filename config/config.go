package config

type Config struct {
	DB     DB     `mapstructure:"db" yaml:"db"`
	Server Server `mapstructure:"server" yaml:"server"`
	Log    Log    `mapstructure:"log" yaml:"log"`
}

type DB struct {
	Dsn string `mapstructure:"dsn" yaml:"dsn" default:"file:test.db?cache=shared&mode=memory"`
}

type Server struct {
	Port string `mapstructure:"port" yaml:"port" default:"8080"`
}

type Log struct {
	Level string `mapstructure:"level" yaml:"level" default:"info"`
}
