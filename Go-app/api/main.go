package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Course struct{
	Courseid string `json:"courseid"`
	Title string `json:"title"`
	Price int `json:"price"`
	Author *Author `json:"author"`
}
type Author  struct{
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}
func main(){
	r:=mux.NewRouter()
	r.HandleFunc("/",servehome).Methods("GET")
	r.HandleFunc("/getAllcourse",getAllcourses).Methods("GET")
	r.HandleFunc("/getonecourse",getoneCourse).Methods("GET")
}

//database
var course []Course

//middleware
func (c *Course)isEmpty()bool{
	if c.Author!=nil{
		return false
	}
	if c.Courseid!=""{
		return false
	}
	if c.Title !="" {
		return false
	}
	return true
}

//controllers

func getAllcourses(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(course)
}

func getoneCourse(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-type","application/json")
	prams:=mux.Vars(r)

	for _,c:=range(course){
		if c.Courseid==prams["id"]{
			json.NewEncoder(w).Encode(c)
			return
		}
	}
	json.NewEncoder(w).Encode("Incorrect course id or course does not found")
}

func servehome(w http.ResponseWriter,r *http.Request) {
	w.Write([]byte("<h1>This is an api route !!!</h1>"))
}