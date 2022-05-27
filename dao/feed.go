package dao

func GetVideoList(nums int) ([]VideoInfo, error) {
	var videos []VideoInfo
	result := db.Table("video_info").Find(&videos).Limit(nums)
	if result.RowsAffected == 0 {
		return nil, result.Error
	}
	return videos, nil
}
