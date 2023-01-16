package main

import (
	"paa/memo/conf"
	"paa/memo/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
