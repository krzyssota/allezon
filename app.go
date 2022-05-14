package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Println("server: req", req)
	body, err := req.GetBody()

	if err == nil {
		fmt.Fprintf(w, "%s", body)
	} else {
		fmt.Println("getBody Err", w, "err", err)

	}

	/*ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(4 * time.Second):
		fmt.Fprintf(w, "%s", req.Context())
	case <-ctx.Done():

		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}*/
}

func main() {
	fmt.Println("server: listening at \\")

	http.HandleFunc("/", hello)
	http.ListenAndServe(":8090", nil)
}
