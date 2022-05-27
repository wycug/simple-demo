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

var userIdSequence = int64(1)

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
	if _, exist := UserIsExist(token, "", username); exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	}

	err := dao.CreateUserInfo(username, password)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "create failed"},
		})
	}
	userInfo, err := dao.GetUserInfoByName(username)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "create failed"},
		})
	}
	newUser := userInfoToUser(userInfo)
	usersLoginInfo[token] = newUser
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0},
		UserId:   int64(userInfo.ID),
		Token:    userInfo.Name + userInfo.Password,
	})

}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if user, exist := UserIsExist(token, "", username); exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	id := c.Query("id")
	token := c.Query("token")

	if user, exist := UserIsExist(token, id, ""); exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     user,
		})
	} else if userInfo, err := dao.GetUserInfoById(id); err != nil {
		//暂时未处理点赞信息
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User: User{
				Id:            int64(userInfo.ID),
				Name:          userInfo.Name,
				FollowCount:   0,
				FollowerCount: 0,
				IsFollow:      false,
			},
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func GetUserById(id string) (User, error) {
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

func UserIsExist(token, user_id, user_name string) (User, bool) {
	var user User
	if user, exist := usersLoginInfo[token]; exist {
		return user, true
	}
	if user_id != "" {
		if userInfo, err := dao.GetUserInfoById(user_id); err == nil {
			user = User{
				Id:            int64(userInfo.ID),
				Name:          userInfo.Name,
				FollowCount:   userInfo.FollowCount,
				FollowerCount: userInfo.FollowerCount,
				IsFollow:      false,
			}
			usersLoginInfo[token] = user
			return user, true
		}
	}
	if user_name != "" {
		if userInfo, err := dao.GetUserInfoByName(user_name); err == nil {
			user = userInfoToUser(userInfo)
			usersLoginInfo[token] = user
			return user, true
		}
	}

	return user, false
}

func userInfoToUser(userInfo dao.UserInfo) User {
	return User{
		Id:            int64(userInfo.ID),
		Name:          userInfo.Name,
		FollowCount:   userInfo.FollowCount,
		FollowerCount: userInfo.FollowerCount,
		IsFollow:      false,
	}
}
