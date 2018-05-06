package mgodb

import (
	"gopkg.in/mgo.v2"
)

var (
	mgoSession *mgo.Session
	dbname     = "ggapidb"
)

// GetSession holds the database firing up
func GetSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial("localhost")
		if err != nil {
			panic(err)
		}
	}
	return mgoSession.Clone()
}
