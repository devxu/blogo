package controllers

import (
	"fmt"
	"regexp"
	"time"
	"unicode/utf8"

	"github.com/revel/revel"
)

func init() {

	revel.InterceptMethod((*AppController).checkAuth, revel.BEFORE)

	initTemplateFun()
}

// Initialize template functions
func initTemplateFun() {

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
