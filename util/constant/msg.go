/**
@author: Junming ZHANG, Haojun YIN
@date: 2022/5/28
*/

package constant

var msg = []string{
	"Success",                          // 成功
	"RequestParamError",                // 请求参数错误
	"UnknownError",                     // 未知错误
	"UsernameHasExistedError",          //用户名已存在
	"GenerateTokenError",               // 生成token出错
	"GetIdByTokenError",                // 通过token获取id出错
	"UserNotExistOrPasswordWrongError", // 用户名不存在或密码错误
	"LoadFileError",                    // 加载文件出错
	"SaveUploadedFileError",            // 保存文件出错
	"Opt Parameter error",              // 操作参数异常
	"User doesn't exist",               // 当前用户不存在
	"Already followed",                 // 当前用户已关注
	"Not follow yet",                   // 当前用户还未关注
	"Follow failed",                    // 关注失败
	"Unfollow failed",                  // 取消关注失败
	"Get list failed",                  // 获取列表失败
}

func Msg(code int) string {
	if code < 0 || code >= len(msg) {
		return ""
	}
	return msg[code]
}
