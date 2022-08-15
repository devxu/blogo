package main

import (
	"blogo/controllers"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
)

func main() {

	session_key := os.Getenv("session_key")
	log.Println("session_key:", session_key)
	if len(session_key) == 0 {
		log.Fatalln("Please set up session_key in environment variables before running")
	}

	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(session_key))))
	e.Use(controllers.CheckAuth())
	//e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Renderer = controllers.NewTemplateRenderer()

	// Routes
	e.Static("/public", "public")
	e.File("/favicon.ico", "public/img/favicon.png")

	e.GET("/", controllers.Index)
	e.GET("/page/:page", controllers.QueryPage)
	e.GET("/post/:slug", controllers.ShowPost)
	e.POST("/addComment", controllers.AddComment)
	e.GET("/archives", controllers.Archives)
	e.GET("/tags", controllers.Tags)
	e.GET("/works", controllers.Works)
	e.GET("/about", controllers.About)
	e.GET("/admin/login", controllers.Login)
	e.POST("/admin/login", controllers.LoginSubmit)
	e.GET("/admin/logout", controllers.Logout)
	e.GET("/admin/home", controllers.Home)
	e.GET("/admin/posts", controllers.ListPost)
	e.POST("/admin/queryPosts", controllers.QueryPosts)
	e.GET("/admin/post/create", controllers.CreatePost)
	e.GET("/admin/post/edit/:id", controllers.EditPost)
	e.GET("/admin/post/delete/:id", controllers.DeletePost)
	e.POST("/admin/post/save", controllers.SavePost)
	e.GET("/admin/comments", controllers.ListComment)
	e.POST("/admin/queryComments", controllers.QueryComments)
	e.GET("/admin/comment/delete/:id", controllers.DeleteComment)

	// Start server
	e.Logger.Fatal(e.Start(":8081"))
}
