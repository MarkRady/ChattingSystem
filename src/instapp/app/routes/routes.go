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


type tApplicationController struct {}
var ApplicationController tApplicationController


func (_ tApplicationController) Home(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ApplicationController.Home", args).URL
}

func (_ tApplicationController) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ApplicationController.Index", args).URL
}

func (_ tApplicationController) Create(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ApplicationController.Create", args).URL
}

func (_ tApplicationController) Show(
		token string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "token", token)
	return revel.MainRouter.Reverse("ApplicationController.Show", args).URL
}

func (_ tApplicationController) Update(
		token string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "token", token)
	return revel.MainRouter.Reverse("ApplicationController.Update", args).URL
}

func (_ tApplicationController) Delete(
		token string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "token", token)
	return revel.MainRouter.Reverse("ApplicationController.Delete", args).URL
}


type tChatController struct {}
var ChatController tChatController


func (_ tChatController) Index(
		token string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "token", token)
	return revel.MainRouter.Reverse("ChatController.Index", args).URL
}

func (_ tChatController) Create(
		token string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "token", token)
	return revel.MainRouter.Reverse("ChatController.Create", args).URL
}

func (_ tChatController) Show(
		token string,
		number int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "token", token)
	revel.Unbind(args, "number", number)
	return revel.MainRouter.Reverse("ChatController.Show", args).URL
}

func (_ tChatController) Delete(
		token string,
		number int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "token", token)
	revel.Unbind(args, "number", number)
	return revel.MainRouter.Reverse("ChatController.Delete", args).URL
}


type tMessageController struct {}
var MessageController tMessageController


func (_ tMessageController) Index(
		token string,
		chat_number int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "token", token)
	revel.Unbind(args, "chat_number", chat_number)
	return revel.MainRouter.Reverse("MessageController.Index", args).URL
}

func (_ tMessageController) Create(
		token string,
		chat_number int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "token", token)
	revel.Unbind(args, "chat_number", chat_number)
	return revel.MainRouter.Reverse("MessageController.Create", args).URL
}

func (_ tMessageController) Search(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("MessageController.Search", args).URL
}

func (_ tMessageController) Show(
		token string,
		chat_number int64,
		message_number int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "token", token)
	revel.Unbind(args, "chat_number", chat_number)
	revel.Unbind(args, "message_number", message_number)
	return revel.MainRouter.Reverse("MessageController.Show", args).URL
}

func (_ tMessageController) Update(
		token string,
		chat_number int64,
		message_number int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "token", token)
	revel.Unbind(args, "chat_number", chat_number)
	revel.Unbind(args, "message_number", message_number)
	return revel.MainRouter.Reverse("MessageController.Update", args).URL
}

func (_ tMessageController) Delete(
		token string,
		chat_number int64,
		message_number int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "token", token)
	revel.Unbind(args, "chat_number", chat_number)
	revel.Unbind(args, "message_number", message_number)
	return revel.MainRouter.Reverse("MessageController.Delete", args).URL
}


