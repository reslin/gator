package main

import (
	"fmt"
	"log"

	"github.com/reslin/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal("error reading config: %v", err)
	}

	cfg.SetUser("bla")

	cfg, err = config.Read()
	if err != nil {
		log.Fatal("error reading config: %v", err)
	}

	fmt.Println(cfg)
}
