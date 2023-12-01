package confutil

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"

	"dario.cat/mergo"
	"gopkg.in/yaml.v3"
)

type Loader[C any] struct {
	fs fs.FS
}

func NewLoader[C any](fs_ fs.FS) *Loader[C] {
	return &Loader[C]{fs: fs_}
}

func (l *Loader[C]) Load(env Env) (*C, error) {
	conf := map[string]interface{}{}
	c, err := l.loadYAML("config.yaml")
	if err != nil {
		return nil, err
	}
	mergo.Merge(&conf, &c, mergo.WithOverride)

	paths := []string{
		"config.local.yaml",
		fmt.Sprintf("config.%s.yaml", env),
		fmt.Sprintf("config.%s.local.yaml", env),
	}
	for _, p := range paths {
		c, err := l.loadYAML(p)
		if err != nil && !os.IsNotExist(err) {
			return nil, err
		}
		mergo.Merge(&conf, &c, mergo.WithOverride)
	}

	result := new(C)
	tmp, err := yaml.Marshal(conf)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(tmp, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (l *Loader[C]) loadYAML(path string) (map[string]interface{}, error) {
	f, err := l.fs.Open(path)
	if err != nil {
		return nil, err
	}
	w := bytes.NewBuffer([]byte{})
	name := "main"
	t := NewYAMLTemplate()
	err = t.Compile(name, f, w)
	if err != nil {
		return nil, err
	}

	c := map[string]interface{}{}
	err = yaml.Unmarshal(w.Bytes(), &c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
