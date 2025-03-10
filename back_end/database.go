package back_end

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var (
	db     *sql.DB
	config mysql.Config
)

func loadDbCreds() {
	f, err := os.Open("dbCredentials")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)
	s := make([]string, 0, 3)
	for fs.Scan() {
		s = append(s, fs.Text())
	}
	fmt.Printf("Loaded DB Credentials: %v\n", s)
	config = mysql.Config{
		User:   s[0],
		Passwd: s[1],
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: s[2],
	}
}

func init() {
	loadDbCreds()
}

func Connect() error {
	var err error
	db, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return err
	}
	pingErr := db.Ping()
	if pingErr != nil {
		return pingErr
	}
	return nil
}

func registerQuestions(ts *sql.Tx, username string, q *UserSecurity) error {
	if q == nil {
		return errors.New("attempting to register a user with nil questions")
	}
	if ts == nil {
		return errors.New("attempting to register user security options with nil transaction")
	}
	h1, _ := Encrypt(q.Sq1Ans)
	h2, _ := Encrypt(q.Sq1Ans)
	h3, _ := Encrypt(q.Sq1Ans)
	_, err := ts.Exec(
		"INSERT INTO tblusersecurity (username, sq1, sq1ans, sq2, sq2ans, sq3, sq3ans) VALUES (?, ?, ?, ?, ?, ?, ?)",
		username, q.Sq1, h1, q.Sq2, h2, q.Sq3, h3,
	)
	if err != nil {
		rollErr := ts.Rollback()
		if rollErr != nil {
			return fmt.Errorf("could not rollback user security registration: %w", rollErr)
		}
		return fmt.Errorf("could not register user security options: %w", err)
	}
	return nil
}

func RegisterUser(u *User) error {
	if u == nil {
		return errors.New("attempting to register a nil user")
	}
	hashedPassword, err := Encrypt(u.Password)
	if err != nil {
		return err
	}
	ts, err := db.Begin()
	if err != nil {
		return fmt.Errorf("could not begin transaction: %w", err)
	}
	_, err = ts.Exec(
		"INSERT INTO tblusers (username, firstName, lastName, email, password, position) VALUES (?, ?, ?, ?, ?, ?)",
		u.Username, u.FirstName, u.LastName, u.Email, hashedPassword, u.Position,
	)
	if err != nil {
		rollErr := ts.Rollback()
		if rollErr != nil {
			return fmt.Errorf("could not rollback user registration: %w", rollErr)
		}
		return fmt.Errorf("could not register user: %w", err)
	}
	err = registerQuestions(ts, u.Username, u.Questions)
	if err != nil {
		return err
	}
	if err = ts.Commit(); err != nil {
		return fmt.Errorf("could not commit user registration: %w", err)
	}
	return nil
}

func EmailExists(email string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM tblusers WHERE email = ?", email).Scan(&count)
	if errors.Is(err, sql.ErrNoRows) {
		return false
	}
	return count > 0
}

func UsernameExists(username string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM tblusers WHERE username = ?", username).Scan(&count)
	if errors.Is(err, sql.ErrNoRows) {
		return false
	}
	return count > 0
}

// ChangePassword changes the password for a given user.
// The password parameter must be hashed.
func ChangePassword(username string, password string) error {
	ts, err := db.Begin()
	if err != nil {
		return fmt.Errorf("could not begin transaction: %w", err)
	}
	_, err = ts.Exec("UPDATE tblUsers SET password = ? WHERE username = ?", password, username)
	if err != nil {
		rollErr := ts.Rollback()
		if rollErr != nil {
			return fmt.Errorf("could not rollback password change: %w", err)
		}
		return fmt.Errorf("could not update user password: %w", err)
	}
	if err = ts.Commit(); err != nil {
		return fmt.Errorf("could not commit user change: %w", err)
	}
	return nil
}

func Authenticate(c *LoginRequest) error {
	var retrivedUser struct {
		Username string
		Password string
	}
	err := db.QueryRow("SELECT username, password FROM tblUsers WHERE username = ?", c.Username).Scan(&retrivedUser.Username, &retrivedUser.Password)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("Invalid Login Request")
	} else if err != nil {
		return fmt.Errorf("dbCredentials error: %w", err)
	}
	if retrivedUser.Username != c.Username || !Compare(retrivedUser.Password, c.Password) {
		return fmt.Errorf("Invalid Login Request")
	}
	return nil
}

func GetQuestions(username string) (questions []string, err error) {
	var q1, q2, q3 int
	err = db.QueryRow("SELECT sq1, sq2, sq3 FROM tblUserSecurity WHERE username = ?", username).Scan(&q1, &q2, &q3)
	if err != nil {
		return nil, fmt.Errorf("could not get questions: %w", err)
	}
	return []string{Questions[q1], Questions[q2], Questions[q3]}, nil
}

func GetAnswer(username string, colName string) (answer string, err error) {
	colName = colName + "Ans"
	query := "SELECT " + colName + " FROM tblUserSecurity WHERE username = ?"
	err = db.QueryRow(query, username).Scan(&answer)
	if err != nil {
		return "", fmt.Errorf("could not get answer: %w", err)
	}
	return answer, nil
}
