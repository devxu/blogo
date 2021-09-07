package controllers

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

const (
	auth_ignore_paths  = "/admin/login,/admin/logout"
	auth_check_pattern = "^/admin/.*$"
	flash_error        = "flash.error"
)

var (
	funcMap = template.FuncMap{
		"delHtml": func(text string) string {
			if text != "" {
				reg, _ := regexp.Compile("<(.[^>]*)>")
				text = reg.ReplaceAllString(text, "")
				if utf8.RuneCountInString(text) > 200 {
					text = string([]rune(text)[0:200]) + "..."
				}
			}
			return text
		},
		"randAvatar": func(n int64) string {
			return fmt.Sprintf("%d.jpg", n%3+1)
		},
		"timeFormat": func(t time.Time, layout string) string {
			return t.Format(layout)
		},
		"set": func(viewArgs map[string]interface{}, key string, value interface{}) template.JS {
			viewArgs[key] = value
			return template.JS("")
		},
		"append": func(viewArgs map[string]interface{}, key string, value interface{}) template.JS {
			if viewArgs[key] == nil {
				viewArgs[key] = []interface{}{value}
			} else {
				viewArgs[key] = append(viewArgs[key].([]interface{}), value)
			}
			return template.JS("")
		},
		"date": func(date time.Time) string {
			return date.Format("2006-01-02")
		},
		"datetime": func(date time.Time) string {
			return date.Format("2006-01-02 15:04:05")
		},
	}
)

type TemplateRenderer struct {
	templates *template.Template
}

func NewTemplateRenderer() *TemplateRenderer {

	renderer := &TemplateRenderer{}
	renderer.templates = template.Must(template.New("renderer").Funcs(funcMap).ParseGlob("views/**/*"))
	//log.Println(renderer.templates.DefinedTemplates())
	return renderer
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if data == nil {
		data = make(echo.Map)
	}
	if dataMap, ok := data.(echo.Map); ok {
		dataMap["session"] = getSession(c).Values
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

func getSession(c echo.Context) *sessions.Session {
	sess, _ := session.Get("session", c)
	return sess
}

func CheckAuth() echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			if strings.Contains(auth_ignore_paths, c.Request().RequestURI) {
				return next(c)
			}

			matched, _ := regexp.MatchString(auth_check_pattern, c.Request().RequestURI)
			if matched {
				sess := getSession(c)
				if sess.Values["loginName"] == nil {
					return echo.ErrForbidden
				}
			}
			return next(c)
		}
	}
}
