package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

var AppConfig struct {
	Gmodel struct {
		Verbose    bool
		Connection map[any]struct {
			Dsn string
		}
	}
}

func Parse(configpath string) (err error) {
	var bs []byte
	bs, err = os.ReadFile(configpath)
	if err == nil {
		err = yaml.Unmarshal(bs, &AppConfig)
	}
	return
}
