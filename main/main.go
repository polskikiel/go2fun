package main

import (
	"net/http"
	"io/ioutil"
	"log"
	"strings"
	"app/model"
	"fmt"
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

	if len(path) > 1 {
		username = string([]byte(path[1])[:len(path[1])-1])
	}

	fmt.Println(len(path))

	switch len(path) {
	case 3:
		user = model.CreateUser(
			username,
			path[len(path)-1])
		break
	case 2:
		user = model.CreateUser(
			username,
			"")
		break
	default:
		user = model.CreateUser("anon", "")
	}

	w.Write([]byte(user.String()))

}
