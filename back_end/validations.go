package back_end

import (
	"net/mail"
	"regexp"
	"strconv"
	"strings"
)

var (
	firstNameRegex = regexp.MustCompile(`[\p{L} '-]*\p{L}$`)
	lastNameRegex  = regexp.MustCompile(`[\p{L}'-]*\p{L}$`)
	suffixRegex    = regexp.MustCompile(`[\p{L}.]*$`)
)

func commonChecks(f, val string) (valid bool, field, reason string) {
	val = strings.TrimSpace(val)
	if len(strings.TrimSpace(val)) < 1 {
		return false, f, "Cannot be empty."
	}
	return true, "", ""
}

func validFirst(first string) (valid bool, field, reason string) {
	f := "Name"
	if v, f, r := commonChecks(f, first); !v {
		return v, f, r
	}
	// if !unicode.IsLetter(rune(first[0])) {
	// 	return false, f, "Must begin with a letter."
	// }
	if !firstNameRegex.MatchString(first) {
		return false, f, "Only letters, hyphens, apostrophes, and spaces are allowed."
	}
	return true, "", ""
}

func validLast(last string) (valid bool, field, reason string) {
	f := "Name"
	if v, f, r := commonChecks(f, last); !v {
		return v, f, r
	}
	// if !unicode.IsLetter(rune(last[0])) {
	// 	return false, f, "Must begin with a letter."
	// }
	if !lastNameRegex.MatchString(last) {
		return false, f, "Only letters, hyphens, and apostrophes are allowed."
	}
	return true, "", ""
}

func validEmail(email string) (valid bool, field, reason string) {
	f := "Email"
	if v, f, r := commonChecks(f, email); !v {
		return v, f, r
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return false, f, "Invalid email address."
	}
	return true, "", ""
}

func validSuffix(suffix string) (valid bool, field, reason string) {
	suffix = strings.TrimSpace(suffix)
	f := "Suffix"
	if len(suffix) < 1 {
		return true, "", ""
	}
	if !suffixRegex.MatchString(suffix) {
		return false, f, "Only letters and periods are allowed"
	}
	return true, "", ""
}

func validLength(s string, length int) bool {
	return len(s) <= length
}

func ValidateRegistration(r *RegistrationRequest) *ValidationResponse {
	if r == nil {
		return &ValidationResponse{Valid: false, Field: "Err", Reason: "Registration is nil."}
	}
	if v, f, r := validFirst(r.FirstName); !v {
		return &ValidationResponse{Valid: v, Field: f, Reason: r}
	}
	if v, f, r := validLast(r.LastName); !v {
		return &ValidationResponse{Valid: v, Field: f, Reason: r}
	}
	if v, f, r := validEmail(r.Email); !v {
		return &ValidationResponse{Valid: v, Field: f, Reason: r}
	}
	if UsernameExists(r.Username) {
		return &ValidationResponse{Valid: false, Field: "User", Reason: "Username already exists."}
	}
	if EmailExists(r.Email) {
		return &ValidationResponse{Valid: false, Field: "Email", Reason: "Email already exists."}
	}
	sq1, _ := strconv.Atoi(r.Q1)
	sq2, _ := strconv.Atoi(r.Q2)
	sq3, _ := strconv.Atoi(r.Q3)
	questions := UserSecurity{
		Sq1:    sq1,
		Sq1Ans: r.Q1Ans,
		Sq2:    sq2,
		Sq2Ans: r.Q2Ans,
		Sq3:    sq3,
		Sq3Ans: r.Q3Ans,
	}
	user := User{
		Username:  r.Username,
		Position:  r.Position,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Email:     r.Email,
		Password:  r.Password,
		Questions: &questions,
	}
	err := RegisterUser(&user)
	if err != nil {
		return &ValidationResponse{Valid: false, Field: "Err", Reason: err.Error()}
	}
	return nil
}
