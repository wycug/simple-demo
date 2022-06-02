/**
@author: Junming ZHANG, Haojun YIN
@date: 2022/5/28
*/

package constant

const (
	Success                          = iota // 成功
	RequestParamError                       // 请求参数错误
	UnknownError                            // 未知错误
	UsernameHasExistedError                 // 用户名已存在
	GenerateTokenError                      // 生成token出错
	GetIdByTokenError                       // 通过token获取id出错
	UserNotExistOrPasswordWrongError        // 用户名不存在或密码错误
	LoadFileError                           // 加载文件出错
	SaveUploadedFileError                   // 保存文件出错
	OptParameterError                       // 操作参数异常
	UserNotExistError                       // 当前用户不存在
	AlreadyFollowedError                    // 当前用户已关注
	NotFollowYetError                       // 当前用户还未关注
	FollowFailed                            // 关注失败
	UnfollowFailed                          // 取消关注失败
	GetListFailed                           // 获取列表失败
)
