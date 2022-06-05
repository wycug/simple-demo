package dao

import (
	"gorm.io/gorm"
)

func CreateVideoInfo(authorId int64, playUrl, coverUrl, title string) error {
	result := db.Table("video_info").Create(&VideoInfo{
		Model:    gorm.Model{},
		AuthorId: authorId,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
		Title:    title,
	})
	if result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}

func GetVideoInfoListById(authorId int64) ([]VideoInfo, error) {
	var videoInfoList []VideoInfo
	db.Table("video_info").Where("author_id = ?", authorId).Order("created_at desc").Find(&videoInfoList)
	return videoInfoList, nil
}
