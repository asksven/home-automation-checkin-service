package dao

import (
	"testing"
	"fmt"
	"time"

	"github.com/asksven/home-automation-checkin-service/models"
	"gopkg.in/mgo.v2/bson"
)

var data = CheckInDAO{}

func init() {
	fmt.Println("Init")
	data.Server = "localhost"
	data.Database = "movies_db"
	data.Connect()
}


func TestDeleteAll(t *testing.T) {
	t.Log("TestDeleteAll")

	data.DeleteAll()

}

func TestInsert(t *testing.T) {
	t.Log("TestInsert")

	var checkin = models.CheckIn{}
	checkin.Name = "paul"
	checkin.Location = "home"
	checkin.ID = bson.NewObjectId()
	checkin.Stamp = fmt.Sprintf("%s", time.Now().UTC().Format(time.RFC1123))

	// we want to be sure that we find what we expect so we first delete everything
	data.DeleteAll()

	err := data.Insert(checkin)
	if err != nil {
		fmt.Print("Error: ")
		fmt.Println(err)
		t.Errorf("An error occured")
	}

	checkins, err := data.FindAll()
	t.Log(checkins)

	res, err := data.FindByName(checkin.Name)
	if err != nil {
		fmt.Print("Error: ")
		fmt.Println(err)
		t.Errorf("An error occured")

	} else {		
		if res.Name != "paul" {
			t.Errorf("res should not be null")
		} else {
			t.Log("Success: " + res.Name)
		}
	}
}

func TestFindByName(t *testing.T) {
	t.Log("TestFindByName")

	var checkin = models.CheckIn{}
	checkin.Name = "paul"
	checkin.Location = "home"
	checkin.ID = bson.NewObjectId()
	checkin.Stamp = fmt.Sprintf("%s", time.Now().UTC().Format(time.RFC1123))

	// we want to be sure that we find what we expect so we first delete everything
	data.DeleteAll()

	err := data.Insert(checkin)
	if err != nil {
		fmt.Print("Error: ")
		fmt.Println(err)
		t.Errorf("An error occured")
	}

	res, err := data.FindByName(checkin.Name)
	if err != nil {
		fmt.Print("Error: ")
		fmt.Println(err)
		t.Errorf("An error occured")

	} else {		
		if res.Name != "paul" {
			t.Errorf("res should not be null")
		} else {
			t.Log("Success: " + res.Name)
		}
	}

}

func TestFindAll(t *testing.T) {
	t.Log("TestFindAll")

	var checkin = models.CheckIn{}
	checkin.Name = "paul"
	checkin.Location = "home"
	checkin.ID = bson.NewObjectId()
	checkin.Stamp = fmt.Sprintf("%s", time.Now().UTC().Format(time.RFC1123))

	var checkin2 = models.CheckIn{}
	checkin2.Name = "mary"
	checkin2.Location = "home"
	checkin2.ID = bson.NewObjectId()
	checkin2.Stamp = fmt.Sprintf("%s", time.Now().UTC().Format(time.RFC1123))


	// we want to be sure that we find what we expect so we first delete everything
	data.DeleteAll()

	err := data.Insert(checkin)
	if err != nil {
		fmt.Print("Error: ")
		fmt.Println(err)
		t.Errorf("An error occured")
	}

	err2 := data.Insert(checkin2)
	if err2 != nil {
		fmt.Print("Error: ")
		fmt.Println(err2)
		t.Errorf("An error occured")
	}

	checkins, err := data.FindAll()

	if err != nil {
		fmt.Print("Error: ")
		fmt.Println(err)
		t.Errorf("An error occured")

	} else {		
		if len(checkins) != 2 {
			t.Errorf("there should be two entries in the database")
		} 
	}
}

func TestFindAllByLocation(t *testing.T) {
	t.Log("TestFindAllByLocation")

	var checkin = models.CheckIn{}
	checkin.Name = "paul"
	checkin.Location = "home"
	checkin.ID = bson.NewObjectId()
	checkin.Stamp = fmt.Sprintf("%s", time.Now().UTC().Format(time.RFC1123))

	var checkin2 = models.CheckIn{}
	checkin2.Name = "mary"
	checkin2.Location = "home2"
	checkin2.ID = bson.NewObjectId()
	checkin2.Stamp = fmt.Sprintf("%s", time.Now().UTC().Format(time.RFC1123))

	// we want to be sure that we find what we expect so we first delete everything
	data.DeleteAll()

	err := data.Insert(checkin)
	if err != nil {
		fmt.Print("Error: ")
		fmt.Println(err)
		t.Errorf("An error occured")
	}

	err2 := data.Insert(checkin2)
	if err2 != nil {
		fmt.Print("Error: ")
		fmt.Println(err2)
		t.Errorf("An error occured")
	}

	checkins, err := data.FindAllByLocation("home")

	if err != nil {
		fmt.Print("Error: ")
		fmt.Println(err)
		t.Errorf("An error occured")

	} else {		
		if len(checkins) != 1 {
			t.Errorf("there should be one entry in the database at home")
		} 
	}
}

func TestDelete(t *testing.T) {
	t.Log("TestDelete")

	var checkin = models.CheckIn{}
	checkin.Name = "paul"
	checkin.Location = "home"
	checkin.ID = bson.NewObjectId()
	checkin.Stamp = fmt.Sprintf("%s", time.Now().UTC().Format(time.RFC1123))

	var checkin2 = models.CheckIn{}
	checkin2.Name = "mary"
	checkin2.Location = "home"
	checkin2.ID = bson.NewObjectId()
	checkin2.Stamp = fmt.Sprintf("%s", time.Now().UTC().Format(time.RFC1123))

	// we want to be sure that we find what we expect so we first delete everything
	data.DeleteAll()

	err := data.Insert(checkin)
	if err != nil {
		fmt.Print("Error: ")
		fmt.Println(err)
		t.Errorf("An error occured")
	}

	err = data.Insert(checkin2)
	if err != nil {
		fmt.Print("Error: ")
		fmt.Println(err)
		t.Errorf("An error occured")
	}

	checkins, err := data.FindAll()

	fmt.Println(checkins)
	
	if err != nil {
		fmt.Print("Error: ")
		fmt.Println(err)
		t.Errorf("An error occured")

	} else {		
		if len(checkins) != 2 {
			t.Errorf("there should be two entries in the database")
		} 
	}

	err = data.Delete("paul")
	if err != nil {
		fmt.Print("Error: ")
		fmt.Println(err)
		t.Errorf("An error occured")

	}

	checkins, err = data.FindAll()

	if err != nil {
		fmt.Print("Error: ")
		fmt.Println(err)
		t.Errorf("An error occured")

	} else {		
		if len(checkins) != 1 {
			t.Errorf("there should be one entry in the database")
		} 
	}

}
