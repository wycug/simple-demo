/**
@author: Junming ZHANG, Haojun YIN
@date: 2022/5/27
**/

package initializer

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

func InitDataBase() {
	m := global.Config.Mysql
	var dsn = fmt.Sprintf("%s:%s@%s", m.Username, m.Password, m.Url)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		panic(fmt.Errorf("gorm open error: %s\n", err))
	}

	sqlDb, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("db db error: %s\n", err))
	}
	sqlDb.SetMaxIdleConns(m.MaxIdleConns)
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime) * time.Second)
	global.Db = db
}
