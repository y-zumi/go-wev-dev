package main

import (
	"github.com/GoesToEleven/golang-web-dev/042_mongodb/10_hands-on/starting-code/controllers"
	"github.com/GoesToEleven/golang-web-dev/042_mongodb/10_hands-on/starting-code/session"
	"html/template"
	"net/http"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	session.LastCleaned = time.Now()
}

func main() {
	c := controllers.NewController(tpl)
	http.HandleFunc("/", c.Index)
	http.HandleFunc("/bar", c.Bar)
	http.HandleFunc("/signup", c.Signup)
	http.HandleFunc("/login", c.Login)
	http.HandleFunc("/logout", c.Logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe("localhost:8080", nil)
}