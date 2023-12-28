package model

type Endpoint struct {
	Base    string `yaml:"base"`
	Backend []struct {
		Pattern string `yaml:"pattern"`
		Method  string `yaml:"method"`
		Target  struct {
			URL       string `yaml:"url"`
			SSL       bool   `yaml:"ssl"`
			Method    string `yaml:"method"`
			Encrypted bool   `yaml:"encrypted"`
		} `yaml:"target"`
		Plugin struct {
			JwtEnabled bool   `yaml:"jwtEnabled"`
			JwtName    string `yaml:"jwtName"`
		} `yaml:"plugin"`
	} `yaml:"backend"`
}
