package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/RaymondCode/simple-demo/dao"
	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid

type FavoriteActionListResponse struct {
	Response
	Videos []Video `json:"video_list"`
}

type FavoriteActionListRequest struct {
	Token      string `json:"token"`
	UserId     string `json:"user_id"`
	VideoId    string `json:"video_id"`
	ActionType string `json:"action_type"`
}

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	//var params *FavoriteActionListRequest
	//err := c.BindQuery(params)
	//params.Token = c.Param("token")
	//params.UserId = c.Param("user_id")
	//params.VideoId = c.Param("video_id")
	//params.ActionType = c.Param("action_type")
	//if err != nil {
	//	log.Println("bind param fail, err =", err.Error())
	//	return
	//}
	token := c.Query("token")
	userId := c.Query("user_id")
	videoId := c.Query("video_id")
	actionType := c.Query("action_type")
	fmt.Println(token, userId, videoId, actionType)
	if user, exist := UserIsExist(token); exist {
		if actionType == "1" {
			err := dao.FavoriteAction(strconv.FormatInt(user.Id, 10), videoId)
			if err != nil {
				return
			}
			c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "favorite action success"})
		} else if actionType == "2" {
			err := dao.CancelFavoriteAction(strconv.FormatInt(user.Id, 10), videoId)
			if err != nil {
				log.Println("cancel favorite video fail, err =", err.Error())
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
		id, _ := strconv.ParseInt(userId, 10, 64)
		res := make([]Video, 0)
		for i := 0; i < len(videos); i++ {
			tmp := VideoInfo2Video(videos[i], id)
			res = append(res, tmp)
		}
		c.JSON(http.StatusOK, FavoriteActionListResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "get favorite video success",
			},
			Videos: res,
		})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

func VideoInfo2Video(video dao.VideoInfo, userId int64) Video {
	userInfo, _ := dao.GetUserInfoById(userId)
	return Video{
		Id: video.Id,
		Author: User{
			Id:            userId,
			Name:          userInfo.Name,
			FollowCount:   userInfo.FollowCount,
			FollowerCount: userInfo.FollowerCount,
			IsFollow:      userInfo.IsFollow,
		},
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		IsFavorite:    true,
		Title:         video.Title,
	}
}

// func CancelFavoriteActionAop(user User, params FavoriteActionListRequest) error {
// 	BeforeUserCacheUpdateByToken(params.Token)
// 	err := dao.FavoriteAction(strconv.FormatInt(user.Id, 10), params.VideoId)
// 	if err != nil {
// 		return err
// 	}
// 	AfterUserCacheUpdateByToken(params.Token)
// 	return nil
// }

// func FavoriteActionAop(user User, params FavoriteActionListRequest) error {
// 	BeforeUserCacheUpdateByToken(params.Token)
// 	err := dao.CancelFavoriteAction(params.UserId, params.VideoId)
// 	if err != nil {
// 		return err
// 	}
// 	AfterUserCacheUpdateByToken(params.Token)
// 	return nil
// }


// func CancelFavoriteActionAop(user User, params FavoriteActionListRequest) error {
// 	BeforeUserCacheUpdateByToken(params.Token)
// 	err := dao.FavoriteAction(strconv.FormatInt(user.Id, 10), params.VideoId)
// 	if err != nil {
// 		return err
// 	}
// 	AfterUserCacheUpdateByToken(params.Token)
// 	return nil
// }

// func FavoriteActionAop(user User, params FavoriteActionListRequest) error {
// 	BeforeUserCacheUpdateByToken(params.Token)
// 	err := dao.CancelFavoriteAction(params.UserId, params.VideoId)
// 	if err != nil {
// 		return err
// 	}
// 	AfterUserCacheUpdateByToken(params.Token)
// 	return nil
// }
