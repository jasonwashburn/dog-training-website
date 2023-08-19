package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jasonwashburn/dog-training-website/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	// For now, lets just return a message saying this is the registration page
	fmt.Fprintf(w, "This is the registration page")
}

func GetLoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/user/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func PostLoginHandler(w http.ResponseWriter, r *http.Request) {
	testUserEmail := "test@example.com"
	testUserPassword := "test"
	email := r.FormValue("email")
	password := r.FormValue("password")
	hash, err := HashPassword(password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user := models.User{
		Email:    email,
		Password: hash,
	}
	if user.Email == testUserEmail && CheckPasswordHash(testUserPassword, user.Password) {
		fmt.Fprintf(w, "Login successful")
	} else {
		fmt.Fprintf(w, "Login failed")
	}
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
