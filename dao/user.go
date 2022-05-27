package dao

type UserDao struct {
}

var userDao UserDao

func GetUserInfoById(id string) (UserInfo, error) {
	var user UserInfo
	result := db.Table("user_info").Where("id = ?", id).Find(&user)
	if result.RowsAffected != 0 {
		return user, result.Error
	}
	return user, nil
}
func GetUserInfoByName(name string) (UserInfo, error) {
	var user UserInfo
	result := db.Table("user_info").Where("name = ?", name).Find(&user)
	if result.RowsAffected == 0 {
		return user, result.Error
	}
	return user, nil
}

func GetUserInfolist() ([]UserInfo, error) {
	var users []UserInfo
	result := db.Table("user_info").Find(&users)
	if result.RowsAffected == 0 {
		return nil, result.Error
	}
	return users, nil
}

func CreateUserInfo(name, password string) error {
	result := db.Table("user_info").Create(&UserInfo{Name: name, Password: password})
	if result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}
