package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func must(err error) {
	if err != nil {
		log.Println("internal error:", err)
		panic(internalError)
	}
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	must(json.NewEncoder(w).Encode(v))
}
