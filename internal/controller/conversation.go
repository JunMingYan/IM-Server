package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"server/constant"
	"server/internal/models"
	"server/internal/service"
)

func ConversationList(c *gin.Context) {
	var form models.ConversationForm
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

	results, err := service.ConversationList(form)
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
	data["singleRecentConversationList"] = results
	c.JSON(http.StatusOK, constant.Response{
		Code:    constant.SuccessCode,
		Message: constant.ApiSuccess,
		Data:    data,
	})
	return
}
