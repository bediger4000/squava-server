package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"squava-server/srvr"
)

func main() {
	portString := flag.String("p", "8072", "TCP port on which to listen")
	debug := flag.Bool("v", false, "verbose output per request")
	flag.Parse()

	srv := &srvr.Srvr{
		Router: http.NewServeMux(),
		Debug:  *debug,
	}

	srv.Routes()

	s := &http.Server{
		Addr:           ":" + *portString,
		Handler:        srv.Router,
		ReadTimeout:    40 * time.Second,
		WriteTimeout:   40 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("Squava solver listening on TCP port %s\n", *portString)

	log.Fatal(s.ListenAndServe())

}

// readStopWords has the name of a file as its formal argument,
// returns a map of strings to bool, if a string appears,
// don't put it in the dictionary.
func readStopWords(stopWordsFileName string) map[string]bool {

	if stopWordsFileName == "" {
		return make(map[string]bool)
	}

	buffer, err := ioutil.ReadFile(stopWordsFileName)
	if err != nil {
		log.Fatalf("Problem reading stop words file %q: %v\n", stopWordsFileName, err)
	}

	lines := bytes.Split(buffer, []byte{'\n'})

	removes := make(map[string]bool)
	for i := range lines {
		if len(lines[i]) == 0 {
			continue
		}
		removes[string(lines[i])] = true
	}

	return removes
}
