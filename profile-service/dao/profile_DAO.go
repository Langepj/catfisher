package dao

import (
	"log"
  "profile-service/model"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProfileDAO struct {
	Server   string
	Database string
}

const (
	COLLECTION = "profiles"
)

func (m *ProfileDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *ProfileDAO) FindAll() ([]Profile, error) {
	var profiles []Profile
	err := db.C(COLLECTION).Find(bson.M{}).All(&profiles)
	return movies, err
}

func (m *ProfileDAO) FindById(id string) (Movie, error) {
	var profile Profile
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&profile)
	return movie, err
}

func (m *ProfileDAO) Insert(profile Profile) error {
	err := db.C(COLLECTION).Insert(&profile)
	return err
}

func (m *ProfileDAO) Delete(profile Profile)  error {
	err := db.C(COLLECTION).Remove(&profile)
	return err
}

func (m *ProfileDAO) Update(profile Profile) error {
	err := db.C(COLLECTION).UpdateId(profile.ID, &profile)
	return err
}
