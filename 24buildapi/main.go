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

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

/* Database */
var courses []Course

/* Middleware and Helpers */
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("API Runnings")

	/* Seeding Database	*/
	courses = append(courses, Course{CourseId: "1", CourseName: "ReactJS", CoursePrice: 199, Author: &Author{
		Fullname: "DJ",
		Website:  "dj.com",
	}})

	courses = append(courses, Course{CourseId: "2", CourseName: "NodeJS", CoursePrice: 299, Author: &Author{
		Fullname: "DJ1",
		Website:  "dj1.com",
	}})

	/* Defining the Router */
	r := mux.NewRouter()

	/* Routing Methods */
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	/* Starting Server */
	log.Fatal(http.ListenAndServe(":3000", r))
}

/* Controllers */
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World from DJ</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all the courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting One Course")
	w.Header().Set("Content-Type", "application/json")

	/* Getting the Id from the User */
	params := mux.Vars(r)

	for _, val := range courses {
		if val.CourseId == params["id"] {
			json.NewEncoder(w).Encode(val)
			return
		}
	}

	message := make(map[string]string)
	message["status"] = "Sorry, The Course by that Id does not exists!"
	json.NewEncoder(w).Encode(message)
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please Give some body")
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	rand.Seed(time.Now().Unix())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update a course with Id")
	w.Header().Set("Content-Type", "application/json")

	var oneCourse Course
	_ = json.NewDecoder(r.Body).Decode(&oneCourse)

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			course = oneCourse
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No Course by that ID Exists")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete One Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode("Delete Operation Successful")
}
