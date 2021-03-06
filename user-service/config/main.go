package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Conf Config

type Config struct {
	Port       int    `yaml:"port"`
	EncryptKey string `yaml:"encryptKey"`
	DB         struct {
		Host     string `yaml:"host"`
		Port     uint   `yaml:"port"`
		User     string `yaml:"user"`
		Name     string `yaml:"name"`
		Password string `yaml:"password"`
	} `yaml:"db"`
}

func IniConfig() error {
	source, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(source, &Conf)
	if err != nil {
		panic(err)
	}
	return nil
}
