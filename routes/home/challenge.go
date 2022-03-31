package home

import (
	"net/http"
	"strings"
	"fmt"
	"hasuragraphql"
	"html/template"
	// "fmt"
)

type ChallengeStruct struct {
	User string
	Username string
	Success bool
	Error template.HTML
}

func (l *ChallengeStruct) VerifyChallenge(r *http.Request, id string) {
	if strings.ToLower(r.Method) == "post"{
		code := DoubleQuote(strings.TrimSpace(r.FormValue("code")))

		fmt.Println(code, "NEH")
		about := strings.TrimSpace(r.FormValue("about"))
		title := strings.TrimSpace(r.FormValue("title"))
		visibility := strings.TrimSpace(r.FormValue("visibility"))
		language := strings.TrimSpace(r.FormValue("language"))

		if len(code) > 1 && len(about) > 1 && len(title) > 1 && len(visibility) > 1 && len(language) > 1{
			result := hasuragraphql.HasuraGraphQlApi(hasuragraphql.InsertPost(about,visibility, id, code, language, title ))
			fmt.Println(result)
			l.Success = true
		}else{
			l.Success = false
			l.Error = template.HTML("<h2 style='color:red'>Something went wrong</h2>")
		}

	}

}