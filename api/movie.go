package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-zoo/bone"
	"gopkg.in/mgo.v2/bson"
)

type Movie struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	Title    string        `json:"title" bson:"title"`
	Runtime  int           `json:"runtime" bson:"runtime"`
	Released time.Time     `json:"released" bson:"released"`
	Genre    string        `json:"genre" bson:"genre"`
	Actors   []string      `json:"actors" bson:"actors"`
	Producer string        `json:"producer" bson:"producer"`
	Imdb     string        `json:"imdb" bson:"imdb"`
	Poster   string        `json:"poster" bson:"poster"`
	Rated    string        `json:"rated" bson:"rated"`
	Rating   int           `json:"rating" bson:"rating"`
}

func GetMovie(rw http.ResponseWriter, req *http.Request) {
	id := bone.GetValue(req, "id")
	movie := &Movie{}
	err := db.C("movie").FindId(bson.ObjectIdHex(id)).One(movie)
	if err != nil {
		errLog.NewHTTPErr("NOT FOUND", "This movie doesn't exist").LogHTTP(req).SendHTTP(rw, 404)
		return
	}
	err = json.NewEncoder(rw).Encode(movie)
	if err != nil {
		errLog.NewHTTPErr("NOT FOUND", err).LogHTTP(req).SendHTTP(rw, 404)
		return
	}
}

func GetMovieByTitle(rw http.ResponseWriter, req *http.Request) {
	title := bone.GetValue(req, "title")
	movie := &Movie{}
	err := db.C("movie").Find(bson.M{"title": title}).One(movie)
	if err != nil {
		data := GetMovieData(title)
		err = json.NewEncoder(rw).Encode(&data)
		if err != nil {
			errLog.NewHTTPErr("NOT FOUND", "This movie doesn't exist").LogHTTP(req).SendHTTP(rw, 404)
			return
		}
		return
	}
	err = json.NewEncoder(rw).Encode(movie)
	if err != nil {
		errLog.NewHTTPErr("NOT FOUND", err).LogHTTP(req).SendHTTP(rw, 404)
		return
	}
}

func PostMovie(rw http.ResponseWriter, req *http.Request) {
	movie := &Movie{}
	err := json.NewDecoder(req.Body).Decode(movie)
	if err != nil {
		errLog.NewHTTPErr("INVALID JSON DATA", err).LogHTTP(req).SendHTTP(rw, 406)
		return
	}

	movie.ID = bson.NewObjectId()
	json.NewEncoder(rw).Encode(map[string]int{"status": 200})
	err = db.C("movie").Insert(movie)
	if err != nil {
		errLog.NewHTTPErr("DB ERROR", err).LogHTTP(req).SendHTTP(rw, 501)
		return
	}
}
