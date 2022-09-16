package routers

import (
	"github.com/gin-gonic/gin"
	controller "server/internal/controller"
	"server/internal/session"
	"server/internal/websocket"
	"server/pkg/jwt"
)

func GetRouter() *gin.Engine {
	// if mode == gin.ReleaseMode {
	// 	gin.SetMode(gin.ReleaseMode) // 设置成发布模式
	// }
	r := gin.Default()
	r.Use(session.SetSession("GOSESSIONID"))
	//r.StaticFS("/api/static", http.Dir("./resources"))
	r.Static("/api/static", "./resources")
	//
	socket := websocket.RunWebSocket
	//
	v1 := r.Group("/")
	//
	user := v1.Group("user")
	{
		user.GET("/getCode", controller.GetCaptcha)
		user.POST("/register", controller.RegisterHandler)
		user.POST("/login", controller.LoginHandler)
		user.GET("/logout", controller.Logout)
		user.GET("/ip", controller.IP)

		user.GET("/getUserInfo", jwt.JWTAuthMiddleware(), controller.GetUserInfo)
		user.POST("/preFetchUser", jwt.JWTAuthMiddleware(), controller.PreFetchUser)
		user.POST("/updateUserPwd", jwt.JWTAuthMiddleware(), controller.UpdateUserPwd)
		user.POST("/updateUserInfo", jwt.JWTAuthMiddleware(), controller.UpdateUserInfo)
		user.POST("/updateUserConfigure", jwt.JWTAuthMiddleware(), controller.UpdateUserConfigure)
		user.POST("/modifyFriendBeiZhu", jwt.JWTAuthMiddleware(), controller.UpdateFriendNote)
		user.POST("/modifyFriendFenZu", jwt.JWTAuthMiddleware(), controller.MoveFriendGroup)
		user.POST("/addFenZu", jwt.JWTAuthMiddleware(), controller.AddFriendGroup)
		user.POST("/delFenZu", jwt.JWTAuthMiddleware(), controller.DeleteFriendGroup)
		user.POST("/editFenZu", jwt.JWTAuthMiddleware(), controller.ModifyFriendGroupName)
	}

	systemUser := v1.Group("sys", jwt.JWTAuthMiddleware())
	{
		systemUser.GET("/getSysUsers", controller.SystemUsers)
		systemUser.POST("/filterMessage", controller.FilterMessage)
		systemUser.POST("/addFeedBack", controller.AddFeedBack)
	}

	group := v1.Group("group", jwt.JWTAuthMiddleware())
	{
		group.GET("/getMyGroupList", controller.MyGroupList)
		group.POST("/recentGroup", controller.RecentGroupList)
		group.GET("/getGroupInfo", controller.GroupInfo)
	}

	groupMessage := v1.Group("groupMessage", jwt.JWTAuthMiddleware())
	{
		groupMessage.POST("/getRecentGroupMessages", controller.GetRecentGroupMessages)
	}

	goodFriend := v1.Group("goodFriend", jwt.JWTAuthMiddleware())
	{
		goodFriend.POST("/recentConversationList", controller.ConversationList)
		goodFriend.GET("/getMyFriendsList", controller.MyFriendList)
	}

	validate := v1.Group("validate", jwt.JWTAuthMiddleware())
	{
		validate.POST("/getValidateMessage", controller.GetValidateMessage)
		validate.POST("/sendValidateMessage", controller.SendValidateMessage)
		validate.GET("/getMyValidateMessageList", controller.GetMyValidateMessageList)
		validate.POST("/assentRequest", controller.AssentRequest)
		validate.POST("/rejectRequest", controller.RejectRequest)
	}
	singleMessage := v1.Group("singleMessage", jwt.JWTAuthMiddleware())
	{
		singleMessage.POST("/getRecentSingleMessages", controller.GetRecentSingleMessages)
		singleMessage.GET("/getLastMessage", controller.LastFriendMsg)
		singleMessage.POST("/isRead", controller.IsRead)
	}

	ws := r.Group("ws")
	{
		ws.GET("/socket.io", socket)
	}
	//
	return r
}
