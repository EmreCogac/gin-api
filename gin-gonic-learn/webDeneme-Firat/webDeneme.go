package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

type User struct {
	Username string
	Password string
}

var users []User

var tmpl = template.Must(template.New("index").Parse(`
<!DOCTYPE html>
<html>
<head>
	<title>Go Web Uygulamasi</title>
</head>
<body>
	<h1>Hoş Geldiniz</h1>

	{{if not .Authenticated}}
		<h2>Kayit Ol</h2>
		<form action="/register" method="POST">
			<label>Kullanici Adi:</label>
			<input type="text" name="username"><br>
			<label>Şifre:</label>
			<input type="password" name="password"><br>
			<input type="submit" value="Kayit Ol">
		</form>

		<h2>Giriş Yap</h2>
		<form action="/login" method="POST">
			<label>Kullanici Adi:</label>
			<input type="text" name="username"><br>
			<label>Şifre:</label>
			<input type="password" name="password"><br>
			<input type="submit" value="Giriş Yap">
		</form>
	{{else}}
		<h2>Hoş Geldiniz, {{.Username}}!</h2>
		<form action="/logout" method="POST">
			<input type="submit" value="Çikiş Yap">
		</form>
	{{end}}
</body>
</html>
`))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	data := map[string]interface{}{
		"Authenticated": session.Values["authenticated"],
		"Username":      session.Values["username"],
	}
	tmpl.Execute(w, data)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		user := User{Username: username, Password: password}
		users = append(users, user)

		session, _ := store.Get(r, "session-name")
		session.Values["authenticated"] = true
		session.Values["username"] = username
		session.Save(r, w)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		for _, user := range users {
			if user.Username == username && user.Password == password {
				session, _ := store.Get(r, "session-name")
				session.Values["authenticated"] = true
				session.Values["username"] = username
				session.Save(r, w)

				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	session.Values["authenticated"] = false
	session.Values["username"] = ""
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

var store = sessions.NewCookieStore([]byte("secret-key"))

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)

	fmt.Println("Server :8080 portunda çalışıyor...")
	http.ListenAndServe(":8080", nil)
}
