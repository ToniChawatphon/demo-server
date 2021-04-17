package app

import (
	"log"
	"net/http"
)

// HTTPServer receives http requests
func HTTPServer(port string) {
	log.Printf("Start listening at port %v", port)
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("- receive request")
}
