package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"

	"go-tcp-server/lib/file"
	"go-tcp-server/lib/logger"
)

type ServerProperties struct {
	Bind           string `yaml:"bind"`
	Port           int    `yaml:"port"`
	AppendOnly     bool   `yaml:"appendOnly"`
	AppendFilename string `yaml:"appendFilename"`
	MaxClients     int    `yaml:"maxClients"`
	RequirePass    string `yaml:"requirePass"`
	Databases      int    `yaml:"databases"`

	Peers []string `yaml:"peers"`
	Self  string   `yaml:"self"`
}

// Properties holds global config properties
var Properties *ServerProperties

func init() {
	// default config
	Properties = &ServerProperties{
		Bind:       "127.0.0.1",
		Port:       6379,
		AppendOnly: false,
	}
}

// SetupConfig reads config file and stores properties into Properties
func SetupConfig(filename, dir string) {
	f, err := file.MustOpen(filename, dir)
	if err != nil {
		logger.Fatal(err)
	}
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		logger.Fatal(err)
	}

	err = yaml.Unmarshal(bs, Properties)
	if err != nil {
		logger.Fatal(err)
	}
}
