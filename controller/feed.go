package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/RaymondCode/simple-demo/dao"
	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

const MaxVideoNums int = 30

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	//可选参数
	// lastTime := c.Query("latest_time")
	token := c.Query("token")
	var user User
	user, _ = UserIsExist(token)
	videoInfoList, err := dao.GetVideoList(MaxVideoNums)
	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 1, StatusMsg: err.Error()},
			VideoList: nil,
			NextTime:  time.Now().Unix(),
		})
	}
	videos := VideoInfoListToVideoList(videoInfoList, user.Id)
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videos,
		NextTime:  time.Now().Unix(),
	})
}

func VideoInfoListToVideoList(videoInfoList []dao.VideoInfo, user_id int64) []Video {
	var videos []Video
	for _, videoInfo := range videoInfoList {
		video, err := VideoInfoToVideo(videoInfo, user_id)
		if err != nil {
			log.Printf(err.Error())
			continue
		}
		videos = append(videos, video)
	}
	return videos
}
func VideoInfoToVideo(videoInfo dao.VideoInfo, user_id int64) (Video, error) {
	author, _ := GetUserById(videoInfo.AuthorId)
	return Video{
		Id:            videoInfo.Id,
		Author:        author,
		PlayUrl:       videoInfo.PlayUrl,
		CoverUrl:      videoInfo.CoverUrl,
		FavoriteCount: videoInfo.FavoriteCount,
		CommentCount:  videoInfo.CommentCount,
		IsFavorite:    dao.IsFavorite(user_id, videoInfo.Id),
		Title:         videoInfo.Title,
	}, nil
}
