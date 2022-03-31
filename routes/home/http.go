package home


import (
	"log"
	"fmt"
	"net/http" 
	"html/template"
	"templates"
	"hasuragraphql"
)


func Home(w http.ResponseWriter, r *http.Request) {
	let := HomeStruct{}
	session, _ := hasuragraphql.Store.Get(r, "session")

	if session.Values["username"] != nil {

		if len(session.Values["username"].(string)) > 1{
			let.Details = template.HTML(ShowView("loggedin", session.Values["usersid"].(string)))
			let.HomeDetails(session.Values["username"].(string))
			templates.Render(w, "assets/templates/home/home.html", let)

		}else{
			http.Redirect(w, r, "/views/", http.StatusSeeOther)
		}
	}else{
		http.Redirect(w, r, "/views/", http.StatusSeeOther)
	}
}



func Profile(w http.ResponseWriter, r *http.Request) {
	let := ProfileStruct{}
	session, _ := hasuragraphql.Store.Get(r, "session")

	if session.Values["username"] != nil {
		// AcceptedChallenge

		if len(session.Values["username"].(string)) > 1{
			let.ProfileDetails(session.Values["username"].(string))

			keys, _ := r.URL.Query()["default"]

			if len(keys) > 0 {
				key := keys[0]

			    if key != "nil" {
					if key == "accept"{
						let.Details = template.HTML(AcceptedChallenge("loggedin", session.Values["usersid"].(string)))
						templates.Render(w, "assets/templates/home/accepts.html", let)
				    }else{
						let.Details = template.HTML(UserChallenge("loggedin", session.Values["usersid"].(string)))
					    templates.Render(w, "assets/templates/home/profile.html", let)
				    }
			    }
			}else{
				let.Details = template.HTML(UserChallenge("loggedin", session.Values["usersid"].(string)))
				templates.Render(w, "assets/templates/home/profile.html", let)
			}
			
				
			// Query()["key"] will return an array of items, 
			// we only want the single item. ?key=name
			// key := keys[0]

			// if key != "nil" {
			// 	if key == "accept"{
			// 		let.Details = template.HTML(AcceptedChallenge("loggedin", session.Values["usersid"].(string)))
			// 		templates.Render(w, "assets/templates/home/accepts.html", let)
			// 	}else{
			// 		templates.Render(w, "assets/templates/home/profile.html", let)
			// 	}
			// }else{
			// 	let.Details = template.HTML(UserChallenge("loggedin", session.Values["usersid"].(string)))
			// 	templates.Render(w, "assets/templates/home/profile.html", let)
			// }

		}else{
			http.Redirect(w, r, "/views/", http.StatusSeeOther)
		}
	}else{
		http.Redirect(w, r, "/views/", http.StatusSeeOther)
	}
}


func Challenge(w http.ResponseWriter, r *http.Request) {
	let := ChallengeStruct{}
	session, _ := hasuragraphql.Store.Get(r, "session")

	if session.Values["username"] != nil {

		if len(session.Values["username"].(string)) > 1{
			let.VerifyChallenge(r, session.Values["usersid"].(string))
			let.User = session.Values["username"].(string)[0:1]
			let.Username = session.Values["username"].(string)

			if let.Success == true{
				http.Redirect(w, r, "/", http.StatusSeeOther)
			}

			templates.Render(w, "assets/templates/home/challenge.html", let)

		}else{
			http.Redirect(w, r, "/views/", http.StatusSeeOther)
		}
	}else{
		http.Redirect(w, r, "/views/", http.StatusSeeOther)
	}
}



func Story(w http.ResponseWriter, r *http.Request) {
	let := HomeStruct{}
	session, _ := hasuragraphql.Store.Get(r, "session")

	if session.Values["username"] != nil {

		if len(session.Values["username"].(string)) > 1{
			let.HomeDetails(session.Values["username"].(string))
			// GetLike(session.Values["usersid"].(string))
			keys, ok := r.URL.Query()["default"]
				
			if !ok || len(keys[0]) < 1 {
				log.Println("Url Param 'key' is missing")
				return
			}
				
			// Query()["key"] will return an array of items, 
			// we only want the single item. ?key=name
			key := keys[0]

			if key != "nil" {
				let.Details = template.HTML(SinglePost(key, session.Values["usersid"].(string)))
				fmt.Println( session.Values["usersid"].(string)," Reg")
				let.RelatedPost = template.HTML(ShowRelated("loggedin", "id", key))
				templates.Render(w, "assets/templates/home/story.html", let)
			}else{
				let.Details = template.HTML("<h1>Cant Process Your Request</h1>")
				templates.Render(w, "assets/templates/home/story.html", let)
			}

		}else{
			http.Redirect(w, r, "/views/", http.StatusSeeOther)
		}
	}else{
		http.Redirect(w, r, "/views/", http.StatusSeeOther)
	}
}



func View(w http.ResponseWriter, r *http.Request) {
	let :=  HomeStruct{}
	session, _ := hasuragraphql.Store.Get(r, "session")

	if session.Values["username"] != nil {

		if len(session.Values["username"].(string)) > 1{
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}else{
			templates.Render(w, "assets/templates/home/views.html", nil)
		}
	}else{
		let.Details = template.HTML(ShowView("out", "id"))
		templates.Render(w, "assets/templates/home/views.html", let)
	}
}