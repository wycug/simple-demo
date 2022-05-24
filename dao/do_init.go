package dao

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// gorm.Model definition
// type Model struct {
// 	ID        uint `gorm:"primaryKey"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
// }

func InitDB() error {

	host := "127.0.0.1" //106.13.196.236
	port := "3306"
	database := "douyin"
	username := "root"
	password := "666666"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db_init, err := gorm.Open(mysql.Open(args), &gorm.Config{})

	if err != nil {
		// panic("failed to connect database, err:" + err.Error())
		return err
	}
	db = db_init
	//迁移
	db.Table("user_info").AutoMigrate(&UserInfo{})
	db.Table("follow_info").AutoMigrate(&FollowInfo{})
	// db.Create(&UserInfo{Name: "zhangsan", Password: "11111"})
	// users, err := userDao.getUserInfolist()
	// fmt.Println("%d", users[0].ID)
	return nil

}
