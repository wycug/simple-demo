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
	// token := c.Query("token")

	videoInfoList, err := dao.GetVideoList(MaxVideoNums)
	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 1, StatusMsg: err.Error()},
			VideoList: nil,
			NextTime:  time.Now().Unix(),
		})
	}
	videos := videoInfoListToVideoList(videoInfoList)
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videos,
		NextTime:  time.Now().Unix(),
	})

}

func videoInfoListToVideoList(videoInfoList []dao.VideoInfo) []Video {
	var videos []Video
	for _, videoInfo := range videoInfoList {
		video, err := videoInfoToVideo(videoInfo)
		if err != nil {
			log.Printf(err.Error())
			continue
		}
		videos = append(videos, video)
	}
	return videos
}
func videoInfoToVideo(videoInfo dao.VideoInfo) (Video, error) {
	user, err := GetUserById(videoInfo.AuthorId)
	if err != nil {
		return Video{}, err
	}
	return Video{
		Id:            videoInfo.Id,
		Author:        user,
		PlayUrl:       videoInfo.PlayUrl,
		CoverUrl:      videoInfo.CoverUrl,
		FavoriteCount: videoInfo.FavoriteCount,
		CommentCount:  videoInfo.CommentCount,
		Title:         videoInfo.Title,
	}, nil
}
