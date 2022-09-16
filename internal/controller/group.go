package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"server/constant"
	"server/internal/models"
	"server/internal/service"
)

func MyGroupList(c *gin.Context) {
	var group models.MyGroupForm
	err := c.Bind(&group)
	if err != nil {
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: constant.GetParamsError,
			Data:    nil,
		})
		return
	}

	results, err := service.MyGroupList(group.UserName)
	if err != nil {
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: constant.ServerError,
			Data:    nil,
		})
		return
	}

	data := make(map[string]interface{})
	data["myGroupList"] = results
	c.JSON(http.StatusOK, constant.Response{
		Code:    constant.GetMyFriendCode,
		Message: constant.ApiSuccess,
		Data:    data,
	})
	return

}

func RecentGroupList(c *gin.Context) {
	var form models.RecentGroupForm
	err := c.Bind(&form)
	if err != nil {
		logrus.Fatal(err)
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: constant.GetParamsError,
			Data:    nil,
		})
		return
	}

	results, err := service.RecentGroupList(form)
	if err != nil {
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: constant.ServerError,
			Data:    nil,
		})
		return
	}

	data := make(map[string]interface{})
	data["recentGroups"] = results
	c.JSON(http.StatusOK, constant.Response{
		Code:    constant.GetMyFriendCode,
		Message: constant.ApiSuccess,
		Data:    data,
	})
	return

}

func GetRecentGroupMessages(c *gin.Context) {
	var form models.RecentGroupMessage
	err := c.Bind(&form)
	if err != nil {
		logrus.Fatal(err)
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: constant.GetParamsError,
			Data:    nil,
		})
		return
	}

	results, err := service.RecentGroupMessage(form)
	if err != nil {
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: constant.ServerError,
			Data:    nil,
		})
		return
	}

	data := make(map[string]interface{})
	data["recentGroupMessages"] = results
	c.JSON(http.StatusOK, constant.Response{
		Code:    constant.GetMyFriendCode,
		Message: constant.ApiSuccess,
		Data:    data,
	})
	return
}

func GroupInfo(c *gin.Context) {
	var form models.GroupInfoForm
	err := c.Bind(&form)
	if err != nil {
		logrus.Fatal(err)
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: constant.GetParamsError,
			Data:    nil,
		})
		return
	}

	result, results, err := service.GroupInfo(form)
	if err != nil {
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: constant.ServerError,
			Data:    nil,
		})
		return
	}

	data := make(map[string]interface{})
	data["groupInfo"] = result
	data["groupUserList"] = results
	c.JSON(http.StatusOK, constant.Response{
		Code:    constant.GetMyFriendCode,
		Message: constant.ApiSuccess,
		Data:    data,
	})
	return

}
