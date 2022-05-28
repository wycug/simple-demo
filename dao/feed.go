package dao

import "errors"

func GetVideoList(nums int) ([]VideoInfo, error) {
	var videos []VideoInfo
	//Order按创建时间倒叙排序
	result := db.Table("video_info").Order("created_at desc").Find(&videos).Limit(nums)
	if result.RowsAffected == 0 {
		return nil, errors.New("not exist video")
	}
	return videos, nil
}
