package handlers

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type Registr struct {
	Username  string
	Fullname  string
	Email     string
	Password  string
	Password2 string
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost, // 10-14 норм
	)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func SuccessHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

func AuthHandler(DB *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodPost:
			err := r.ParseMultipartForm(10 << 20) // 10MB
			if err != nil {
				http.Error(w, "Ошибка обработки формы", http.StatusBadRequest)
				return
			}

			username := r.FormValue("login")
			password := r.FormValue("password")

			verification, err := VerificationUserDB(username, password, DB)
			if err != nil {
				log.Println("Ошибка проверки пользователя:", err)
				http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
				return
			}

			if verification {
				http.Redirect(w, r, "/success", http.StatusSeeOther)
				return
			}

			// если логин или пароль неверный
			http.Error(w, "Неверный логин или пароль", http.StatusUnauthorized)
			return

		case http.MethodGet:
			http.ServeFile(w, r, "./front/login.html")
			return

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
}

func RegistrHandler(DB *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			err := r.ParseMultipartForm(10 << 20)
			if err != nil {
				http.Error(w, "Error parsing form data", http.StatusBadRequest)
				return
			}
			log.Println("username:", r.FormValue("username"))
			log.Println("fullname:", r.FormValue("fullname"))
			log.Println("email:", r.FormValue("email"))

			reg := Registr{
				r.FormValue("username"),
				r.FormValue("fullname"),
				r.FormValue("email"),
				r.FormValue("password"),
				r.FormValue("password2"),
			}

			if reg.Password != reg.Password2 {
				http.Error(w, "Пароли не совпадают", http.StatusBadRequest)
				return
			}
			password_hash, err := HashPassword(reg.Password)
			if err != nil {
				log.Fatalln("error hash")
			}
			reg = Registr{
				r.FormValue("username"),
				r.FormValue("fullname"),
				r.FormValue("email"),
				password_hash,
				r.FormValue("password2"),
			}
			CreateUserDB(reg, DB)
		} else if r.Method == http.MethodGet {
			http.ServeFile(w, r, "./front/index.html")
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

	}

}
