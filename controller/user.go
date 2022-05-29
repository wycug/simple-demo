package controller

import (
	"net/http"

	"github.com/RaymondCode/simple-demo/dao"
	"github.com/gin-gonic/gin"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	// "zhangleidouyin": {
	// 	Id:            1,
	// 	Name:          "zhanglei",
	// 	FollowCount:   10,
	// 	FollowerCount: 5,
	// 	IsFollow:      true,
	// },
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password
	//先从缓存中判断用户是否存在
	if _, exist := UserIsExist(token); exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
		return
	}

	err := dao.CreateUserInfo(username, password)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "create failed"},
		})
		return
	}
	userInfo, err := dao.GetUserInfoByName(username)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "create failed"},
		})
		return
	}
	newUser := userInfoToUser(userInfo)
	usersLoginInfo[token] = newUser
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0},
		UserId:   int64(userInfo.Id),
		Token:    userInfo.Name + userInfo.Password,
	})

}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	if userInfo, isLogin, errMsg := dao.CheckLoginInfo(username, password); isLogin {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   userInfo.Id,
			Token:    userInfo.Token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: errMsg},
		})
	}

}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	if user, exist := UserIsExist(token); exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserIsExist(token string) (User, bool) {
	var user User
	if user, exist := usersLoginInfo[token]; exist {
		return user, true
	}
	if user, err := GetUserByToken(token); err == nil {
		usersLoginInfo[token] = user
		return user, true
	}
	return user, false
}

func GetUserById(id int64) (User, error) {
	userInfo, err := dao.GetUserInfoById(id)
	var user User
	if err != nil {
		return user, err
	}
	user = userInfoToUser(userInfo)

	//等待读取点赞数据
	user.FollowCount = 0
	user.FollowerCount = 0

	return user, nil
}

func GetUserByToken(token string) (User, error) {
	userInfo, err := dao.GetUserInfoByToken(token)
	var user User
	if err != nil {
		return user, err
	}
	user = userInfoToUser(userInfo)

	//等待读取点赞数据
	user.FollowCount = 0
	user.FollowerCount = 0

	return user, nil
}

func userInfoToUser(userInfo dao.UserInfo) User {
	return User{
		Id:            int64(userInfo.Id),
		Name:          userInfo.Name,
		FollowCount:   userInfo.FollowCount,
		FollowerCount: userInfo.FollowerCount,
		IsFollow:      false,
	}
}
