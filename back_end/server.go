package back_end

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Forgot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Forgot Request Received.")
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "POST requests only.", http.StatusMethodNotAllowed)
		fmt.Printf("Incorrect method: %s\n", r.Method)
		return
	}

	var req ResetRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("cannot parse body")
		}
		fmt.Printf("cannot decode body: %s\n", string(body))
		return
	}

	fmt.Printf("Request: %v\n", req)

	if !UsernameExists(req.Username) {
		fmt.Println("Username does not exist.")
		err := json.NewEncoder(w).Encode(&ResetResponse{Success: false, Field: "User", Reason: "Username does not exist."})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	fmt.Println("Username exists, attempting to retrive/validate response")

	if req.ValUser && !req.ValQues {
		if res, err := GetQuestions(req.Username); err != nil {
			fmt.Printf("Database related error: %s\n", err.Error())
			err = json.NewEncoder(w).Encode(&ResetResponse{Success: false, Field: "Err", Reason: err.Error()})
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			return
		} else {
			fmt.Printf("Sent questions to the client: %s\n", res)
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
			fmt.Printf("Database related error: %s\n", err.Error())
			err = json.NewEncoder(w).Encode(&ResetResponse{Success: false, Field: "Err", Reason: err.Error()})
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			return
		}
		if !Compare(ans, req.Answer) {
			fmt.Println("Provided answer is incorrect.")
			err = json.NewEncoder(w).Encode(&ResetResponse{Success: false, Field: "Answer", Reason: "Incorrect answer."})
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			return
		}
	}

	hashedPassword, err := Encrypt(req.Password)

	if err != nil {
		fmt.Printf("Hashing related error: %s\n", err)
		err = json.NewEncoder(w).Encode(&ResetResponse{Success: false, Field: "Err", Reason: err.Error()})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	err = ChangePassword(req.Username, hashedPassword)
	if err != nil {
		fmt.Printf("Database related error: %s\n", err)
		err = json.NewEncoder(w).Encode(&ResetResponse{Success: false, Field: "Err", Reason: err.Error()})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	fmt.Printf("Successfully changed %s's password.", req.Username)

	err = json.NewEncoder(w).Encode(&ResetResponse{Success: true})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

}

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register Request Received")
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "POST requests only.", http.StatusMethodNotAllowed)
		fmt.Printf("Incorrect method: %s\n", r.Method)
		return
	}

	var reg RegistrationRequest
	err := json.NewDecoder(r.Body).Decode(&reg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("cannot parse body")
		}
		fmt.Printf("cannot decode body: %s\n", string(body))
		return
	}

	fmt.Printf("Request body: %v\n", reg)

	if res := ValidateRegistration(&reg); res != nil {
		fmt.Printf("Cannot register user: %s\n", res.Reason)
		e := json.NewEncoder(w).Encode(res)
		if e != nil {
			http.Error(w, e.Error(), http.StatusBadRequest)
		}
		return
	}

	fmt.Printf("Successfully registered this user: %v\n", reg)

	err = json.NewEncoder(w).Encode(ValidationResponse{Valid: true, Reason: ""})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login Request Received")
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "POST requests only.", http.StatusMethodNotAllowed)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("cannot parse body")
		}
		fmt.Printf("cannot decode body: %s\n", string(body))
		return
	}

	var c LoginRequest

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Request Body: %v\n", c)

	err = Authenticate(&c)

	if err != nil {
		fmt.Printf("Cannot authenticate %s: %v\n", c.Username, err.Error())
		e := json.NewEncoder(w).Encode(AuthenticationResponse{LoggedIn: false, Reason: err.Error()})
		if e != nil {
			http.Error(w, e.Error(), http.StatusBadRequest)
		}
		return
	}

	fmt.Printf("Successfully authenticated %s.\n", c.Username)

	err = json.NewEncoder(w).Encode(AuthenticationResponse{LoggedIn: true, Reason: ""})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

}
