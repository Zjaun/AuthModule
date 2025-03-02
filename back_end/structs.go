package back_end

type User struct {
	Username  string
	Position  string
	FirstName string
	LastName  string
	Email     string
	Password  string
	Questions *UserSecurity
}

type UserSecurity struct {
	Sq1    int
	Sq1Ans string
	Sq2    int
	Sq2Ans string
	Sq3    int
	Sq3Ans string
}

type Question struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
