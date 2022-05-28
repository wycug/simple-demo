package config

// import (
// 	"log"

// 	"gopkg.in/ini.v1"
// )

// var (
// 	Dbtype     string //数据库类型
// 	DbHost     string //数据库服务器主机
// 	DbPort     string //数据服务器端口
// 	DbUser     string //数据库用户
// 	DbPassWord string //数据库密码
// 	DbName     string //数据库名

// )

// func init() {
// 	f, err := ini.Load("./config/config.ini")
// 	if err != nil {
// 		log.Fatal("初始化失败")
// 	}
// 	loadDb(f)
// }

// // loadDb 加载数据库相关配置
// func loadDb(file *ini.File) {
// 	s := file.Section("database")
// 	Dbtype = s.Key("Dbtype").MustString("mysql")
// 	DbName = s.Key("DbName").MustString("douyin")
// 	DbPort = s.Key("DbPort").MustString("3306")
// 	DbHost = s.Key("DbHost").MustString("106.13.196.236")
// 	DbUser = s.Key("DbUser").MustString("root")
// 	DbPassWord = s.Key("DbPassWord").MustString("666666")
// }
