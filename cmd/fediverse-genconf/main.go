package main

import (
	"github.com/majestrate/clicker-rick/config"
	"os"
)

func main() {
	if len(os.Args) != 5 {
		return
	}

	conf := config.DefaultConfig
	conf.Instance.Email = os.Args[1]
	conf.Instance.Domain = os.Args[2]
	conf.Assets.Root = os.Args[3]
	err := conf.Save(os.Args[4])
	if err != nil {
		panic(err)
	}
}
