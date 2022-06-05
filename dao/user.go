package dao

import "errors"

func GetUserInfoById(id int64) (UserInfo, error) {
	var user UserInfo
	result := db.Table("user_info").Where("id = ?", id).Find(&user)
	if result.Error != nil {
		return user, result.Error
	}
	if result.RowsAffected == 0 {
		return user, errors.New("user not exist")
	}
	return user, nil
}

func GetUserInfoByName(name string) (UserInfo, error) {
	var user UserInfo
	result := db.Table("user_info").Where("name = ?", name).Find(&user)
	if result.Error != nil {
		return user, result.Error
	}
	if result.RowsAffected == 0 {
		return user, errors.New("user not exist")
	}
	return user, nil
}

func GetUserInfoByToken(token string) (UserInfo, error) {
	var user UserInfo
	result := db.Table("user_info").Where("token = ?", token).Find(&user)
	if result.Error != nil {
		return user, result.Error
	}
	if result.RowsAffected == 0 {
		return user, errors.New("user not exist")
	}
	return user, nil
}

func GetUserInfolist() ([]UserInfo, error) {
	var users []UserInfo
	result := db.Table("user_info").Find(&users)
	if result.Error != nil {
		return users, result.Error
	}
	if result.RowsAffected == 0 {
		return users, errors.New("not exist user")
	}
	return users, nil
}

func CreateUserInfo(name, password string) error {
	result := db.Table("user_info").Create(&UserInfo{Name: name, Password: password, Token: name + password})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not exist")
	}
	return nil
}

func CheckLoginInfo(name, password string) (UserInfo, bool, string) {
	var user UserInfo
	result := db.Table("user_info").Where("name = ?", name).Find(&user)
	if result.Error != nil {
		return user, false, "请稍后重试"
	}
	if result.RowsAffected == 0 {
		return user, false, "用户名不存在"
	}
	if user.Password == password {
		return user, true, ""
	} else {
		return user, false, "密码错误"
	}

}
