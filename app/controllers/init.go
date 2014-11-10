package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"regexp"
	"time"
	"unicode/utf8"
)

func init() {

	revel.InterceptMethod((*AppController).checkAuth, revel.BEFORE)

	addTemplateFun()
}

//添加模板函数
func addTemplateFun() {

	revel.TemplateFuncs["delHtml"] = func(text string) string {
		if text != "" {
			reg, _ := regexp.Compile("<(.[^>]*)>")
			text = reg.ReplaceAllString(text, "")
			if utf8.RuneCountInString(text) > 200 {
				text = string([]rune(text)[0:200]) + "..."
			}
		}
		return text
	}

	revel.TemplateFuncs["randAvatar"] = func(n int64) string {
		return fmt.Sprintf("%d.jpg", n%3+1)
	}

	revel.TemplateFuncs["timeFormat"] = func(t time.Time, layout string) string {
		return t.Format(layout)
	}
}
