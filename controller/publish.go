package controller

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {

	// title := c.PostForm("title")
	// token := c.PostForm("token")
	// fmt.Println(title, token)
	// // fmt.Println(title)
	// // var user User
	// // if _, exist := usersLoginInfo[token]; !exist {
	// // 	// user_, err := GetUserById(id)

	// // 	// if err != nil {
	// // 	// 	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	// // 	// 	return
	// // 	// }
	// // 	// user = user_
	// // } else {
	// // 	user = usersLoginInfo[token]
	// // }

	// data, err := c.FormFile("data") //318672   //319097
	// // fmt.Println(data.Filename)
	// if err != nil {
	// 	c.JSON(http.StatusOK, Response{
	// 		StatusCode: 1,
	// 		StatusMsg:  err.Error(),
	// 	})
	// 	return
	// }

	// filename := filepath.Base(data.Filename)
	// // user := usersLoginInfo[token]
	// finalName := fmt.Sprintf("%d_%s", 1, filename)
	// saveFile := filepath.Join("./public/", finalName)
	// if err := c.SaveUploadedFile(data, saveFile); err != nil {
	// 	c.JSON(http.StatusOK, Response{
	// 		StatusCode: 1,
	// 		StatusMsg:  err.Error(),
	// 	})
	// 	return
	// }

	// // c.JSON(http.StatusOK, Response{
	// // 	StatusCode: 0,
	// // 	StatusMsg:  finalName + " uploaded successfully",
	// // })
	// c.JSON(http.StatusOK, Response{
	// 	StatusCode: 0,
	// 	StatusMsg:  " uploaded successfully",
	// })

	token := c.PostForm("token")

	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	user := usersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})

}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
