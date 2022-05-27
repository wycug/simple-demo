package dao

func GetFavoriteList(userId int64) ([]VideoInfo, error) {
	var videos []VideoInfo
	db.Raw("select * from video_info where video_id in select video_id from favorite_info where user_id = ?", userId).Scan(&videos)
	return videos, nil
}

func FavoriteAction(userId, videoId int64) error {
	var favoriteInfo FavoriteInfo
	var video VideoInfo
	db.Table("favorite_info").Where("user_id = ? and & video_id = ?", userId, videoId).First(&favoriteInfo)
	db.Table("video_info").Where("video_id = ?", video).Find(&video)
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
		favoriteInfo.UserId = userId
		favoriteInfo.VideoId = videoId
		favoriteInfo.IsFavorite = 1
		video.FavoriteCount++
		db.Table("favorite_info").Create(&favoriteInfo)
		db.Table("video_info").Save(&video)
	}
	return nil
}

func CancelFavoriteAction(userId, videoId int64) error {
	var favoriteInfo FavoriteInfo
	var video VideoInfo
	db.Table("favorite_info").Where("user_id = ? and & video_id = ?", userId, videoId).First(&favoriteInfo)
	db.Table("video_info").Where("video_id = ?", video).Find(&video)
	favoriteInfo.IsFavorite = 0
	video.FavoriteCount--
	db.Table("favorite_info").Save(&favoriteInfo)
	db.Table("video_info").Save(&video)
	return nil
}
