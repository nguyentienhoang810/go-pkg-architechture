package app

import (
	"net/http"
	"fmt"
)

func (sv *Service) SetRouters() {
	fmt.Println("set routers from root")
	sv.Httpsv.MuxRouter().HandleFunc("/", sv.GetIndex).Methods("GET")
	sv.Httpsv.MuxRouter().HandleFunc("/home", sv.GetIndex).Methods("GET")
	sv.Httpsv.MuxRouter().HandleFunc("/login", sv.GetLogin).Methods("GET")
	sv.Httpsv.MuxRouter().HandleFunc("/login", sv.PostLogin).Methods("POST")
	sv.Httpsv.MuxRouter().HandleFunc("/register", sv.GetRegister).Methods("GET")
	sv.Httpsv.MuxRouter().HandleFunc("/register", sv.PostRegister).Methods("POST")
}

func (sv *Service) GetIndex(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "index, get all products")
}

func (sv *Service) GetLogin(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "get login page")
}

func (sv *Service) PostLogin(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "excute login")
}

func (sv *Service) GetRegister(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "get register page")
}

func (sv *Service) PostRegister(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "excute register")
}