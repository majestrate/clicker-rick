package main

import (
	"github.com/majestrate/clicker-rick/config"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	conf := config.DefaultConfig
	conf.Instance.Email = os.Args[1]
	conf.Instance.Domain = os.Args[2]
	err := conf.Save(os.Args[3])
	if err != nil {
		panic(err)
	}
}
