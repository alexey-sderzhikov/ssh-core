package main

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/go-playground/validator"
)

type config struct {
	dbAddr string `env:"DB_ADDRESS"`
	dbUser string `env:"DB_USERNAME"`
	dbPass string `env:"DB_PASSWORD"`

	httpAddr string `env:"HTTP_LISTEN"`
}

func createConfig() (config, error) {
	var c config

	err := env.Parse(&c)
	if err != nil {
		return config{}, err
	}

	err = c.validate()
	if err != nil {
		return config{}, err
	}

	c.dump()

	return c, nil
}

func (c *config) validate() error {
	v := validator.New()
	return v.Struct(c)
}

func (c *config) dump() {
	fmt.Println("config is ", c)
}
