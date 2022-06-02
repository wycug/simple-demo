package dao

func GetVideoList(nums int) ([]VideoInfo, error) {
	var videos []VideoInfo
	// result := db.Table("video_info").Find(&videos).Limit(nums)
	//Order按创建时间倒叙排序
	db.Table("video_info").Order("created_at desc").Find(&videos).Limit(nums)
	return videos, nil
}
