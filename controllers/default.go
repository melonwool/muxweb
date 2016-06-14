package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"log"
	"mlonz/models"
	"net/http"
	"os"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var pc map[string]string
	pc = make(map[string]string)
	query := r.URL.Query()
	if query["act"] != nil {
		pc["act"] = query["act"][0]
	}
	pc["qingdao"] = "青岛"
	pc["jinan"] = "济南"
	pc["yantai"] = "烟台"
	jsonStr := json.NewEncoder(w)
	jsonStr.Encode(pc)
}

func ShowHandler(w http.ResponseWriter, r *http.Request) {
	type Output struct {
		Title   string
		Details map[int]map[string]string
	}
	datas, err := models.Getdata()
	if err != nil {
		log.Fatal(err)
	}
	data := &Output{Title: "数据库列表", Details: datas}
	t, err := template.ParseFiles("template/idname.html")
	err = t.Execute(w, data)
}

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	buf, err := ioutil.ReadFile("static/" + string(vars["category"]) + "/" + string(vars["file"]))
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	}
	w.Write([]byte(buf))
}
