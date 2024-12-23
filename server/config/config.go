package config

import "github.com/kelseyhightower/envconfig"

type Configuration struct {
	Port string `default:"8080"`
}

func Load() (*Configuration, error) {
	envconfig.Process()

}
