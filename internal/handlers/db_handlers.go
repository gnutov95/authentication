package handlers

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

func CreateUserDB(reg Registr, db *sql.DB) error {

	t := time.Now()
	createData := t.Format("2006-01-02 15:04:05")
	role := "student"

	res, err := db.Exec("INSERT INTO new_table(username, email, password_hash, full_name, role, created_at, updated_at) VALUES (?,?,?,?,?,?,?)", reg.Username, reg.Email, reg.Password, reg.Fullname, role, createData, createData)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	lastID, _ := res.LastInsertId()
	log.Printf("Добавлено строк: %d, последний ID: %d\n", rowsAffected, lastID)
	return nil
}

func VerificationUserDB(login, password string, db *sql.DB) (bool, error) {
	var passwordBD string
	log.Println(login)
	log.Println(password)

	err := db.QueryRow(
		"SELECT password_hash FROM new_table WHERE username = ? OR email = ? LIMIT 1",
		login, login,
	).Scan(&passwordBD)

	// если пользователь не найден
	if err == sql.ErrNoRows {
		log.Println("пользователь не найден")
		return false, nil
	}

	// если другая ошибка БД
	if err != nil {
		log.Println("ошибка бд")
		return false, err
	}

	// сравниваем пароль с хешем из БД
	err = bcrypt.CompareHashAndPassword([]byte(passwordBD), []byte(password))
	if err != nil {
		log.Println("пароль не совпал с хешем")
		return false, nil // пароль неверный
	}
	log.Println("всё ок")

	return true, nil
}
