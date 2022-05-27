package controller

import (
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

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}

func videoInfoListToVideoList(videoInfoList []dao.VideoInfo) []Video {
	var videos []Video
	for _, videoInfo := range videoInfoList {
		videos = append(videos)
	}
	return videos
}
func videoInfoToVideo(videoInfo dao.VideoInfo) Video {
	return Video{
		Id:            videoInfo.Id,
		Author:        videoInfo.AuthorId,
		PlayUrl:       videoInfo.PlayUrl,
		CoverUrl:      videoInfo.CoverUrl,
		FavoriteCount: videoInfo.FavoriteCount,
		CommentCount:  videoInfo.CommentCount,
		Title:         videoInfo.Title,
	}
}
