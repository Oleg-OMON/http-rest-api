package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Oleg-OMON/http-rest-api.git/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/apiserver.json", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	configFile, err := os.Open(configPath)
	if err != nil {
		fmt.Println(err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}
