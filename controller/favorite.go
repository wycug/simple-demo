package controller

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type FavoriteActionListResponse struct {
	Response
	Videos []dao.VideoInfo
}

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	userId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		log.Println("userId format to int64 fail, err =", err.Error())
		return
	}
	videoId, err := strconv.Atoi(c.Query("video_id"))
	if err != nil {
		log.Println("videoId format to int64 fail, err =", err.Error())
		return
	}
	actionType, err := strconv.Atoi(c.Query("action_type"))
	if err != nil {
		log.Println("actionType format to int64 fail, err =", err.Error())
		return
	}

	if _, exist := usersLoginInfo[token]; exist {
		if actionType == 1 {
			err = dao.FavoriteAction(int64(userId), int64(videoId))
			if err != nil {
				return
			}
			c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "favorite action success"})
		} else if actionType == 2 {
			err = dao.CancelFavoriteAction(int64(userId), int64(videoId))
			if err != nil {
				return
			}
			c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "cancel favorite action success"})
		} else {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "action_type is not valid"})
		}
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	userId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		log.Println("userId format to int64 fail, err =", err.Error())
		c.JSON(http.StatusOK, Response{
			StatusCode: 0,
			StatusMsg:  "userId is invalid",
		})
		return
	}
	if _, exist := UserIsExist(token, "", ""); exist {
		videos, err := dao.GetFavoriteList(int64(userId))
		if err != nil {
			log.Println("get favorite list fail, err =", err.Error())
			c.JSON(http.StatusOK, FavoriteActionListResponse{
				Response: Response{
					StatusCode: 1,
					StatusMsg:  "get favorite list fail",
				},
				Videos: nil,
			})
			return
		}
		c.JSON(http.StatusOK, FavoriteActionListResponse{
			Response: Response{
				StatusCode: 0,
			},
			Videos: videos,
		})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}
