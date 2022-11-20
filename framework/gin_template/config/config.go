package config

import _ "github.com/mitchellh/mapstructure"

type Server struct {
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
}
