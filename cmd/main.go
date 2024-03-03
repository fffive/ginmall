package main

import (
	"fmt"
	"ginmall/conf"
	"ginmall/routes"
)

func main() {
	conf.Init()

	r := routes.NewRouter()
	err := r.Run(conf.HttpPort)
	if err != nil {
		fmt.Println("router run with error -- ", err)
	}
}
