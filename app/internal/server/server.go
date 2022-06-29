package server

import (
	"database/sql"
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"
)

type Server struct {
	DB  *sql.DB
	Log *logrus.Logger
}

func (s *Server) handleFunc(w http.ResponseWriter, r *http.Request) {
	var err error

	db := s.DB
	log := s.Log

	q, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		log.Info()
	}

	name := q.Get("n")

	queryADD := "insert into names(name_user) values($1)"

	_, err = db.Exec(queryADD, name)
	if err != nil {
		log.Error(err)
	}

	_, err = w.Write([]byte(name))
	if err != nil {
		log.Error(err)
	}
}

func (s *Server) Run() error {
	var err error

	mux := http.NewServeMux()
	mux.HandleFunc("/logger/", s.handleFunc)

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		return err
	}

	return err
}
