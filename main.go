package main

import (
	"fmt"
	"go-tcp-server/config"
	"go-tcp-server/lib/file"
	"go-tcp-server/lib/logger"
	"go-tcp-server/tcp"
)

const configFile string = "etc/config.yaml"

func main() {
	logger.Setup(&logger.Settings{
		Path:       "logs",
		Name:       "tcp-server",
		Ext:        "log",
		TimeFormat: "2006-01-02",
	})

	if !file.CheckNotExist(configFile) {
		config.SetupConfig(configFile, ".")
	}

	err := tcp.ListenAndServeWithSignal(
		&tcp.Config{Address: fmt.Sprintf("%s:%d", config.Properties.Bind, config.Properties.Port)},
		tcp.MakeHandler())

	if err != nil {
		logger.Fatal(err)
	}
}
