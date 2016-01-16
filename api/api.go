package api

import (
	"log"
	"os"

	"github.com/go-zoo/trash"
	"gopkg.in/mgo.v2"
)

var (
	session = &mgo.Session{}
	db      = &mgo.Database{}
	errLog  = trash.New(log.New(os.Stdout, "<!SPOON!> ", 0), "json")
)

func init() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		errLog.NewErr("NO DB CONNECTION", err).Log()
		return
	}
	db = session.DB("spoon")
}
