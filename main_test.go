package main

import (
	"os"
	"testing"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func Test_StoreAndFind(t *testing.T) {
	session, err := mgo.Dial(os.Getenv("WERCKER_MONGODB_HOST"))
	if err != nil {
		panic(err)
	}
	defer session.Close()

	conn := session.DB("test").C("decepticons")
	err = conn.Insert(&Decepticon{"Shockwave"}, &Decepticon{"Starscream"})
	if err != nil {
		t.Error("Could not insert a Decepticon")
	}

	result := Decepticon{}
	err = conn.Find(bson.M{"name": "Shockwave"}).One(&result)
	if err != nil {
		t.Error("Could not find a Decepticon")
	}
	t.Log("Test Passed")
}
