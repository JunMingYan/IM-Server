package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//func getSession() sessions.Store {
//	sessionMaxAge := 60 * 60
//	sessionSecret := "JULIA"
//	store := cookie.NewStore([]byte(sessionSecret))
//	store.Options(sessions.Options{
//		MaxAge: sessionMaxAge,
//		Path:   "/",
//	})
//	return store
//}

func getSession() sessions.Store {
	sessionMaxAge := 60 * 60
	sessionSecret := "JULIA"
	store, err := redis.NewStore(10, "tcp", "localhost:6379", "123456", []byte(sessionSecret))
	if err != nil {
		logrus.Error(err)
		return nil
	}
	store.Options(sessions.Options{
		MaxAge: sessionMaxAge,
		Path:   "/",
	})
	return store
}

func SetSession(KeyParis string) gin.HandlerFunc {
	store := getSession()
	return sessions.Sessions(KeyParis, store)
}
