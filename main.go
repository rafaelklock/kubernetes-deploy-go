package main

import (
	"fmt"
	"net/http"
	"log"
)

const webContent = "dev-ops-ninja:v1.0.1-beta"

func main() {
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, webContent)
}
