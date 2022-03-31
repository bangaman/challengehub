package accept

import (
	"net/http"
	"strings"
	"fmt"
	"hasuragraphql"
	"encoding/json"
)


type AccepStruct struct {
	Success bool
	Postid string
}

type CheckAccept struct {
	Data struct {
		Accept []interface{} `json:"accept"`
	} `json:"data"`
}

func (l *AccepStruct) VerifyAccept(r *http.Request, id string) {
	if strings.ToLower(r.Method) == "post"{
		postid := strings.TrimSpace(r.FormValue("postid"))
		l.Postid = postid
		fmt.Println(postid, " jjjj")

		result := hasuragraphql.HasuraGraphQlApi(FindAccept(id, postid))

		var x CheckAccept
		json.Unmarshal([]byte(result), &x)

		if len(x.Data.Accept) > 0 {
			delete := hasuragraphql.HasuraGraphQlApi(DeleteAccept(id, postid))

			fmt.Println(delete)
		}else{
			insert := hasuragraphql.HasuraGraphQlApi(InsertAccept(id, postid))

			fmt.Println(insert)
		}

		l.Success = true
		// fmt.Println(result)


	}
}