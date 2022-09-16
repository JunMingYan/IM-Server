package constant

const (
	SuccessCode     = 200
	FailCode        = 400
	ServerCode      = 500
	GetMyFriendCode = 2000

	VerifyCodeInvalidCode = 1000
	LoginSuccessCode      = 1001
	Blacklist             = 1002
	RegisterSuccessCode   = 1003
)

const (
	SessionID = "Session"
)

const (
	RegisterSuccess = "注册成功"
	UpdateSuccess   = "更新成功"
	LoginSuccess    = "登录成功"
	ApiSuccess      = "成功请求接口"

	UpdateFailed         = "更新失败"
	GetVerifyCodeSuccess = "成功获取验证码"
	GetVerifyCodeError   = "获取验证码失败,请稍后再试"
	VerifyCodeError      = "验证码不存在,请重新输入"
	PasswordError        = "密码错误,请输入密码"
	SaltError            = "加密出错"
	UpdatePasswordError  = "新密码不能和旧密码相同,请重新输入"
	ServerError          = "服务器内部错误"
	GetParamsError       = "获取参数错误"
	UserNameError        = "用户名已存在,请重新输入"
	NotUserNameError     = "用户名不存在,请重新输入"
	RequestHeaderNull    = "请求头中auth为空"
	RequestHeaderError   = "请求头中auth格式有误"
	TokenError           = "无效的token"
)

type Response struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}
