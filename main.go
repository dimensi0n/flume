package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/dimensi0n/flume/auth"
	"github.com/dimensi0n/flume/db"
	"github.com/dimensi0n/flume/fs"
	"github.com/gorilla/mux"
)

func main() {
	wordPtr := flag.String("t", "1234", "")
	flag.Parse()
	fmt.Println("token: ", *wordPtr)
	auth.GetToken(*wordPtr)
	db.GetToken(*wordPtr)
	fs.GetToken(*wordPtr)
	r := mux.NewRouter()
	db.DbHandler(r)
	auth.AuthHandler(r)
	fs.FsHandler(r)

	fmt.Println("Listening on port :8080")
	http.ListenAndServe(":8080", r)
}
