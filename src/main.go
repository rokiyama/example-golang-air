package main

import (
	"io"
	"log"
	"net/http"
)

const port = "8080"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Println(req)
		w.Header().Set("Content-Type", "application/json")
		if _, err := io.WriteString(w, `{"result":"ok"}`); err != nil {
			log.Print(err)
		}
	})
	log.Printf("running on port:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
