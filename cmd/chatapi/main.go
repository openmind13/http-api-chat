package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/openmind13/http-api-chat/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/chatapi.toml", "path to config file")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	flag.Parse()

	// scan config
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("start api server on port %v\n", config.BindAddr)

	// start api server
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
