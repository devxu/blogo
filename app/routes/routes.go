// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tAppController struct {}
var AppController tAppController



type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


type tAdmin struct {}
var Admin tAdmin


func (_ tAdmin) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Admin.Index", args).Url
}

func (_ tAdmin) ListPost(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Admin.ListPost", args).Url
}

func (_ tAdmin) QueryPosts(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Admin.QueryPosts", args).Url
}

func (_ tAdmin) CreatePost(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Admin.CreatePost", args).Url
}

func (_ tAdmin) EditPost(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Admin.EditPost", args).Url
}

func (_ tAdmin) DeletePost(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Admin.DeletePost", args).Url
}

func (_ tAdmin) SavePost(
		post interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "post", post)
	return revel.MainRouter.Reverse("Admin.SavePost", args).Url
}

func (_ tAdmin) ListComment(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Admin.ListComment", args).Url
}

func (_ tAdmin) QueryComments(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Admin.QueryComments", args).Url
}

func (_ tAdmin) DeleteComment(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Admin.DeleteComment", args).Url
}

func (_ tAdmin) Login(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Admin.Login", args).Url
}

func (_ tAdmin) LoginSubmit(
		username string,
		password string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "password", password)
	return revel.MainRouter.Reverse("Admin.LoginSubmit", args).Url
}

func (_ tAdmin) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Admin.Logout", args).Url
}


type tBlog struct {}
var Blog tBlog


func (_ tBlog) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Blog.Index", args).Url
}

func (_ tBlog) QueryPage(
		page int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "page", page)
	return revel.MainRouter.Reverse("Blog.QueryPage", args).Url
}

func (_ tBlog) ShowPost(
		slug string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "slug", slug)
	return revel.MainRouter.Reverse("Blog.ShowPost", args).Url
}

func (_ tBlog) AddComment(
		comment interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "comment", comment)
	return revel.MainRouter.Reverse("Blog.AddComment", args).Url
}

func (_ tBlog) Archives(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Blog.Archives", args).Url
}

func (_ tBlog) Tags(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Blog.Tags", args).Url
}

func (_ tBlog) Works(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Blog.Works", args).Url
}

func (_ tBlog) About(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Blog.About", args).Url
}


