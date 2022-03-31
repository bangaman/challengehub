package logout

import (
	"net/http"
	"hasuragraphql"
)

func Logout(w http.ResponseWriter, r *http.Request){
	session, _ := hasuragraphql.Store.Get(r, "session")
	session.Values["username"] = nil
	session.Values["email"] = nil
	session.Values["usersid"] = nil
	session.Save(r, w)
	http.Redirect(w, r, "/view", http.StatusSeeOther)
}