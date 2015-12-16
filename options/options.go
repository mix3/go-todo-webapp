package options

import "github.com/kelseyhightower/envconfig"

type Options struct {
	DBUser string `envconfig:"DB_USER" default:"root"`
	DBPass string `envoncifg:"DB_PASS" default:""`
	DBHost string `envconfig:"DB_HOST" default:"127.0.0.1"`
	DBPort int    `envconfig:"DB_PORT" default:"3306"`
	DBName string `envconfig:"DB_NAME" default:"todo"`
	Host   string `envconfig:"HOST" default:"localhost"`
	Port   int    `envconfig:"PORT" default:"5000"`
	Static string `envconfig:"STATIC" default:"./static"`
	Debug  bool   `envconfig:"DEBUG" default:"false"`
}

func Get() (Options, error) {
	var opts Options
	err := envconfig.Process("todo", &opts)
	return opts, err
}
