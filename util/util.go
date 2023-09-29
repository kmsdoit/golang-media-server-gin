package util

import (
	"go-media-server/api/camera/schema"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path/filepath"
)

type DatabaseConfig struct {
	Type     string `yaml:"Type"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	Host     string `yaml:"Host"`
	Name     string `yaml:"Name"`
}

type Config struct {
	Database DatabaseConfig `yaml:"database"`
}

func Util_Init() {
	filename, _ := filepath.Abs("config.yaml")
	yamlFile, err := ioutil.ReadFile(filename)
	var config Config

	err = yaml.Unmarshal(yamlFile, &config)

	if err != nil {
		panic(err)
	}

	db := DB_INIT(config.Database.User, config.Database.Password, config.Database.Host, config.Database.Name)
	db.AutoMigrate(&schema.CameraSchema{})
}
