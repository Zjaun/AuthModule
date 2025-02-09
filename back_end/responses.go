package back_end

type ValidationResponse struct {
	Valid  bool   `json:"valid"`
	Field  string `json:"field,omitempty"`
	Reason string `json:"reason,omitempty"`
}

type AuthenticationResponse struct {
	LoggedIn bool   `json:"logged_in"`
	Reason   string `json:"reason,omitempty"`
}

type ResetResponse struct {
	Success   bool   `json:"success"`
	Field     string `json:"field,omitempty"`
	Reason    string `json:"reason,omitempty"`
	Question1 string `json:"q1,omitempty"`
	Question2 string `json:"q2,omitempty"`
	Question3 string `json:"q3,omitempty"`
}
