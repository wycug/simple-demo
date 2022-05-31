package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/RaymondCode/simple-demo/dao"
	"github.com/gin-gonic/gin"
)

type CommentListResponse struct {
	Response

	CommentList []Comment `json:"comment_list,omitempty"` //omitempty作用是在json数据结构转换时，当该字段的值为该字段类型的零值时，忽略该字段。
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")

	if _, exist := UserIsExist(token); exist {

		if actionType == "1" {
			actionAdd(c)
			return
		} else {
			actionDel(c)
			return
		}
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}

}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {

	//查询videoID的所有评论，将查询到的信息封装到comment类中，然后再通过response返回前端
	videoID, err := strconv.ParseInt(c.Query("video_id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusOK, CommentListResponse{
			Response:    Response{StatusCode: 1, StatusMsg: "没找到视频"},
			CommentList: nil,
		})
	}

	commentList, err := dao.List(videoID)
	if err != nil {
		c.JSON(http.StatusOK, CommentListResponse{
			Response:    Response{StatusCode: 1, StatusMsg: "没找到视频"},
			CommentList: nil,
		})
	}
	//找到了返回列表
	comments := comment_to_list(commentList)
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: comments,
	})

	// userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	// if err != nil {
	// 	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "没查到用户id"})
	// }

	// token := c.Query("token")
	// if user, exist := usersLoginInfo[token]; exist {
	// 	text := c.Query("comment_text")
	// 	res := time.Now().Format("2006-01-02 15:04:05")
	// 	var DemoComments = []Comment{
	// 		{
	// 			Id:         userID,
	// 			User:       user,
	// 			Content:    text,
	// 			CreateDate: res,
	// 		},
	// 	}

	// 	c.JSON(http.StatusOK, CommentListResponse{
	// 		Response:    Response{StatusCode: 0},
	// 		CommentList: DemoComments,
	// 	})

	// } else {
	// 	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "当前没有评论"})
	// }
}

//产生一个评论列表
func comment_to_list(commentInfoList []dao.CommentInfo) []Comment {
	var comments []Comment

	for _, commentInfo := range commentInfoList {
		comment, err := newcommnet(commentInfo)
		if err != nil {
			log.Printf(err.Error())
			continue
		}
		comments = append(comments, comment)
	}
	return comments

}

//新的评论
func newcommnet(commentInfo dao.CommentInfo) (Comment, error) {
	user, err := GetUserById(commentInfo.UserID)
	if err != nil {
		return Comment{}, err
	}

	return Comment{
		Id:         commentInfo.Id,
		User:       user,
		Content:    commentInfo.Content,
		CreateDate: commentInfo.Comment_date,
	}, nil
}

// 新增评论
func actionAdd(c *gin.Context) {
	token := c.Query("token")
	//拿到user结构体信息
	if user, exist := UserIsExist(token); exist {
		userID := user.Id
		// fmt.Println(userID)
		text := c.Query("comment_text")
		res := time.Now().Format("2006-01-02 15:04:05")

		videoID, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: Response{StatusCode: 1, StatusMsg: "没找到视频"},
			})
		}

		//将数据添加到数据库中
		err = dao.Addcomment(userID, videoID, text, res)

		if err != nil {
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: Response{StatusCode: 1, StatusMsg: "create failed"},
			})
			return
		}
		c.JSON(http.StatusOK, CommentActionResponse{
			Response: Response{StatusCode: 0},
			Comment: Comment{
				Id:         userID,
				User:       user,
				Content:    text,
				CreateDate: res,
			}})

	} else {
		c.JSON(http.StatusOK, CommentActionResponse{
			Response: Response{StatusCode: 1, StatusMsg: "新增评论失败"},
		})
	}
}

// 只处理删除评论的操作
func actionDel(c *gin.Context) {
	commentID, err := strconv.ParseInt(c.Query("comment_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, CommentActionResponse{
			Response: Response{StatusCode: 1, StatusMsg: "删除评论失败"},
		})

	}
	fmt.Println(commentID)
	//在数据库中删除此评论
	dao.Deletecomment(commentID)
	// resp := dao.Deletecomment(commentID)
	// c.JSON(http.StatusOK, resp)
	c.JSON(http.StatusOK, Response{StatusCode: 0})
	//实时更新列表
}