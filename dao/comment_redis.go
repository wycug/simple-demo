package dao

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type CommentRedis struct{}

var (
	commentRedis CommentRedis
	rdb          *redis.Client
	ctx          = context.Background()
)

func Addcomment_redis(userID, videoID int64, text, res string) error {
	//缓冲管道
	pipe := rdb.TxPipeline()

	key1 := fmt.Sprintf("%v:comment_userID", userID)
	if err := pipe.SAdd(ctx, key1, userID).Err(); err != nil {
		return err
	}

	key2 := fmt.Sprintf("%v:comment_videoID", videoID)
	if err := pipe.SAdd(ctx, key2, videoID).Err(); err != nil {
		return err
	}

	key3 := fmt.Sprintf("%v:comment_text", text)
	if err := pipe.SAdd(ctx, key3, text).Err(); err != nil {
		return err
	}

	key4 := fmt.Sprintf("%v:comment_res", res)
	if err := pipe.SAdd(ctx, key4, res).Err(); err != nil {
		return err
	}

	if _, err := pipe.Exec(ctx); err != nil {
		return err
	}

	return nil
}

//删除评论
func Deletecomment_redis(id int64) error {
	pipe := rdb.TxPipeline()

	key1 := fmt.Sprintf("%v:comment_id", id)
	if err := pipe.SAdd(ctx, key1, id).Err(); err != nil {
		return err
	}

	if _, err := pipe.Exec(ctx); err != nil {
		return err
	}

	return nil
}

// //显示评论
// func List_redis(videoID int64) ([]CommentInfo, error) {
// 	var comments []CommentInfo
// 	//Order按创建时间倒叙排序
// 	result := db.Table("comment_info").Where("video_id = ?", videoID).Order("created_at desc").Find(&comments)
// 	if result.RowsAffected == 0 {
// 		return nil, errors.New("not exist video")
// 	}
// 	return comments, nil
// }
