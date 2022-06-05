package controller

import (
	"net/http"
	"strconv"

	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/util/constant"
	"github.com/gin-gonic/gin"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

//实例化service
var (
	followService     = service.NewRelationServiceInstance()
	followListService = service.NewFollowerDaoInstance()
)

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")

	if user, exist := UserIsExist(token); exist {
		//判断操作类型 如果是关注则走关注的逻辑， 如果是取消关注则走取消关注的逻辑
		actionType, err := strconv.ParseInt(c.Query("action_type"), 10, 64)
		if err != nil {
			c.JSON(http.StatusOK,
				Response{
					StatusCode: constant.RequestParamError,
					StatusMsg:  err.Error(),
				})
			return
		}
		followFromID := user.Id
		followToID, err := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusOK,
				Response{
					StatusCode: constant.RequestParamError,
					StatusMsg:  err.Error(),
				})
			return
		}

		if actionType == constant.FOLLOW {
			//关注操作
			if err := FollowAop(followFromID, followToID); err != nil {
				c.JSON(http.StatusOK,
					Response{
						StatusCode: constant.FollowFailed,
						StatusMsg:  err.Error(),
					})
				return
			}
		} else if actionType == constant.UNFOLLOW {
			//取消关注操作
			if err := UnFollowAop(followFromID, followToID); err != nil {
				c.JSON(http.StatusOK,
					Response{
						StatusCode: constant.UnfollowFailed,
						StatusMsg:  err.Error(),
					})
				return
			}
		} else {
			c.JSON(http.StatusOK,
				Response{
					StatusCode: constant.OptParameterError,
					StatusMsg:  constant.Msg(constant.OptParameterError),
				})
			return
		}
		c.JSON(http.StatusOK,
			Response{
				StatusCode: constant.Success,
				StatusMsg:  constant.Msg(constant.Success),
			})
	} else {
		c.JSON(http.StatusOK,
			Response{
				StatusCode: constant.UserNotExistError,
				StatusMsg:  constant.Msg(constant.UserNotExistError),
			})
	}
}

// FollowList all users have same follow list
//  关注列表
func FollowList(c *gin.Context) {
	token := c.Query("token")

	if user, exist := UserIsExist(token); exist {
		// 调用Service层里的获取关注列表函数
		// 获取用户id
		userId := user.Id
		followInfo, _, _, err := followListService.FollowerList(userId, constant.FOLLOWLIST)
		if err != nil {
			c.JSON(http.StatusOK,
				Response{
					StatusCode: constant.GetListFailed,
					StatusMsg:  err.Error(),
				})
			return
		}
		var users []User

		for i := 0; i < len(followInfo); i++ {
			// 拷贝到user变量
			var user User
			user.Id = followInfo[i].ID
			user.FollowCount = followInfo[i].FollowCount
			user.FollowerCount = followInfo[i].FollowerCount
			user.IsFollow = true
			user.Name = followInfo[i].Name

			users = append(users, user)
		}

		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: constant.Success,
				StatusMsg:  constant.Msg(constant.Success),
			},
			UserList: users,
		})
	} else {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: constant.UserNotExistError,
				StatusMsg:  constant.Msg(constant.UserNotExistError),
			},
			UserList: nil,
		})
	}
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	token := c.Query("token")

	if user, exist := UserIsExist(token); exist {
		// 	调用Service层里的获取粉丝列表函数
		// 获取用户id
		userId := user.Id
		followerInfo, _, isFollow, err := followListService.FollowerList(userId, constant.FANSLIST)
		if err != nil {
			c.JSON(http.StatusOK,
				Response{
					StatusCode: constant.GetListFailed,
					StatusMsg:  err.Error(),
				})
			return
		}

		var users []User
		for i := 0; i < len(followerInfo); i++ {
			var user User
			user.Id = followerInfo[i].ID
			user.Name = followerInfo[i].Name
			user.FollowCount = followerInfo[i].FollowCount
			user.FollowerCount = followerInfo[i].FollowerCount
			// 如果当前userid 在isFollow里，返回true，否则是false
			user.IsFollow = IsFollow(isFollow, user)
			users = append(users, user)
		}

		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: constant.Success,
				StatusMsg:  constant.Msg(constant.Success),
			},
			UserList: users,
		})
	} else {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: constant.UserNotExistError,
				StatusMsg:  constant.Msg(constant.UserNotExistError),
			},
			UserList: nil,
		})
	}
}

func IsFollow(isFollow []int64, user User) bool {
	for i := 0; i < len(isFollow); i++ {
		if isFollow[i] == user.Id {
			return true
		}
	}
	return false
}

func FollowAop(followFromID, followToID int64) error {
	BeforeUserCacheUpdateById(followFromID)
	BeforeUserCacheUpdateById(followToID)
	_, err := followService.Follow(followFromID, followToID)
	if err != nil {
		return err
	}
	AfterUserCacheUpdateById(followFromID)
	AfterUserCacheUpdateById(followToID)
	return nil
}

func UnFollowAop(followFromID, followToID int64) error {
	BeforeUserCacheUpdateById(followFromID)
	BeforeUserCacheUpdateById(followToID)
	_, err := followService.UnFollow(followFromID, followToID)
	if err != nil {
		return err
	}
	AfterUserCacheUpdateById(followFromID)
	AfterUserCacheUpdateById(followToID)
	return nil
}
