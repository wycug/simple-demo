
# simple-demo


## 抖音项目服务端简单示例

具体功能内容参考飞书说明文档

工程无其他依赖，直接编译运行即可

```shell
go run main.go router.go
```
main.go的可以修改启动参数

const URL string = "http://192.168.137.1"
const PORT string = ":8080"
const GROUPPATH string = "/douyin"
const STATICPATH string = "./public"
const SQLHOST string = "106.13.196.236"
const SQLPORT string = "3306"
const SQLDATABASE string = "douyin"
const SQLUSER string = ""
const SQLPASSWORD string = ""
const SQLCHARSET string = "utf8"

### 功能说明

接口功能不完善，仅作为示例

* 用户登录数据保存在内存中，单次运行过程中有效
* 视频上传后会保存到本地 public 目录中，访问时用 127.0.0.1:8080/static/video_name 即可

### 测试数据


# 字节训练营-抖音服务端开发

