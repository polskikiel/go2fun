package main

import (
	"net/http"
	"io/ioutil"
	"log"
	"strings"
	"app/model"
	"fmt"
	"encoding/json"
)

type myHandler struct {
}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("template/base.html")
	if err != nil {
		w.Write([]byte(http.StatusText(404)))
		log.Fatal(err)
		return
	}

	w.Write(data)
}

func main() {
	http.Handle("/", new(myHandler))
	http.HandleFunc("/user", UserHandler)
	http.ListenAndServe(":9090", nil)
}
func UserHandler(w http.ResponseWriter, r *http.Request) {

	path := strings.Split(
		r.URL.Path[len("/user"):], "/")

	var user model.User
	var username string

	i := len(path)

	if i > 1 {
		username = string([]byte(path[1])[:len(path[1])-1])
	}

	fmt.Println(i)

	switch i {
	case 3:
		user = model.CreateUser(
			username,
			path[i-1])
		break
	case 2:
		user = model.CreateUser(
			username,
			"")
		break
	default:
		user = model.CreateUser("anon", "")
	}

	js, err := json.Marshal(user)
	if err != nil{
		log.Fatal("User cannot be created")
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(js)
}
