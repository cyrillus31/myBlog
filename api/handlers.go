package api

func handleBozo(w http.ResponseWriter, r *http.Request) {
  number := r.PathValue("number")
  fmt.Fprintf(w, "What a bozo. %s", number)
  }

