package views

import (
	"html/template"
)

func StartPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/user_home.html"))
	return temp
}