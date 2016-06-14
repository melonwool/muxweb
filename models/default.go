package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/widuu/goini"
	"log"
)

func conn() (*sql.DB, error) {
	conf := goini.SetConfig("conf/app.ini")
	user := conf.GetValue("dev", "database.user")
	password := conf.GetValue("dev", "database.password")
	host := conf.GetValue("dev", "database.host")
	port := conf.GetValue("dev", "database.port")
	dbname := conf.GetValue("dev", "database.dbname")

	db, err := sql.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+dbname)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	//defer db.Close()
	return db, nil
}

func Getdata() (map[int]map[string]string, error) {
	var (
		id   string
		name string
	)
	datas := make(map[int]map[string]string)
	//db, err := sql.Open("mysql", "test:NfUA0gYnp0Php0ELVqxFNx2RU8nAYcyw@tcp(127.0.0.1:3306)/test")
	db, err := conn()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer db.Close()
	//rows, err := db.Query("select id, name from info order by id DESC LIMIT 5")
	rows, err := db.Query("select id, name from info LIMIT 5")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	i := 1
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		data := make(map[string]string)
		data["id"] = id
		data["name"] = name
		datas[i] = data
		//fmt.Println(datas)
		i = i + 1
	}
	return datas, nil
}
