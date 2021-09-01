package controllers

import (
	"blogo/app"
	"regexp"
	"strings"

	"github.com/revel/revel"
)

const (
	auth_ignore_paths  = "/admin/login,/admin/logout"
	auth_check_pattern = "^/admin/.*$"
)

type AppController struct {
	*revel.Controller
}

// Session 获取当前Session，没有则创建
func (c *AppController) Session() *app.CachedSession {
	session := app.GetCachedSession(c.Request)
	if session == nil {
		session = app.NewCachedSession()
		c.SetCookie(session.Cookie())
	}
	return session
}

// checkAuth 检查权限
func (c *AppController) checkAuth() revel.Result {
	if strings.Contains(auth_ignore_paths, c.Request.GetRequestURI()) {
		//忽略检查地址
		return nil
	}
	matched, _ := regexp.MatchString(auth_check_pattern, c.Request.GetRequestURI())
	if matched {
		loginName := c.Session().Get("loginName")
		if loginName == nil || loginName == "" {
			return c.Forbidden("没有权限访问")
		}
	}
	return nil
}
