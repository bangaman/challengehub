package accept

import (
	"net/http"
	"hasuragraphql"

)


func Accept(w http.ResponseWriter, r *http.Request) {
	let := AccepStruct{}
	session, _ := hasuragraphql.Store.Get(r, "session")

	if session.Values["username"] != nil {

		if len(session.Values["username"].(string)) > 1{
			let.VerifyAccept(r, session.Values["usersid"].(string))
			if let.Success == true {
				http.Redirect(w, r, "/story/?default="+let.Postid, http.StatusSeeOther)
			}

		}else{
			http.Redirect(w, r, "/views/", http.StatusSeeOther)
		}
	}else{
		http.Redirect(w, r, "/views/", http.StatusSeeOther)
	}
}

func Publish(w http.ResponseWriter, r *http.Request) {
	let := PublishStruct{}
	session, _ := hasuragraphql.Store.Get(r, "session")

	if session.Values["username"] != nil {

		if len(session.Values["username"].(string)) > 1{
			let.VerifyPublish(r, session.Values["usersid"].(string))

			if let.Success == true {
				http.Redirect(w, r, "/story/?default="+let.Postid, http.StatusSeeOther)
			}
			// let.VerifyReact(r, session.Values["usersid"].(string))
			// fmt.Fprintf(w, "%s", let.Details)

		}else{
			http.Redirect(w, r, "/views/", http.StatusSeeOther)
		}
	}else{
		http.Redirect(w, r, "/views/", http.StatusSeeOther)
	}
}