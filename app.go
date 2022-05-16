package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var counter int = 20

func debug(w http.ResponseWriter, req *http.Request) {

	/*if counter%10 == 0 {
		fmt.Println("server: req", req, "\nheader", req.Header, "\nbody", req.Body)
	}
	counter++*/
	//if _, err := fmt.Fprintf(w, "%s", req.Body); err != nil {
	defer func(w io.Writer, format string, a ...any) {
		if _, err := fmt.Fprintf(w, format, a); err != nil {
			fmt.Println("sending 'hej' error", err)
		}
	}(w, "hej")
}

func main() {
	fmt.Println("server: listening at \\")

	http.HandleFunc("/", debug)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
