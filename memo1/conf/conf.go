package conf

import (
	"fmt"
	"strings"

	"gopkg.in/ini.v1"
)

var (
	AppMode    string
	HttpPort   string
	DbUser     string
	DbPassWord string
	DbHost     string
	DbName     string
	DbPort     string
	Path       string
)

func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误")
	}
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbName = file.Section("mysql").Key("DbName").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	fmt.Println(AppMode, HttpPort, DbUser, DbPassWord, DbHost, DbName)
	//"用户名:密码@tcp(ip:port)/dbName?charset=utf8mb4&parseTime=True&loc=Local"
	Path = strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort,
		")/", DbName, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")

}
