package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"os"
	"time"
)

// define data structure
type Decepticon struct {
	Name string
	Date time.Time
}

func config(k string) string {
	v := os.Getenv(k)
	if v == "" {
		return "localhost"
	}
	return v
}

func main() {

	session, err := mgo.Dial(config("WERCKER_MONGODB_HOST"))
	if err != nil {
		panic(err)
	}
	defer session.Close()

	conn := session.DB("test").C("decepticons")
	err = conn.Insert(&Decepticon{"Shockwave", time.Now()}, &Decepticon{"Starscream", time.Now()})

	if err != nil {
		panic(err)
	}

	result := Decepticon{}
	err = conn.Find(bson.M{"name": "Shockwave"}).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decepticon: ", result.Name)
}
