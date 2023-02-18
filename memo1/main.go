package main

import (
	"jiuxia/memo1/conf"
	"jiuxia/memo1/model"
	"jiuxia/memo1/routes"
)

func main() {
	conf.Init()      //加载配置文件
	model.Database() //连接数据库
	r := routes.NewRouter()
	r.Run(conf.HttpPort)

}
