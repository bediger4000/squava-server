package srvr

import "fmt"

// Routes sets the URIs and functions that work them.
// "pattern" string documented here: https://golang.org/pkg/net/http/#ServeMux
func (s *Srvr) Routes() {
	if s.Debug {
		fmt.Printf("Enter Srvr.Routes closure\n")
		defer fmt.Printf("Exit Srvr.Routes closure\n")

	}
	s.Router.HandleFunc("/move", s.handleMove())
	s.Router.HandleFunc("/newgame", s.handleNewGame())
}
