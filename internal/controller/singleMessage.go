package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"server/constant"
	"server/internal/models"
	"server/internal/service"
)

func GetRecentSingleMessages(c *gin.Context) {
	var form models.RecentSingleMessages
	if err := c.ShouldBind(&form); err != nil {
		logrus.Info(err)
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: constant.GetParamsError,
			Data:    nil,
		})
		return
	}

	results, err := service.GetRecentSingleMessages(form)
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
	if results == nil {
		result := new([]models.SingleMessage)
		data["recentMessage"] = result
	} else {
		data["recentMessage"] = results
	}
	c.JSON(http.StatusOK, constant.Response{
		Code:    constant.GetMyFriendCode,
		Message: constant.ApiSuccess,
		Data:    data,
	})
	return
}

func LastFriendMsg(c *gin.Context) {
	var form models.LastFriendMsg
	if err := c.ShouldBind(&form); err != nil {
		logrus.Info(err)
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: constant.GetParamsError,
			Data:    nil,
		})
		return
	}

	result, err := service.LastFriendMsg(form)
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
	data["lastMsg"] = result
	c.JSON(http.StatusOK, constant.Response{
		Code:    constant.GetMyFriendCode,
		Message: constant.ApiSuccess,
		Data:    data,
	})
	return
}

func IsRead(c *gin.Context) {
	var form models.IsRead
	err := c.ShouldBind(&form)
	if err != nil {
		logrus.Info(err)
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: constant.GetParamsError,
			Data:    nil,
		})
		return
	}

	err = service.IsRead(form)
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
		Data:    nil,
	})
}
