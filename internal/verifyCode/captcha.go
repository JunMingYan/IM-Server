package verifyCode

import (
	"bytes"
	"encoding/base64"
	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	Key = "verifyCode"
)

const (
	Length = captcha.DefaultLen
	Width  = 120
	Height = 40
)

type VerifyCode struct {
	ID   string
	Code string
}

func GetVerifyCode(c *gin.Context) (code string, err error) {
	id := captcha.NewLen(Length)
	var content bytes.Buffer
	err = captcha.WriteImage(&content, id, Width, Height)
	if err != nil {
		return "", err
	}
	//
	session := sessions.Default(c)
	session.Set(Key, id)
	session.Options(sessions.Options{MaxAge: 5 * 60})
	_ = session.Save()
	//
	code = base64.StdEncoding.EncodeToString(content.Bytes())
	return
}

func Verify(c *gin.Context, code string) bool {
	session := sessions.Default(c)
	captchaID := session.Get(Key)
	if captchaID != nil {
		session.Delete(Key)
		_ = session.Save()
		return captcha.VerifyString(captchaID.(string), code)
	} else {
		return false
	}
}
