package jwt

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/constant"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const key = "Bearer"

type MyClaims struct {
	userID   uint64
	nickName string `json:"nickName"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour // 过期时间

var MySecret = []byte("JULIA")

func GenToken(userID uint64, nickName string) (string, error) {
	c := MyClaims{
		userID,
		nickName,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "Books",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(MySecret)
}

func ParseToken(tokenStr string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("无效的token")
}

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, constant.Response{
				Code:    constant.FailCode,
				Message: constant.RequestHeaderNull,
				Data:    nil,
			})
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == key) {
			c.JSON(http.StatusOK, constant.Response{
				Code:    constant.FailCode,
				Message: constant.RequestHeaderError,
				Data:    nil,
			})
			c.Abort()
			return
		}
		mc, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, constant.Response{
				Code:    constant.FailCode,
				Message: constant.TokenError,
				Data:    nil,
			})
			c.Abort()
			return
		}
		c.Set("userID", mc.userID)
		c.Set("userName", mc.nickName)
		c.Next()
	}
}
