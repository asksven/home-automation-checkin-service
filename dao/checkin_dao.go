package dao

import (
	"log"

	"github.com/golang/glog"

	"github.com/asksven/home-automation-checkin-service/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// CheckInDAO contains the database connection
type CheckInDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "checkins"
)

// Connect Establish a Connection to database
func (m *CheckInDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// FindAll returns the list of all checkins
func (m *CheckInDAO) FindAll() ([]models.CheckIn, error) {
	var checkins []models.CheckIn
	err := db.C(COLLECTION).Find(bson.M{}).All(&checkins)
	return checkins, err
}

// DeleteAll deletes all checkins
func (m *CheckInDAO) DeleteAll() () {
	db.C(COLLECTION).RemoveAll(bson.M{})
}

// FindAllByLocation returns the list of checkins for one location
func (m *CheckInDAO) FindAllByLocation(location string) ([]models.CheckIn, error) {
	var checkins []models.CheckIn
	err := db.C(COLLECTION).Find(bson.M{"location": location}).All(&checkins)
	return checkins, err
}

// FindByName returns one checkin for the given name (or none)
func (m *CheckInDAO) FindByName(name string) (models.CheckIn, error) {
	var checkin models.CheckIn
	err := db.C(COLLECTION).Find(bson.M{"name": name}).One(&checkin)
	return checkin, err
}

// Insert adds a checkin into database
func (m *CheckInDAO) Insert(checkin models.CheckIn) error {
	err := db.C(COLLECTION).Insert(&checkin)
	return err
}

// Delete deletes an existing checkin for a given name from the database
func (m *CheckInDAO) Delete(name string) error {
	var result models.CheckIn
	err := db.C(COLLECTION).Find(bson.M{"name": name}).One(&result)
	if err != nil {
		glog.Error("could not name object by name:" + name + ". No deletion")
		return err
	}
	err = db.C(COLLECTION).Remove(&result)
	return err
}
