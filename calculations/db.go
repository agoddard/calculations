package calculations

import (
	"fmt"
	"labix.org/v2/mgo"
)

func GetSession() *mgo.Session {
	session, err := mgo.Dial(mongoUrl)
	if err != nil {
		fmt.Printf("Error connecting to mongo: %s", err)
	}
	return session
}

func GetCollection() *mgo.Collection {
	session := GetSession()
	return session.DB("").C(collection)
}
