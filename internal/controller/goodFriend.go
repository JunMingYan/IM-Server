package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"server/constant"
	"server/internal/models"
	"server/internal/service"
)

func MyFriendList(c *gin.Context) {
	var form models.MyFriendForm
	err := c.Bind(&form)
	if err != nil {
		logrus.Info(err)
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: constant.GetParamsError,
			Data:    nil,
		})
		return
	}

	friends, err := service.MyFriendList(form.UserID)
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
	data["myFriendsList"] = friends
	c.JSON(http.StatusOK, constant.Response{
		Code:    constant.GetMyFriendCode,
		Message: constant.ApiSuccess,
		Data:    data,
	})
	return
}

func UpdateFriendNote(c *gin.Context) {
	var form models.UpdateFriendNote
	err := c.Bind(&form)
	if err != nil {
		logrus.Info(err)
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: constant.GetParamsError,
			Data:    nil,
		})
		return
	}

	results, err := service.UpdateFriendNote(form)
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
	data["friendNotes"] = results
	c.JSON(http.StatusOK, constant.Response{
		Code:    constant.GetMyFriendCode,
		Message: constant.ApiSuccess,
		Data:    data,
	})
	return
}

func MoveFriendGroup(c *gin.Context) {
	var form models.UpdateFriendGroup
	err := c.Bind(&form)
	if err != nil {
		logrus.Info(err)
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: constant.GetParamsError,
			Data:    nil,
		})
		return
	}

	results, err := service.MoveFriendGroup(form)
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
	data["friendGroups"] = results
	c.JSON(http.StatusOK, constant.Response{
		Code:    constant.GetMyFriendCode,
		Message: constant.ApiSuccess,
		Data:    data,
	})
	return
}

func AddFriendGroup(c *gin.Context) {
	var form models.AddFriendGroup
	err := c.Bind(&form)
	if err != nil {
		logrus.Info(err)
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: constant.GetParamsError,
			Data:    nil,
		})
		return
	}

	results, err := service.SaveFriendGroup(form)
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
	data["friendGroups"] = results
	c.JSON(http.StatusOK, constant.Response{
		Code:    constant.GetMyFriendCode,
		Message: constant.ApiSuccess,
		Data:    data,
	})
	return
}

func DeleteFriendGroup(c *gin.Context) {
	var form models.DeleteFriendGroup
	err := c.Bind(&form)
	if err != nil {
		logrus.Info(err)
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: constant.GetParamsError,
			Data:    nil,
		})
		return
	}

	results, err := service.DeleteFriendGroup(form)
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
	data["friendGroups"] = results
	c.JSON(http.StatusOK, constant.Response{
		Code:    constant.GetMyFriendCode,
		Message: constant.ApiSuccess,
		Data:    data,
	})
	return
}

func ModifyFriendGroupName(c *gin.Context) {
	var form models.UpdateFriendGroupName
	err := c.Bind(&form)
	if err != nil {
		logrus.Info(err)
		c.JSON(http.StatusOK, constant.Response{
			Code:    constant.FailCode,
			Message: constant.GetParamsError,
			Data:    nil,
		})
		return
	}

	results, err := service.ModifyFriendGroupName(form)
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
	data["friendGroups"] = results
	c.JSON(http.StatusOK, constant.Response{
		Code:    constant.GetMyFriendCode,
		Message: constant.ApiSuccess,
		Data:    data,
	})
	return
}
