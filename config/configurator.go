package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/atsuiest/gapigate/model"
	"gopkg.in/yaml.v2"
)

type Configuration struct {
	JwtSecret string `yaml:"jwtSecret"`
	WebClient struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"webClient"`
	Endpoints   []model.Endpoint `yaml:"endpoints"`
	Plugins     []model.Plugin   `yaml:"plugins"`
	Credentials struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"credentials"`
}

var (
	ErrorFileNotFound = "Configuration file not found"
	ErrorYamlFormat   = "Invalid yaml file"
	GlobalConf        = &Configuration{}
	ValidationsMap    = map[string]model.Validation{}
	BackendMap        = map[string]model.Backend{}
)

func init() {
	urlconfig := os.Getenv("CONFIG")
	if urlconfig == "" {
		println("Using PWD os " + urlconfig)
		pwd, _ := os.Getwd()
		urlconfig = pwd + "/config.yaml"
	}
	file, err := os.ReadFile(urlconfig)
	if err != nil {
		println(ErrorFileNotFound + " trying accessing " + urlconfig)
		os.Exit(0)
	}
	err = yaml.Unmarshal(file, &GlobalConf)
	if err != nil {
		println(err.Error())
		os.Exit(0)
	}
	ValidationsMap = map[string]model.Validation{}
	for _, v := range GlobalConf.Plugins {
		for _, w := range v.Validations {
			ValidationsMap[v.Type+"|"+w.Name] = w
		}
	}
	BackendMap = map[string]model.Backend{}
	for _, v := range GlobalConf.Endpoints {
		for _, w := range v.Backend {
			target := strings.Split(w.Pattern, "/:")
			if len(target) > 1 {
				target[0] = target[0] + "/::"
			}
			BackendMap[w.Method+"|"+v.Base+target[0]] = w
		}
	}
	for k, v := range BackendMap {
		fmt.Println(k, "value is", v)
	}
}
