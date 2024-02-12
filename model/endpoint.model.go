package model

type Endpoint struct {
	Base    string    `yaml:"base"`
	Backend []Backend `yaml:"backend"`
}

type Backend struct {
	Pattern string `yaml:"pattern"`
	Method  string `yaml:"method"`
	Target  struct {
		URL       string `yaml:"url"`
		SSL       bool   `yaml:"ssl"`
		Encrypted bool   `yaml:"encrypted"`
	} `yaml:"target"`
	Plugin struct {
		JwtEnabled bool   `yaml:"jwtEnabled"`
		JwtName    string `yaml:"jwtName"`
	} `yaml:"plugin"`
}
