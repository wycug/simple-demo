package dao

import "errors"

func GetVideoList(nums int) ([]VideoInfo, error) {
	var videos []VideoInfo
	result := db.Table("video_info").Find(&videos).Limit(nums)
	if result.RowsAffected == 0 {
		return nil, errors.New("not exist video")
	}
	return videos, nil
}
