// GENERATED CODE - DO NOT EDIT
// This file provides a way of creating URL's based on all the actions
// found in all the controllers.
package routes

import "github.com/revel/revel"


type tJobs struct {}
var Jobs tJobs


func (_ tJobs) Status(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Jobs.Status", args).URL
}


type tController struct {}
var Controller tController


func (_ tController) Begin(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Controller.Begin", args).URL
}

func (_ tController) Rollback(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Controller.Rollback", args).URL
}

func (_ tController) Commit(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Controller.Commit", args).URL
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeDir(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeDir", args).URL
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
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}

func (_ tStatic) ServeModuleDir(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModuleDir", args).URL
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).URL
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).URL
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).URL
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).URL
}


type tApplication struct {}
var Application tApplication


func (_ tApplication) AddUser(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Application.AddUser", args).URL
}

func (_ tApplication) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Application.Index", args).URL
}

func (_ tApplication) Register(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Application.Register", args).URL
}

func (_ tApplication) SaveUser(
		user interface{},
		verifyPassword string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	revel.Unbind(args, "verifyPassword", verifyPassword)
	return revel.MainRouter.Reverse("Application.SaveUser", args).URL
}

func (_ tApplication) Login(
		username string,
		password string,
		remember bool,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "password", password)
	revel.Unbind(args, "remember", remember)
	return revel.MainRouter.Reverse("Application.Login", args).URL
}

func (_ tApplication) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Application.Logout", args).URL
}


type tHotels struct {}
var Hotels tHotels


func (_ tHotels) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Hotels.Index", args).URL
}

func (_ tHotels) ListJson(
		search string,
		size uint64,
		page uint64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "search", search)
	revel.Unbind(args, "size", size)
	revel.Unbind(args, "page", page)
	return revel.MainRouter.Reverse("Hotels.ListJson", args).URL
}

func (_ tHotels) List(
		search string,
		size uint64,
		page uint64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "search", search)
	revel.Unbind(args, "size", size)
	revel.Unbind(args, "page", page)
	return revel.MainRouter.Reverse("Hotels.List", args).URL
}

func (_ tHotels) Show(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Hotels.Show", args).URL
}

func (_ tHotels) Settings(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Hotels.Settings", args).URL
}

func (_ tHotels) SaveSettings(
		password string,
		verifyPassword string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "password", password)
	revel.Unbind(args, "verifyPassword", verifyPassword)
	return revel.MainRouter.Reverse("Hotels.SaveSettings", args).URL
}

func (_ tHotels) ConfirmBooking(
		id int,
		booking interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "booking", booking)
	return revel.MainRouter.Reverse("Hotels.ConfirmBooking", args).URL
}

func (_ tHotels) CancelBooking(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Hotels.CancelBooking", args).URL
}

func (_ tHotels) Book(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Hotels.Book", args).URL
}


