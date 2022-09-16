package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"server/constant"
	"server/internal/models"
	"server/internal/mysql"
	"server/internal/service"
)

func SystemUsers(c *gin.Context) {
	list, err := service.SystemUserList()
	if err != nil {
		logrus.Info(err)
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: constant.ServerError,
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, constant.Response{
		Code:    constant.GetMyFriendCode,
		Message: constant.ApiSuccess,
		Data:    list,
	})
	return
}

func FilterMessage(c *gin.Context) {
	var msg models.FilterMessage
	err := c.Bind(&msg)
	if err != nil {
		logrus.Info(err)
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: constant.GetParamsError,
			Data:    nil,
		})
		return
	}

	data := make(map[string]interface{})
	data["message"] = msg.Message
	c.JSON(http.StatusOK, constant.Response{
		Code:    constant.GetMyFriendCode,
		Message: constant.ApiSuccess,
		Data:    data,
	})
	return
}

func AddFeedBack(c *gin.Context) {
	var back models.FeedBack
	err := c.Bind(&back)
	if err != nil {
		logrus.Info(err)
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: constant.GetParamsError,
			Data:    nil,
		})
		return
	}

	err = mysql.SaveFeeBack(back)
	if err != nil {
		logrus.Info(err)
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: constant.GetParamsError,
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, constant.Response{
		Code:    constant.GetMyFriendCode,
		Message: constant.ApiSuccess,
		Data:    nil,
	})
	return
}
