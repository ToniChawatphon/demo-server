package app

import (
	"encoding/json"
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
	message := HTTPMessgae{"OK", 200}
	json, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
