package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fraczles/glog/shared/db"
	"github.com/fraczles/glog/shared/models"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	//TODO: write JSONAPI to response
	w.Write([]byte("Hey!\n"))
}

func main() {
	// initialize DB
	dbmap := db.InitDb()
	defer dbmap.Db.Close()

	// delete existing rows
	fmt.Println("3")
	// create fake data
	p1 := models.NewPost("a", "a")
	p2 := models.NewPost("b", "b")

	err := dbmap.Insert(&p1, &p2)
	checkErr(err, "Cannot insert")

	r := mux.NewRouter()
	r.HandleFunc("/about", About)

	http.ListenAndServe(":8001", r)
}
