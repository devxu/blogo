package models

import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
	// _ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/revel/revel"
	"github.com/robfig/config"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

var (
	Engine   *xorm.Engine
	MyConfig *config.Config
)

func init() {
	revel.OnAppStart(func() {
		syncDB()
		addTemplateFun()
		insertTestData()
	})
}

//获取数据库连接，同步数据库表结构
func syncDB() {
	var err error
	MyConfig, _ = config.ReadDefault(revel.BasePath + "/conf/my.conf")
	dbDriver, _ := MyConfig.String("db", "db.driver")
	dbUrl, _ := MyConfig.String("db", "db.url")
	Engine, err = xorm.NewEngine(dbDriver, dbUrl)
	if err != nil {
		panic(err)
	}

	err = Engine.Sync(
		new(Post),
		new(Comment),
	)

	if err != nil {
		panic(err)
	}
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

//添加测试内容
func insertTestData() {

	var affects int64
	firstPost := &Post{}
	count, _ := Engine.Count(firstPost)
	if count == 0 {
		firstPost.Slug = strings.Replace(uuid.NewUUID().String(), "-", "", -1)
		firstPost.Title = "Hello world!"
		firstPost.Content = "第一篇测试内容，<strong>Hello world!</strong>"
		firstPost.Tags = "测试"
		firstPost.CommentCount = 0
		firstPost.Created = time.Now()
		affects, _ = Engine.InsertOne(firstPost)

		if affects > 0 && firstPost.Id > 0 {
			firstComment := &Comment{}
			firstComment.PostId = firstPost.Id
			firstComment.Name = "System"
			firstComment.Message = "第一个评论测试"
			firstComment.Created = time.Now()
			affects, _ = Engine.InsertOne(firstComment)
			fmt.Println("insert first comment affects = ", affects)
			if affects > 0 {
				firstPost.CommentCount = 1
				Engine.Update(firstPost)
			}
		}

	}

}
