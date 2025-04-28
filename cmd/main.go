package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type User struct {
	Name     string
	Password string
	Email    string
}

type TemplateData struct {
	Error string
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("templates/register.html")
		if err != nil {
			http.Error(w, "Ошибка загрузки шаблона", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
		return
	}
	if r.Method == http.MethodPost {
		var user User
		tmpl, err := template.ParseFiles("templates/register.html")
		if err != nil {
			http.Error(w, "Ошибка загрузки шаблона", http.StatusInternalServerError)
			return
		}
		r.ParseForm()

		user.Name = r.FormValue("name")
		password := r.FormValue("password")
		user.Email = r.FormValue("email")
		confirmPassword := r.FormValue("confirm_password")

		if password != confirmPassword {
			data := TemplateData{
				Error: "Пароли не совпадают!",
			}
			tmpl.Execute(w, data)
			return
		}
		user.Password = password

		http.Redirect(w, r, "/success", http.StatusSeeOther)
		return
	}
}

func SuccessHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Регистрация успешна!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/register", RegisterHandler)
	mux.HandleFunc("/success", SuccessHandler)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	srv.ListenAndServe()
}
