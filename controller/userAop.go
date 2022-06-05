package controller

import "github.com/RaymondCode/simple-demo/dao"

func BeforeUserCacheUpdateByToken(token string) {
	if _, exist := usersLoginInfo[token]; exist {
		delete(usersLoginInfo, token)
	}
}

func AfterUserCacheUpdateByToken(token string) {
	user, _ := GetUserByToken(token)
	usersLoginInfo[token] = user
}

func BeforeUserCacheUpdateById(id int64) {
	userInfo, _ := dao.GetUserInfoById(id)
	if _, exist := usersLoginInfo[userInfo.Token]; exist {
		delete(usersLoginInfo, userInfo.Token)
	}
}

func AfterUserCacheUpdateById(id int64) {
	userInfo, _ := dao.GetUserInfoById(id)
	usersLoginInfo[userInfo.Token] = userInfoToUser(userInfo)
}
