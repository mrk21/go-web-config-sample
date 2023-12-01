package config

import (
	"embed"

	"github.com/mrk21/go-web-config-sample/confutil"
)

//go:embed *.yaml
var ConfFS embed.FS

type Config struct {
	Hoge string     `yaml:"hoge"`
	Fuga ConfigFuga `yaml:"fuga"`
}
type ConfigFuga struct {
	Bar ConfigFugaBar `yaml:"bar"`
}
type ConfigFugaBar struct {
	Value1 int `yaml:"value1"`
	Value2 int `yaml:"value2"`
}

var globalLoader = confutil.NewGlobalLoader[Config](ConfFS)

func Load() error  { return globalLoader.Load() }
func Get() *Config { return globalLoader.Get() }
