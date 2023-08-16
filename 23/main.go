package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for course
type Course struct {
	CourseId string `json:"id"`
	CourseName string `json:"courseName"`
	Price int `json:"price"`
	Author *Author `json:"author"`
}

type Author struct {
	Name string `json:"fullname"`
	Domain string `json:"website"`
}

// dummy db
var courses []Course

// middlewares
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func (a *Author) IsNil() bool {
	return a.Name == ""
}

func main() {
	// create route
	r:=mux.NewRouter()
	// add data
	courses=append(courses, Course{
		CourseId: "3",
		CourseName: "GoLang",
		Price: 30,
		Author: &Author{
			Name: "Devil",
			Domain: "www.google.com",
		},
	})
	courses=append(courses, Course{
		CourseId: "5",
		CourseName: "React-Native",
		Price: 50,
		Author: &Author{
			Name: "Devil",
			Domain: "www.youtube.com",
		},
	})
	// routes
	r.HandleFunc("/",serveHome).Methods("GET")
	r.HandleFunc("/api/v1/courses",getCourses).Methods("GET")
	r.HandleFunc("/api/v1/course/{id}",getCourseById).Methods("GET")
	r.HandleFunc("/api/v1/course/add",createCourse).Methods("POST")
	r.HandleFunc("/api/v1/course/update/{id}",updateCourse).Methods("PUT")
	r.HandleFunc("/api/v1/course/delete/{id}",deleteCourse).Methods("DELETE")
	// listen port 
	log.Fatal(http.ListenAndServe(":4000",r))
}

// controllers 
// home route 
// always response first then request 
func serveHome(w http.ResponseWriter ,r *http.Request){
	w.Write([]byte("<h1>Homee</h1>"))
}

// get courses
func getCourses(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	// send the json 
	json.NewEncoder(w).Encode(courses)
}

func getCourseById(w http.ResponseWriter , r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	//get id from request 
	params:=mux.Vars(r)
	// loop through it 
	// finding the matching course with id and return it 
	for _, i := range courses{
		if i.CourseId == params["id"]{
			json.NewEncoder(w).Encode(i)
			return
		}
	}
	message:=map[string]string{"error":"No course found with given id","id": params["id"]}
	json.NewEncoder(w).Encode(message)
}

func createCourse(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type","application/json")

	// empty body validation
	if r.Body == nil{
		// basically as we do in node express 
		// res.json({message:"Your message"})
		json.NewEncoder(w).Encode("Can't Send Empty Data")
	}

	// handling other cases
	var course Course
	_ =json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty(){
		json.NewEncoder(w).Encode("Cannot Send Empty Data")
		return
	}

	// checking if the course name exists 
	for _, res := range courses {
		if res.CourseName == course.CourseName{
			json.NewEncoder(w).Encode("Course already exists ")
			return
		}
	}

	// generate unique id and then int -> string convo 
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses=append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateCourse(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type","application/json")
	// taking the ids 
	params := mux.Vars(r)
	var updatedCourse Course
	found:=false
	// iteration
	for index,course:=range courses {
		if course.CourseId == params["id"]{
			found =true
			// decode json
			if err:= json.NewDecoder(r.Body).Decode(&updatedCourse); err!=nil{
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			// update course
			updatedCourse.CourseId = params["id"]
			courses[index] = updatedCourse
			json.NewEncoder(w).Encode(updatedCourse)
			return
		}
	}
	if !found {
		message := map[string]string{"error": "No course found with given id", "id": params["id"]}
		json.NewEncoder(w).Encode(message)
	}
}

func deleteCourse(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type","application/json")
	// extract the id 
	params:=mux.Vars(r)
	found:=false
	// loop , & if id found delete the course
	for index,course:= range courses {
		if course.CourseId == params["id"]{
			courses= append(courses[:index],courses[index+1:]... )
			found=true
			json.NewEncoder(w).Encode("Deleted the course")
			break
		}
	}
	if !found {
		message := map[string]string{"error": "No course found with given id", "id": params["id"]}
		json.NewEncoder(w).Encode(message)
	}
}

func ErrorHandler (err error ){
	if err != nil {
		panic(err)
	}
}