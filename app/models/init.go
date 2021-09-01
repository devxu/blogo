package models

import (
	"github.com/google/uuid"
	// _ "github.com/go-sql-driver/mysql"
	"strings"
	"time"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/revel/config"
	"github.com/revel/revel"
)

var (
	Engine   *xorm.Engine
	MyConfig *config.Config
)

func init() {
	revel.OnAppStart(func() {
		syncDB()
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

//添加测试内容
func insertTestData() {

	var affects int64
	firstPost := &Post{}
	count, _ := Engine.Count(firstPost)
	if count == 0 {
		firstPost.Slug = strings.Replace(uuid.NewString(), "-", "", -1)
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
			revel.AppLog.Info("insert first comment affects = ", affects)
			if affects > 0 {
				firstPost.CommentCount = 1
				Engine.Update(firstPost)
			}
		}

	}

}
