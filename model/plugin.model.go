package model

type Plugin struct {
	Type        string       `yaml:"type"`
	Validations []Validation `yaml:"validations"`
}

type Validation struct {
	Name   string `yaml:"name"`
	Claim  string `yaml:"claim"`
	Value  string `yaml:"value"`
	Secret string `yaml:"secret"`
}
