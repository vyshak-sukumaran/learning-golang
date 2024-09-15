package main

import (
	"fmt"
	"net/http"
)

type DB interface {
	Store(string) error
}

type Store struct{}

func (s *Store) Store(value string) error {
	fmt.Println("storing into db", value)
	return nil
}

func makeHTTPFunc(db DB, fn httpFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(db, w, r); err != nil {
			//
		}
	}
}

func handler(db DB, w http.ResponseWriter, r *http.Request) error {
	return nil
}

type httpFunc func(db DB, w http.ResponseWriter, r *http.Request) error
