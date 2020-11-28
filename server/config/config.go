package config

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Database   *DatabaseConfig `yaml:"database"`
	HttpConfig *HttpConfig     `yaml:"http"`
	Mask       bool            `yaml:"mask"`
}

type DatabaseConfig struct {
	Path   string        `yaml:"path"`
	Tables *TablesConfig `yaml:"tables"`
}

type TablesConfig struct {
	QQ bool `yaml:"qq"`
	JD bool `yaml:"jd"`
	SF bool `yaml:"sf"`
}

type HttpConfig struct {
	Host string     `yaml:"host"`
	Port uint16     `yaml:"port"`
	TLS  *TLSConfig `yaml:"tls"`
}

type TLSConfig struct {
	CertPath string `yaml:"cert_path"`
	KeyPath  string `yaml:"key_path"`
}

func NewConfig(configPath string) (*Config, error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to read config file: %s", err))
	}
	config := new(Config)
	if err = yaml.Unmarshal(data, config); err != nil {
		return nil, errors.New(fmt.Sprintf("failed to parse config file: %s", err))
	}
	return config, nil
}
