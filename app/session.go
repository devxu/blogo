package app

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/revel/revel"
	"github.com/revel/revel/cache"
	"net/http"
	"strings"
	"sync"
	"time"
)

const (
	COOKIE_NAME_SESSIONID = "GO_SESSIONID"
	EXPIRE_AFTER_DURATION = 30 * time.Minute
)

//缓存Session，数据放在内存中
type CachedSession struct {
	sync.Mutex
	Id    string
	Items map[string]interface{}
}

/**
 * 放入数据
 */
func (s *CachedSession) Set(key string, value interface{}) {
	s.Lock()
	defer s.Unlock()
	s.Items[key] = value
}

/**
 * 获取数据
 */
func (s *CachedSession) Get(key string) interface{} {
	value, found := s.Items[key]
	if found {
		return value
	}
	return nil
}

/**
 * 删除数据
 */
func (s *CachedSession) Remove(key string) {
	s.Lock()
	defer s.Unlock()
	delete(s.Items, key)
}

/**
 * 使Session无效
 */
func (s *CachedSession) Invalidate() {
	s.Lock()
	defer s.Unlock()
	for k, _ := range s.Items {
		delete(s.Items, k)
	}
	cache.Delete(s.Id)
}

/**
 * Session的Cookie
 */
func (s *CachedSession) Cookie() *http.Cookie {
	return &http.Cookie{
		Name:     COOKIE_NAME_SESSIONID,
		Value:    s.Id,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
	}
}

func CachedSessionFilter(c *revel.Controller, fc []revel.Filter) {
	fc[0](c, fc[1:])

	session := GetCachedSession(c.Request.Request)
	if session != nil {
		cache.Replace(session.Id, *session, EXPIRE_AFTER_DURATION)
		c.RenderArgs["session"] = session
	}

}

func GetCachedSession(req *http.Request) *CachedSession {
	cookie, _ := req.Cookie(COOKIE_NAME_SESSIONID)
	if cookie != nil {
		session := &CachedSession{Items: map[string]interface{}{}}
		err := cache.Get(cookie.Value, session)
		if err == nil {
			return session
		}
	}
	return nil
}

func NewCachedSession() *CachedSession {
	sid := "SID-" + strings.ToUpper(strings.Replace(uuid.NewUUID().String(), "-", "", -1))
	session := CachedSession{Id: sid, Items: map[string]interface{}{}}
	cache.Add(session.Id, session, EXPIRE_AFTER_DURATION)
	return &session
}
