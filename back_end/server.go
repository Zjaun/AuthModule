package back_end

import (
	"encoding/json"
	"net/http"
)

func Request(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "POST requests only.", http.StatusMethodNotAllowed)
		return
	}

	var name struct {
		First string `json:"first"`
		Last  string `json:"last"`
	}

	err := json.NewDecoder(r.Body).Decode(&name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := ValidFirstName(name.First); err != nil {
		e := json.NewEncoder(w).Encode(ValidationResponse{Valid: false, Reason: err.Error()})
		if e != nil {
			http.Error(w, e.Error(), http.StatusBadRequest)
		}
		return
	}

	if err := ValidLastName(name.Last); err != nil {
		e := json.NewEncoder(w).Encode(ValidationResponse{Valid: false, Reason: err.Error()})
		if e != nil {
			http.Error(w, e.Error(), http.StatusBadRequest)
		}
		return
	}

	err = json.NewEncoder(w).Encode(ValidationResponse{Valid: true})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
