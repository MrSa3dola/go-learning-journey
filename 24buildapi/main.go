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

// Model for course 0 file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake DB
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("APIs")
	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299,
		Author: &Author{FullName: "Saad Mohamed", Website: "saad.dev"},
	}, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 199,
		Author: &Author{FullName: "Saad Mohamed", Website: "go.dev"},
	})
	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by Saad</h1>"))
}
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab the id from the request
	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with that id")
	return
}
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")
	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}
	// data sent
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course) // no need to take the return, we passed the course as a refrece &
	// every item will be inside it
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside")
		return
	}
	for _, curCourse := range courses {
		if curCourse.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Can't insert with the same course name")
			return
		}
	}
	// generate unitque id, string
	// append course into courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// grab the id
	params := mux.Vars(r)
	// loop, id, remove, add with my ID
	for idx, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:idx], courses[idx+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// send response when id is not found
	json.NewEncoder(w).Encode("Id not found")
	return
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// loop, id, remove
	for idx, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:idx], courses[idx+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode("Course deleted successfully")
}
