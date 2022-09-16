package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"server/constant"
	"server/internal/models"
	"server/internal/service"
	"server/internal/verifyCode"
	"server/pkg/jwt"
)

// 1.生成验证码(captchaID,code)
// 2.验证码保存到redis,过期时间60s key->value(captchaID->code)
// 3.返回验证码
func GetCaptcha(c *gin.Context) {
	code, err := verifyCode.GetVerifyCode(c)
	if err != nil {
		logrus.Info(err)
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.VerifyCodeInvalidCode,
			Message: constant.GetVerifyCodeError,
			Data:    nil,
		})
		return
	}
	//
	c.JSON(http.StatusOK, constant.Response{
		Code:    constant.SuccessCode,
		Message: constant.GetVerifyCodeSuccess,
		Data:    code,
	})
	return
}

// 1.解析参数
// 2.检查验证码
// 3.检查参数(昵称、手机号)是否重复
// 4.检查密码是否相同
// 5.注册用户
// 6.响应
func RegisterHandler(c *gin.Context) {
	var user models.UserRegisterForm
	if err := c.ShouldBind(&user); err != nil {
		logrus.Info(err)
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: constant.GetParamsError,
			Data:    nil,
		})
		return
	}
	//
	if !verifyCode.Verify(c, user.Code) {
		logrus.Info(constant.VerifyCodeError)
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.VerifyCodeInvalidCode,
			Message: constant.VerifyCodeError,
			Data:    nil,
		})
		return
	}
	//
	if err := service.Register(user); err != nil {
		logrus.Error(err)
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	//5.响应
	data := make(map[string]interface{})
	data["userCode"] = user.UserName
	c.JSON(http.StatusOK, constant.Response{
		Code:    constant.RegisterSuccessCode,
		Message: constant.RegisterSuccess,
		Data:    data,
	})
	return
}

// 1.解析参数
// 2.检查参数是否存在
// 3.检查密码是否正确
// 4.生成token
// 5.响应
func LoginHandler(c *gin.Context) {
	var user models.UserLoginForm
	err := c.Bind(&user)
	if err != nil {
		logrus.Info(err)
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: constant.GetParamsError,
			Data:    nil,
		})
		return
	}
	//
	if !verifyCode.Verify(c, user.Code) {
		logrus.Info(constant.VerifyCodeError)
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.VerifyCodeInvalidCode,
			Message: constant.VerifyCodeError,
			Data:    nil,
		})
		return
	}
	//
	userInfo, groups, notes, err := service.Login(user)
	if err != nil {
		logrus.Info(err)
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	//
	data := make(map[string]interface{})
	data["userInfo"] = userInfo
	data["friendGroups"] = groups
	data["friendNotes"] = notes
	//
	token, err := jwt.GenToken(userInfo.UserID, userInfo.UserName)
	data["token"] = token
	//
	session := sessions.Default(c)
	session.Set("userID", userInfo.UserID)
	session.Set("isLogin", true)
	session.Options(sessions.Options{MaxAge: 60 * 60})
	session.Save()
	//
	c.JSON(http.StatusOK, constant.Response{
		Code:    constant.LoginSuccessCode,
		Message: constant.LoginSuccess,
		Data:    data,
	})
	return
}

func GetUserInfo(c *gin.Context) {
	var form models.UserInfoForm
	err := c.Bind(&form)
	if err != nil {
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: constant.GetParamsError,
			Data:    nil,
		})
		return
	}

	result, err := service.GetUserInfo(form.UserID)
	if err != nil {
		logrus.Info(err)
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: constant.ServerError,
			Data:    nil,
		})
		return
	}

	data := make(map[string]interface{})
	data["userInfo"] = result
	c.JSON(http.StatusOK, constant.Response{
		Code:    constant.GetMyFriendCode,
		Message: constant.ApiSuccess,
		Data:    data,
	})
	return
}

func PreFetchUser(c *gin.Context) {
	var user models.SearchUser
	err := c.Bind(&user)
	if err != nil {
		logrus.Info(err)
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: constant.GetParamsError,
			Data:    nil,
		})
		return
	}

	results, err := service.SearchUser(user)
	if err != nil {
		logrus.Info(err)
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: constant.ServerError,
			Data:    nil,
		})
		return
	}

	data := make(map[string]interface{})
	data["userList"] = results
	c.JSON(http.StatusOK, constant.Response{
		Code:    constant.GetMyFriendCode,
		Message: constant.ApiSuccess,
		Data:    data,
	})
	return
}

func UpdateUserPwd(c *gin.Context) {
	var user models.UpdateUserPwd
	err := c.Bind(&user)
	if err != nil {
		logrus.Info(err)
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: constant.GetParamsError,
			Data:    nil,
		})
		return
	}

	err = service.UpdatePwd(user)
	if err != nil {
		logrus.Info(err)
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, constant.Response{
		Code:    constant.GetMyFriendCode,
		Message: constant.UpdateSuccess,
		Data:    nil,
	})
}

func UpdateUserInfo(c *gin.Context) {
	var user models.UpdateUserInfo
	err := c.Bind(&user)
	if err != nil {
		logrus.Info(err)
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: constant.GetParamsError,
			Data:    nil,
		})
		return
	}

	err = service.UpdateUserInfo(user)
	if err != nil {
		logrus.Info(err)
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: constant.UpdateFailed,
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, constant.Response{
		Code:    constant.GetMyFriendCode,
		Message: constant.UpdateSuccess,
		Data:    nil,
	})
}

func UpdateUserConfigure(c *gin.Context) {
	var user models.UpdateUserConfigure
	err := c.Bind(&user)
	if err != nil {
		logrus.Info(err)
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: constant.GetParamsError,
			Data:    nil,
		})
		return
	}

	err = service.UpdateUserConfigure(user)
	if err != nil {
		logrus.Info(err)
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: constant.UpdateFailed,
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, constant.Response{
		Code:    constant.GetMyFriendCode,
		Message: constant.UpdateSuccess,
		Data:    nil,
	})
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

func IP(c *gin.Context) {
	err := service.BrowserSetting()
	if err != nil {
		logrus.Info(err)
	}
}
