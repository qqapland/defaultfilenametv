package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("App Starting...")

	fmt.Println("YT_API_KEY:", os.Getenv("YT_API_KEY"))

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3000", nil)) // 3000 here
}
func handler(w http.ResponseWriter, r *http.Request) {
	simpleOutput := fmt.Sprintf("Got hit from: %s", r.URL.Path[1:])
	fmt.Println(simpleOutput)
	fmt.Fprintf(w, simpleOutput)
}
