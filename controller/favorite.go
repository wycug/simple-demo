package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/RaymondCode/simple-demo/dao"
	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid

type FavoriteActionListResponse struct {
	Response
	Videos []dao.VideoInfo
}

type FavoriteActionListRequest struct {
	Token      string `form:"token"`
	UserId     string `form:"user_id"`
	VideoId    string `form:"video_id"`
	ActionType string `form:"action_type"`
}

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	// token := c.Param("token")
	// userId := c.Param("user_id")
	// videoId := c.Param("video_id")
	// actionType := c.Param("action_type")

	var params FavoriteActionListRequest
	c.BindQuery(&params)
	//fmt.Println(params)
	if user, exist := UserIsExist(params.Token); exist {
		if params.ActionType == "1" {
			err := dao.FavoriteAction(strconv.FormatInt(user.Id, 10), params.VideoId)
			if err != nil {
				return
			}
			c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "favorite action success"})
		} else if params.ActionType == "2" {
			err := dao.CancelFavoriteAction(params.UserId, params.VideoId)
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
	userId := c.Query("user_id")
	if _, exist := UserIsExist(token); exist {
		videos, err := dao.GetFavoriteList(userId)
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
