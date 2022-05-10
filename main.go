package main

import (
	"os"

	"github.com/RaymondCode/simple-demo/dao"
	"github.com/gin-gonic/gin"
)

func main() {
	err := Init()
	if err != nil {
		os.Exit(-1)
	}
	r := gin.Default()
	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func Init() error {
	err := dao.InitDB()
	if err != nil {
		return err
	}
	return nil
}
