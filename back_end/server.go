package back_end

import (
	"encoding/json"
	"net/http"
)

func Forgot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "POST requests only.", http.StatusMethodNotAllowed)
		return
	}

	var req ResetRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !UsernameExists(req.Username) {
		err := json.NewEncoder(w).Encode(&ResetResponse{Success: false, Field: "User", Reason: "Username does not exist."})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	if req.ValUser && !req.ValQues {
		if res, err := GetQuestions(req.Username); err != nil {
			err = json.NewEncoder(w).Encode(&ResetResponse{Success: false, Field: "Err", Reason: err.Error()})
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			return
		} else {
			err = json.NewEncoder(w).Encode(&ResetResponse{Question1: res[0], Question2: res[1], Question3: res[2]})
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			return
		}
	}

	if req.ValQues {
		ans, err := GetAnswer(req.Username, req.Question)
		if err != nil {
			err = json.NewEncoder(w).Encode(&ResetResponse{Success: false, Field: "Err", Reason: err.Error()})
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			return
		}
		if ans != req.Answer {
			err = json.NewEncoder(w).Encode(&ResetResponse{Success: false, Field: "Answer", Reason: "Incorrect answer."})
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			return
		}
	}

	hashedPassword, err := Encrypt(req.Password)

	if err != nil {
		err = json.NewEncoder(w).Encode(&ResetResponse{Success: false, Field: "Err", Reason: err.Error()})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	err = ChangePassword(req.Username, hashedPassword)
	if err != nil {
		err = json.NewEncoder(w).Encode(&ResetResponse{Success: false, Field: "Err", Reason: err.Error()})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	err = json.NewEncoder(w).Encode(&ResetResponse{Success: true})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "POST requests only.", http.StatusMethodNotAllowed)
		return
	}

	var reg RegistrationRequest
	err := json.NewDecoder(r.Body).Decode(&reg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if res := ValidateRegistration(&reg); res != nil {
		e := json.NewEncoder(w).Encode(res)
		if e != nil {
			http.Error(w, e.Error(), http.StatusBadRequest)
		}
		return
	}

	err = json.NewEncoder(w).Encode(ValidationResponse{Valid: true, Reason: ""})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "POST requests only.", http.StatusMethodNotAllowed)
		return
	}

	var c LoginRequest

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = Authenticate(&c)

	if err != nil {
		e := json.NewEncoder(w).Encode(AuthenticationResponse{LoggedIn: false, Reason: err.Error()})
		if e != nil {
			http.Error(w, e.Error(), http.StatusBadRequest)
		}
		return
	}

	err = json.NewEncoder(w).Encode(AuthenticationResponse{LoggedIn: true, Reason: ""})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

}
