package ajax

import (
	"net/http"
	"hasuragraphql"
	"fmt"
)


func React(w http.ResponseWriter, r *http.Request) {
	let := ReactStruct{}
	session, _ := hasuragraphql.Store.Get(r, "session")

	if session.Values["username"] != nil {

		if len(session.Values["username"].(string)) > 1{
			let.VerifyReact(r, session.Values["usersid"].(string))
			fmt.Fprintf(w, "%s", let.Details)

		}else{
			http.Redirect(w, r, "/views/", http.StatusSeeOther)
		}
	}else{
		http.Redirect(w, r, "/views/", http.StatusSeeOther)
	}
}

