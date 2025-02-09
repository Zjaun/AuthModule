package back_end

type RegistrationRequest struct {
	Username  string `json:"username"`
	FirstName string `json:"first"`
	LastName  string `json:"last"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Q1        string `json:"q1"`
	Q1Ans     string `json:"q1Ans"`
	Q2        string `json:"q2"`
	Q2Ans     string `json:"q2Ans"`
	Q3        string `json:"q3"`
	Q3Ans     string `json:"q3Ans"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResetRequest struct {
	ValUser  bool   `json:"valUser"`
	Username string `json:"username"`
	ValQues  bool   `json:"valQues"`
	Question string `json:"question"` // column name!
	Answer   string `json:"answer"`
	Password string `json:"password"`
}
