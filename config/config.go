package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Config represents the application configuration
type Config struct {
	Server   ServerConfig `yaml:"server"`
	Database Database     `yaml:"Database"`
	JWT      JWTConfig    `yaml:"jwt"`
}

// ServerConfig represents the server configuration
type ServerConfig struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	Protocol string `yaml:"Protocol"`
	CertFile string `yaml:"CertFile"`
	KeyFile  string `yaml:"KeyFile"`
}

type Database struct {
	Danymo Danymo `yaml:"danymo"`
	Maria  Maria  `yaml:"maria"`
}

type Danymo struct {
	Prifix             string `yaml:"prifix"`
	TransactionHistory string `yaml:"transactionHistory"`
}

type Maria struct {
	Prifix          string `yaml:"prifix"`
	PersonalWallets string `yaml:"personal_wallets"`
	PublicWallets   string `yaml:"public_wallets"`
	StoreWallets    string `yaml:"store_wallets"`
	MaxIdleConns    int    `yaml:"MaxIdleConns"`
	MaxOpenConns    int    `yaml:"MaxOpenConns"`
	ConnMaxLifetime uint64 `yaml:"ConnMaxLifetime"`
}

// JWTConfig represents the JWT configuration
type JWTConfig struct {
	Secret     string `yaml:"secret"`
	ExpireTime uint64 `yaml:"ExpireTime"`
}

// LoadConfig loads the configuration from a YAML file
func LoadConfig(filename string) (*Config, error) {
	// Read the YAML file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Unmarshal the YAML data into Config struct
	Cfg := &Config{}
	err = yaml.Unmarshal(data, Cfg)
	if err != nil {
		return nil, err
	}

	return Cfg, nil
}
