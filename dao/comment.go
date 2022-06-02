package dao

import (
	"time"
)

type CommentDao struct{}

var commentDao CommentDao

//添加评论  gorm  curd语句
func Addcomment(userID, videoID int64, text, res string) error {
	var video VideoInfo
	//添加评论到数据库
	result := db.Table("comment_info").Create(&CommentInfo{UserID: userID, VideoID: videoID, Content: text, Comment_date: res})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return result.Error
	}
	//评论总数+1，更新video
	db.Table("video_info").Where("id = ?", videoID).Find(&video)
	video.CommentCount++
	db.Table("video_info").Save(&video)

	return nil
}

//删除评论
func Deletecomment(id int64) error {
	var video VideoInfo
	var comment CommentInfo
	//评论总数-1，更新video
	db.Table("comment_info").Where("id = ?", id).Find(&comment)
	db.Table("video_info").Where("id = ?", comment.VideoID).Find(&video)
	video.CommentCount--
	db.Table("video_info").Save(&video)

	//防止两个操作发生冲突:需要靠评论的id找到对应的视频
	time.Sleep(100)

	//软删除
	// db.Table("comment_info").Where("id = ? ", id).Unscoped().Delete(&CommentInfo{})
	db.Table("comment_info").Where("id = ? ", id).Delete(&CommentInfo{})

	return nil
}

//显示评论
func List(videoID int64) ([]CommentInfo, error) {
	var comments []CommentInfo
	//Order按创建时间倒叙排序
	db.Table("comment_info").Where("video_id = ?", videoID).Order("created_at desc").Find(&comments)
	return comments, nil
}

// func Comment_count(videoID int64) error {
// 	var count CommentInfo
// 	result := db.Table("comment_info").Where("video_id = ? ", videoID).Count(&count)
// 	if result.RowsAffected == 0 {
// 		return errors.New("not exist video")
// 	}
// 	fmt.Println(count)
// 	db.Table("video_info").Where("id = ? ", videoID).Update("comment_count", count)
// 	return nil
// }
