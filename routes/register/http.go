package register

import (
	"net/http"
	"hasuragraphql"
	"templates"
	// "fmt"
)



func Register(w http.ResponseWriter, r *http.Request) {
	let := RegisterStruct{}
	let.RegisterVerify(r)
	session, _ := hasuragraphql.Store.Get(r, "session")
	if session.Values["username"] != nil {

		if len(session.Values["username"].(string)) > 1{
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}else{
			templates.Render(w, "assets/templates/register/register.html", let)
		}
	}else{
		if let.Sucess == true {
			http.Redirect(w, r, "/login/", http.StatusSeeOther)
		}else{
			templates.Render(w, "assets/templates/register/register.html", let)
		}
	}
}



func Login(w http.ResponseWriter, r *http.Request) {
	let :=  LoginStruct{}
	let.LoginVerify(r)
	session, _ := hasuragraphql.Store.Get(r, "session")
	if session.Values["username"] != nil {

		if len(session.Values["username"].(string)) > 1{
			http.Redirect(w, r, "/views/", http.StatusSeeOther)
		}else{
			templates.Render(w, "assets/templates/register/login.html", let)
		}
	}else{
		if let.Sucess == true {
			session, _ := hasuragraphql.Store.Get(r, "session")
			session.Values["username"] = let.Username
			session.Values["email"] = let.Email
			session.Values["usersid"] = let.Usersid
			session.Save(r, w)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}else{
			templates.Render(w, "assets/templates/register/login.html", let)
		}
	}
}