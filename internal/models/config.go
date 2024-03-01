package models

type Config struct {
	Port         int      `yaml:"port"`
	Debug        bool     `yaml:"debug"`
	DebugReqBody bool     `yaml:"debugReqBody"`
	Destinations []string `yaml:"destinations"`
}
