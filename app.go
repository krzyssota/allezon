package main

import (
	"fmt"
	"net/http"
)

var counter int = 20

func debug(w http.ResponseWriter, req *http.Request) {

	if counter%10 == 0 {
		fmt.Println("server: req", req, "\nheader", req.Header, "\nbody", req.Body)
	}
	counter++
	//defer fmt.Println("server: hello handler ended")
	if _, err := fmt.Fprintf(w, "%s", req.Body); err != nil {
		fmt.Println("sending body error", err)
	}
}

func main() {
	fmt.Println("server: listening at \\")

	http.HandleFunc("/", debug)
	http.ListenAndServe(":8090", nil)
}
