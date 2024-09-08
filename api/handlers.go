package api

import (
	"fmt"
	"net/http"
)

func HandleBozo(w http.ResponseWriter, r *http.Request) {
	number := r.PathValue("number")
	fmt.Fprintf(w, "What a superhero. %s", number)
}
