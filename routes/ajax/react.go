package ajax

import (
	"net/http"
	"fmt"
	"hasuragraphql"
	"html/template"
	"strings"
	"encoding/json"
	"home"
)

type ReactStruct struct {
	Details template.HTML
}

type CheckReact struct {
	Data struct {
		React []interface{} `json:"react"`
	} `json:"data"`
}

func (l *ReactStruct) VerifyReact(r *http.Request, id string) {
	fmt.Println(id, "ID")
	if strings.ToLower(r.Method) == "post"{
		postid := strings.TrimSpace(r.FormValue("postid"))
		result := hasuragraphql.HasuraGraphQlApi(hasuragraphql.FindReact(id, postid))
		var x CheckReact
		json.Unmarshal([]byte(result), &x)
		// fmt.Println(x)

		if len(x.Data.React) > 0 {
			fmt.Println("inserted")

			delete := hasuragraphql.HasuraGraphQlApi(hasuragraphql.DeleteReact(id, postid))

			fmt.Println(delete)
		}else{
			fmt.Println("not liked")

			insert := hasuragraphql.HasuraGraphQlApi(hasuragraphql.InsertReact(id, postid))

			//unmarshal returned json
			// var i CheckIfRegistered
			// json.Unmarshal([]byte(insert), &i)
			fmt.Println(insert)

		}

		l.Details = template.HTML(home.GetLike(id, postid))

	}
}