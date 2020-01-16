package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/go-yaml/yaml"
)

var Config EnvironmentConfig

type ServerConfig struct {
	Host         string        `yaml:"host"`
	Port         string        `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	IdleTimeout  time.Duration `yaml:"idle_timeout"`
	SecureConn   bool          `yaml:"secure_conn"`
	SecureCert   string        `yaml:"secure_cert"`
	SecureKey    string        `yaml:"secure_key"`
	SecureCA     string        `yaml:"secure_ca"`
	HashCheck    bool          `yaml:"hash_check"`
	HashSalt     string        `yaml:"hash_salt"`
}

type StorageConfig struct {
	Path          string `yaml:"path"`
	ThumbnailSize string `yaml:"thumbnail"`
}

type DbConfig struct {
	Database string `yaml:"db_name"`
	Bucket   string `yaml:"bucket"`
}

type EnvironmentConfig struct {
	Server  ServerConfig  `yaml:"server"`
	Storage StorageConfig `yaml:"storage"`
	Db      DbConfig      `yaml:"db"`
}

func LoadConfig(name string) error {
	cwd, _ := os.Getwd()

	data, err := ioutil.ReadFile(filepath.Join(cwd, name))
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		return err
	}

	return nil
}
