package dao

import "errors"

type CommentDao struct{}

var commentDao CommentDao

//添加评论  gorm  curd语句
func Addcomment(userID, videoID int64, text, res string) error {
	result := db.Table("comment_info").Create(&CommentInfo{UserID: userID, VideoID: videoID, Content: text, Comment_date: res})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}

//删除评论
func Deletecomment(id int64) error {
	result := db.Table("comment_info").Where("id = ? ", id).Delete(&CommentInfo{})
	if result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}

//显示评论
func List(videoID int64) ([]CommentInfo, error) {
	var comments []CommentInfo
	//Order按创建时间倒叙排序
	result := db.Table("comment_info").Where("video_id = ?", videoID).Order("created_at desc").Find(&comments)
	if result.RowsAffected == 0 {
		return nil, errors.New("not exist video")
	}
	return comments, nil
}
