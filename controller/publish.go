package controller

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/RaymondCode/simple-demo/dao"
	"github.com/gin-gonic/gin"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	title := c.PostForm("title")
	fmt.Println(token)
	if _, exist := UserIsExist(token); !exist {
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
	i := 1
	//绝对路径
	path := "./public/" + filename
	//判断文件名是否存在，如果存在name前加i，i++
	if flag, _ := PathExists(path); flag {
		finalName = strconv.Itoa(i) + finalName
		i++
		fmt.Println("存在")
	}
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	imageName := strings.Replace(filename, ".mp4", ".jpeg", -1)
	saveImage := filepath.Join("./public/", imageName)

	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(saveFile).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		fmt.Println(err)
	}

	img, _ := jpeg.Decode(buf)
	imgw, _ := os.Create(saveImage)
	jpeg.Encode(imgw, img, &jpeg.Options{100})

	ipport := fmt.Sprintf("%s%s/static/", URL, PORT)
	url := ipport + finalName
	imgUrl := ipport + imageName
	err = dao.CreateVideoInfo(user.Id, url, imgUrl, title)
	if err != nil {
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
	id, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
			VideoList: nil,
		})
	}
	videoInfoList, err := dao.GetVideoInfoListById(id)
	videos := VideoInfoListToVideoList(videoInfoList, id)
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
			VideoList: nil,
		})
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})
}


// PathExists 判断文件是否存在，如果存在返回true，不存在返回false
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
