package controllers

import (
	//	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func AdminIndex(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["foo"] = "bar"
	session.Values[42] = 43
	session.Values["data"] = "malong"
	session.Save(r, w)
	w.Write([]byte("Hello, admin!"))
}

func AdminLogin(w http.ResponseWriter, r *http.Request) {
	// Get a session.
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	value, ok := session.Values["data"].(string)
	if ok {
		//已经登录
		w.Write([]byte(value))
	} else {
		//未登录去登录页面
		w.Write([]byte("No Ok"))
	}
}
