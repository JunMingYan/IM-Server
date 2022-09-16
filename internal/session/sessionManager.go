package session

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type SessionProvider interface {
	create(sessionID string, data map[string]interface{}) error
	get(sessionID, key string) (string, error)
	getAll(sessionID string) (map[string]string, error)
	set(sessionID, key string, value interface{}) error
	destroy(sessionID string) error
	gc(expire int64) error
}

type SessionManager struct {
	cookieName    string
	cookieExpire  int
	sessionExpire int64
	gcDuration    int
	provider      SessionProvider
}

func NewManager(cookieName string, cookieExpire int, sessionExpire int64, gcDuration int, provider SessionProvider) *SessionManager {
	return &SessionManager{
		cookieName,
		cookieExpire,
		sessionExpire,
		gcDuration,
		provider,
	}
}

func (sm *SessionManager) CreateSessionID(req *http.Request) string {
	addr := req.RemoteAddr
	userAgent := req.Header.Get("User-Agent")
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10000)
	str := addr + "_" + userAgent + "_" + strconv.Itoa(n)
	h := md5.New()
	h.Write([]byte(str))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func (sm *SessionManager) GetSessionID(req *http.Request) (string, error) {
	c, err := req.Cookie(sm.cookieName)
	if err != nil {
		return "", errors.New("Reading cookie failed :" + err.Error())
	}
	if len(c.Value) == 0 {
		return "", errors.New("Cookie does not exists: " + sm.cookieName)
	}
	return c.Value, err
}

func (sm *SessionManager) Create(c *gin.Context, data map[string]interface{}) error {
	sessionID, _ := sm.GetSessionID(c.Request)
	if len(sessionID) > 0 {
		data, _ := sm.provider.getAll(sessionID)
		if data != nil {
			return nil
		}
	}
	sessionID = sm.CreateSessionID(c.Request)
	if len(sessionID) == 0 {
		return errors.New("length of sessionID is 0")
	}
	err := sm.provider.create(sessionID, data)
	if err != nil {
		return err
	}
	if sm.cookieExpire == 0 {
		http.SetCookie(c.Writer, &http.Cookie{
			Name:     sm.cookieName,
			Value:    sessionID,
			Path:     "/",
			HttpOnly: true,
		})
	} else {
		expire, _ := time.ParseDuration(strconv.Itoa(sm.cookieExpire) + "m")
		http.SetCookie(c.Writer, &http.Cookie{
			Name:     sm.cookieName,
			Value:    sessionID,
			Path:     "/",
			Expires:  time.Now().Add(expire),
			HttpOnly: true,
		})
	}
	return nil
}

func (sm *SessionManager) Get(req *http.Request, key string) (string, error) {
	sessionID, _ := sm.GetSessionID(req)
	if len(sessionID) == 0 {
		return "", errors.New("length of session ID is 0")
	}
	return sm.provider.get(sessionID, key)
}

func (sm *SessionManager) GetAll(c *gin.Context) (map[string]string, error) {
	sessionID, _ := sm.GetSessionID(c.Request)
	if len(sessionID) == 0 {
		return nil, errors.New("length of session ID is 0")
	}
	return sm.GetAll(c)
}

func (sm *SessionManager) Set(c *gin.Context, key string, value interface{}) error {
	sessionID, _ := sm.GetSessionID(c.Request)
	if len(sessionID) == 0 {
		return errors.New("length of session ID is 0")
	}
	return sm.provider.set(sessionID, key, value)
}

func (sm *SessionManager) Destroy(req *http.Request) error {
	sessionID, _ := sm.GetSessionID(req)
	if len(sessionID) == 0 {
		return errors.New("length of session ID is 0")
	}
	return sm.provider.destroy(sessionID)
}

func (sm *SessionManager) Gc() error {
	err := sm.provider.gc(sm.sessionExpire)
	duration, _ := time.ParseDuration(strconv.Itoa(sm.gcDuration) + "m")
	time.AfterFunc(duration, func() {
		sm.Gc()
	})
	return err
}
