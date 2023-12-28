package model

type Plugin struct {
	Type        string       `yaml:"type"`
	Validations []Validation `yaml:"validations"`
}

type Validation struct {
	Name   string  `yaml:"name"`
	Claims []Claim `yaml:"claims"`
	Value  string  `yaml:"value"`
	Secret string  `yaml:"secret"`
}

type Claim struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}
