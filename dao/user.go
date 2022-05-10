package dao

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	Id       uint   `gorm:"primaryKey"`
	Name     string `gorm:"varchar(50);not null;index;unique"`
	Password string `gorm:"varchar(50);not null"`
}

func (UserInfo) TableName() string {
	return "user_info"
}

type UserDao struct {
}

var userDao UserDao

func GetUserInfoById(id string) (UserInfo, error) {
	var user UserInfo
	result := db.Where("id = ?", id).First(&user)
	if result.RowsAffected != 0 {
		return user, result.Error
	}
	return user, nil
}
func GetUserInfoByName(name string) (UserInfo, error) {
	var user UserInfo
	result := db.Where("name = ?", name).First(&user)
	if result.RowsAffected != 0 {
		return user, result.Error
	}
	return user, nil
}

func GetUserInfolist() ([]UserInfo, error) {
	var users []UserInfo
	result := db.Find(&users)
	if result.RowsAffected == 0 {
		return nil, result.Error
	}
	return users, nil
}

func CreateUserInfo(name, password string) error {
	result := db.Create(&UserInfo{Name: name, Password: password})
	if result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}
