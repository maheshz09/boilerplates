package main

// add at top of file
import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

var (
	coursesMu sync.RWMutex
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// getAllCourses
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	coursesMu.RLock()
	defer coursesMu.RUnlock()
	if err := json.NewEncoder(w).Encode(courses); err != nil {
		http.Error(w, "failed to encode courses", http.StatusInternalServerError)
	}
}

// getOneCourse
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	coursesMu.RLock()
	defer coursesMu.RUnlock()
	for _, course := range courses {
		if course.CourseId == id {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	http.Error(w, fmt.Sprintf("no course with id %s", id), http.StatusNotFound)
}

// createOneCourse
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		http.Error(w, "empty body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var course Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		http.Error(w, "invalid json: "+err.Error(), http.StatusBadRequest)
		return
	}
	if course.IsEmpty() {
		http.Error(w, "missing course name", http.StatusBadRequest)
		return
	}

	// generate ID - avoid collisions in small pools
	course.CourseId = strconv.FormatInt(time.Now().UnixNano(), 36) // better than rand.Intn
	coursesMu.Lock()
	courses = append(courses, course)
	coursesMu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(course)
}

// updateOneCourse
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	if id == "" {
		http.Error(w, "id required", http.StatusBadRequest)
		return
	}
	if r.Body == nil {
		http.Error(w, "empty body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var updated Course
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "invalid json: "+err.Error(), http.StatusBadRequest)
		return
	}
	if updated.IsEmpty() {
		http.Error(w, "missing fields", http.StatusBadRequest)
		return
	}

	coursesMu.Lock()
	defer coursesMu.Unlock()
	for i, c := range courses {
		if c.CourseId == id {
			updated.CourseId = id
			courses[i] = updated // replace in-place
			json.NewEncoder(w).Encode(updated)
			return
		}
	}
	http.Error(w, "course not found", http.StatusNotFound)
}

// deleteOneCourse
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	if id == "" {
		http.Error(w, "id required", http.StatusBadRequest)
		return
	}

	coursesMu.Lock()
	defer coursesMu.Unlock()
	for i, c := range courses {
		if c.CourseId == id {
			courses = append(courses[:i], courses[i+1:]...)
			w.WriteHeader(http.StatusNoContent) // 204 - no body
			return
		}
	}
	http.Error(w, "course not found", http.StatusNotFound)
}
