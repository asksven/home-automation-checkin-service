package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"fmt"

	"github.com/golang/glog"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"

	"github.com/asksven/home-automation-checkin-service/config"
	"github.com/asksven/home-automation-checkin-service/dao"
	"github.com/asksven/home-automation-checkin-service/models"
)

var configuration = config.Config{}
var data = dao.CheckInDAO{}

// allCheckInsEndPoint GETs list of checkins
func allCheckInsEndPoint(w http.ResponseWriter, r *http.Request) {
	checkins, err := data.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, checkins)
}

// deleteAllCheckInsEndPoint DELETEs all checkins
func deleteAllCheckInsEndPoint(w http.ResponseWriter, r *http.Request) {
	data.DeleteAll()
	respondWithJSON(w, http.StatusOK, "{}")
}


// findCheckInByNameEndpoint GETs a checkin by its name
func findCheckInByNameEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	checkin, err := data.FindByName(params["name"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid object name")
		return
	}
	respondWithJSON(w, http.StatusOK, checkin)
}

// findCheckInByLocationEndpoint GETs all checkins by location
func findCheckInByLocationEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	checkins, err := data.FindAllByLocation(params["location"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid object name")
		return
	}
	respondWithJSON(w, http.StatusOK, checkins)
}

// createCheckInEndPoint POSTs a new checkin
func createCheckInEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var checkin models.CheckIn
	if err := json.NewDecoder(r.Body).Decode(&checkin); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// check if a checkin for that name exists and if yes, delete it
	existingcheckin, err := data.FindByName(checkin.Name)
	if err == nil {
		glog.Info("A checkin already exists for name: " + checkin.Name + ". We need to delete it first.")
		err = data.Delete(existingcheckin.Name)
		if err != nil {
			glog.Error("An error occured attempting to delete name: " + checkin.Name)
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	checkin.ID = bson.NewObjectId()
	now := time.Now().UTC()
	str := fmt.Sprintf("%s", now.Format(time.RFC1123))
//	broken, delete does not work properly if indroducing that attribute: checkin.Stamp = str
	glog.Info("Checked in at : " + str + ": "+ checkin.Name)


	if err := data.Insert(checkin); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, checkin)
}

// deleteCheckInEndPoint DELETEs an existing checkin
func deleteCheckInEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var checkin models.CheckIn
	if err := json.NewDecoder(r.Body).Decode(&checkin); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := data.Delete(checkin.Name); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	configuration.Read()

	data.Server = configuration.Server
	data.Database = configuration.Database
	data.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/checkins", allCheckInsEndPoint).Methods("GET")
	r.HandleFunc("/checkins", deleteAllCheckInsEndPoint).Methods("DELETE")
	r.HandleFunc("/checkins/{location}", findCheckInByLocationEndpoint).Methods("GET")
	r.HandleFunc("/checkin", createCheckInEndPoint).Methods("POST")
	r.HandleFunc("/checkin", deleteCheckInEndPoint).Methods("DELETE")
	r.HandleFunc("/checkin/{name}", findCheckInByNameEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
