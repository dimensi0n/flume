package fs

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var Token string

func GetToken(t string) {
	Token = t
}

func generateName() string {
	var name string
	for i := 0; i < 30; i++ {
		name += strconv.Itoa(rand.Intn(9))
	}
	return name
}

func post(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("token") == Token {
		file, handler, err := r.FormFile("file")
		defer file.Close()
		if err != nil {
			panic(err)
		}

		os.MkdirAll("./public/", 0777)
		rand.Seed(time.Now().UnixNano())
		filename := generateName() + filepath.Ext(handler.Filename)
		f, err := os.OpenFile("./public/"+filename, os.O_WRONLY|os.O_CREATE, 0666)
		defer f.Close()
		io.Copy(f, file)
		json.NewEncoder(w).Encode(filename)
	} else {
		json.NewEncoder(w).Encode("Wrong Authentification Token")
	}
}

func FsHandler(r *mux.Router) {
	r.HandleFunc("/fs/post", post).Methods("POST")
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))
}
