package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Courses struct{
	Courseid string `json:"courseid"`
	Course string `json:"course"`
	Price int `json:"price"`
	Author *Authors `json:"author"`
}

type Authors struct{
	Name string `json:"name"`
	Website string `json:"website"`
}

var db []Courses

func main()  {
	fmt.Println("we come to crud operation in golang")
	r:=mux.NewRouter();
	r.HandleFunc("/",serve).Methods("GET")
	r.HandleFunc("/getallcourse",getAllcourses).Methods("GET")
	r.HandleFunc("/getacourse",getacourse).Methods("GET")
	r.HandleFunc("/createcourse",createacourse).Methods("POST")
	r.HandleFunc("/updateacourse",updateacourse).Methods("PUT")
	r.HandleFunc("/deleteacourse",deleteacourse).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":4000",r))
}

func serve(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("<h1>Hello</h1>"))
}

func (c *Courses)isEmpty()bool{
	return c.Course==""&&c.Courseid==""
}
func getAllcourses(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-type","application/json")
	json.NewEncoder(w).Encode(db)
}

func getacourse(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-type","application/json")

	prams:=mux.Vars(r)

	for _,co:=range(db){
		if co.Courseid==prams["id"]{
			json.NewEncoder(w).Encode(co)
			return
		}
	}
	json.NewEncoder(w).Encode("Wrong course ID please send the correct one !!!")
}

func createacourse(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-type","application/json")

	var course Courses

	json.NewDecoder(r.Body).Decode(&course)

	if course.isEmpty(){
		json.NewEncoder(w).Encode("Header is empty")
		return
	}
	rand.Seed(time.Now().Unix())
	course.Courseid=strconv.Itoa(rand.Intn(100))

	db=append(db, course)
	json.NewEncoder(w).Encode("New Course created !!")
}

func updateacourse(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-type","application/json")

	var course Courses

	_= json.NewDecoder(r.Body).Decode(&course)

	if course.isEmpty(){
		_= json.NewEncoder(w).Encode("there is no course details !!!")
		return
	}

	for i,c:=range(db){
		if c.Courseid==course.Courseid{
			db[i]=course
			_= json.NewEncoder(w).Encode("Course details have been updated")
			return
		}
	}
	json.NewEncoder(w).Encode("There no such courseid !!!")
}

func deleteacourse(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-type","application/json")

	prams:=mux.Vars(r)

	for i,co:=range(db){
		if co.Courseid==prams["id"]{
			db = append(db[:i], db[i+1:]...)
			json.NewEncoder(w).Encode("Just deleted a course...")
			return
		}
	}
	json.NewEncoder(w).Encode("Sorry course not found!!!")
}