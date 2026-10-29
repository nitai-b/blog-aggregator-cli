package main

import (
	"fmt"
	"log"

	"github.com/nitai-b/blog-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	if err := cfg.SetUser("nitai-b"); err != nil {
		log.Fatal(err)
	}

	newcfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(newcfg.DBURL)
	fmt.Println(newcfg.CurrentUserName)
}
