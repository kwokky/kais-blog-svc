package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

type Config struct {
	Debug    bool `toml:"debug"`
	Port     int  `toml:"port"`
	Database `toml:"database"`
	Default  `toml:"default"`
}

type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	DB       string
}

type Default struct {
	PageSize int64
}

var (
	configFile = ""
	config     Config
	loaded     = false
)

func Get() Config {
	if !loaded {
		filepath := getPath()
		if _, err := toml.DecodeFile(filepath, &config); err != nil {
			log.Fatal("配置文件读取失败！", err)
		}
		loaded = true
	}

	return config
}

func SetPath(path string) {
	configFile = path
}

func getPath() string {
	if configFile != "" {
		return configFile
	}

	// 根据环境变量获取配置文件
	if path := os.Getenv("KAIS_CONFIG"); path != "" {
		return path
	}

	// 默认配置文件
	filepath := "config.toml"

	return filepath
}
