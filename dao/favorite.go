package dao

import (
	"log"
	"strconv"
)

func GetFavoriteList(userId string) ([]VideoInfo, error) {
	var videos []VideoInfo
	uId, _ := strconv.ParseInt(userId, 10, 64)
	db.Raw("select * from video_info where id in (select video_id from favorite_info where user_id = ? and is_favorite = 1)", uId).Scan(&videos)
	return videos, nil
}

func FavoriteAction(userId, videoId string) error {
	var favoriteInfo FavoriteInfo
	var video VideoInfo
	db.Table("favorite_info").Where("user_id = ? and video_id = ?", userId, videoId).First(&favoriteInfo)
	db.Table("video_info").Where("id = ?", videoId).Find(&video)
	if favoriteInfo.VideoId != 0 {
		if favoriteInfo.IsFavorite == 1 {
			return nil
		} else {
			favoriteInfo.IsFavorite = 1
			video.FavoriteCount++
			db.Table("favorite_info").Save(&favoriteInfo)
			db.Table("video_info").Save(&video)
		}
	} else {
		userId, err := strconv.ParseInt(userId, 10, 64)
		if err != nil {
			log.Println("userId parse int fail, err =", err.Error())
			return err
		}
		videoId, err := strconv.ParseInt(videoId, 10, 64)
		if err != nil {
			log.Println("userId parse int fail, err =", err.Error())
			return err
		}
		favoriteInfo.UserId = userId
		favoriteInfo.VideoId = videoId
		favoriteInfo.IsFavorite = 1
		video.FavoriteCount++
		db.Table("favorite_info").Create(&favoriteInfo)
		db.Table("video_info").Save(&video)
	}
	return nil
}

func CancelFavoriteAction(userId, videoId string) error {
	var favoriteInfo FavoriteInfo
	var video VideoInfo
	db.Table("favorite_info").Where("user_id = ? and video_id = ?", userId, videoId).First(&favoriteInfo)
	db.Table("video_info").Where("id = ?", videoId).Find(&video)
	if favoriteInfo.IsFavorite == 0 {
		return nil
	}
	favoriteInfo.IsFavorite = 0
	video.FavoriteCount--
	db.Table("favorite_info").Save(&favoriteInfo)
	db.Table("video_info").Save(&video)
	return nil
}
