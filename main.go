package main

import (
	"os"

	"github.com/RaymondCode/simple-demo/controller"
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/gin-gonic/gin"
)

const URL string = "http://106.13.196.236"
const PORT string = ":8080"
const GROUPPATH string = "/douyin"
const STATICPATH string = "./public"

const SQLHOST string = "106.13.196.236"
const SQLPORT string = "3306"
const SQLDATABASE string = "douyin"
const SQLUSER string = "root"
const SQLPASSWORD string = "666666"
const SQLCHARSET string = "utf8"

func main() {
	err := Init()
	if err != nil {
		os.Exit(-1)
	}
	r := gin.Default()
	r.MaxMultipartMemory = 100 << 20
	initRouter(r)
	r.Run(PORT) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func Init() error {
	err := dao.InitDB(SQLHOST, SQLPORT, SQLDATABASE, SQLUSER, SQLPASSWORD, SQLCHARSET)
	if err != nil {
		return err
	}
	controller.URL = URL
	controller.PORT = PORT
	controller.GROUPPATH = GROUPPATH
	controller.STATICPATH = STATICPATH
	return nil
}
