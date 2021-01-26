package srvr

import (
	"fmt"
	"net/http"
)

func (s *Srvr) handleNewGame() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if s.Debug {
			fmt.Printf("Enter handleNewGame closure\n")
			defer fmt.Printf("Exit handleNewGame closure\n")
		}
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, newGameHTML)
	}
}

func (s *Srvr) handleMove() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if s.Debug {
			fmt.Printf("Enter handleMove closure\n")
			defer fmt.Printf("Exit handleMove closure\n")
		}
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, moveHTML)
	}
}

var newGameHTML = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <title>title</title>
    <!-- link rel="stylesheet" href="style.css">
    <script src="script.js"></script -->
  </head>
<body>
<head1>New Game</head1>
</body>
</html>
`

var moveHTML = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <title>title</title>
    <!-- link rel="stylesheet" href="style.css">
    <script src="script.js"></script -->
  </head>
<body>
<head1>Move</head1>
</body>
</html>
`
