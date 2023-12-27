package config

import (
	"os"

	"github.com/atsuiest/gapigate/model"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Configuration struct {
	WebClient struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"webClient"`
	Endpoints   []model.Endpoint `yaml:"endpoints"`
	Plugins     []model.Plugin   `yaml:"plugin"`
	Credentials struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"credentials"`
	Secrets []model.Secret `yaml:"secrets"`
}

var (
	ErrorFileNotFound = "Configuration file not found"
	ErrorYamlFormat   = "Invalid yaml file"
	GlobalConf        = &Configuration{}
)

func init() {
	urlconfig := os.Getenv("CONFIG")
	if urlconfig == "" {
		log.Info("Using PWD os " + urlconfig)
		pwd, _ := os.Getwd()
		urlconfig = pwd + "/config.yaml"
	}
	file, err := os.ReadFile(urlconfig)
	if err != nil {
		println(ErrorFileNotFound + " intentando acceder a la ruta: " + urlconfig)
		os.Exit(0)
	}
	err = yaml.Unmarshal(file, &GlobalConf)
	if err != nil {
		log.Error(err.Error())
		os.Exit(0)
	}
	// if yaml.Unmarshal(file, &GlobalConf) != nil {
	// 	println(ErrorYamlFormat)
	// 	os.Exit(0)
	// }
}
