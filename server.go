package main

import (
	"github.com/gorilla/mux"
	"github.com/widuu/goini"
	"mlonz/controllers"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.IndexHandler)
	r.HandleFunc("/show", controllers.ShowHandler)
	r.HandleFunc("/admin", controllers.AdminIndex)
	r.HandleFunc("/session", controllers.AdminLogin)
	//处理静态文件(开发中使用),感觉生产环境最好还是使用nginx直接rewrite到静态文件去处理
	//r.HandleFunc("/static/{category}/{file}", controllers.StaticHandler)
	r.HandleFunc("/static/{category}/{file}", controllers.ServeStatic)
	//获取conf
	conf := goini.SetConfig("conf/app.ini")
	port := conf.GetValue("dev", "port")
	http.ListenAndServe(":"+port, r)
}
