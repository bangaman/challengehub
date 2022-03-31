package accept

import (
	"net/http"
	"strings"
	"fmt"
	"hasuragraphql"
	"encoding/json"
)


type PublishStruct struct {
	Success bool
	Postid string
}

type CheckPublish struct {
	Data struct {
		Publish []interface{} `json:"publish"`
	} `json:"data"`
}

func (l *PublishStruct) VerifyPublish(r *http.Request, id string) {
	if strings.ToLower(r.Method) == "post"{
		del := strings.TrimSpace(r.FormValue("delete"))
		postid := strings.TrimSpace(r.FormValue("postid"))
		l.Postid = postid

		if len(del) > 0 {
			delete := hasuragraphql.HasuraGraphQlApi(DeletePublish(del))

			fmt.Println(delete)
			l.Success = true
		}else{
			story := strings.TrimSpace(r.FormValue("story"))
		    code := DoubleQuote(strings.TrimSpace(r.FormValue("code")))
		    // fmt.Println(postid, " jjjj")

		    result := hasuragraphql.HasuraGraphQlApi(FindPublish(id))

		    var x CheckPublish
		    json.Unmarshal([]byte(result), &x)

		    fmt.Println(len(x.Data.Publish))

		    if len(x.Data.Publish) > 0 {
			   fmt.Println("Inserterd already")
		    }else{
			   insert := hasuragraphql.HasuraGraphQlApi(InsertPublish(story, code, id, postid))

			   fmt.Println(insert)
			   l.Success = true
		    }

		}
		
	}
}



func DoubleQuote(text string) string {
	brks := strings.Split(text, " ")


	for i, k := range brks {
		brks[i] = strings.Replace(strings.ToLower(k), `"`, "-----", -1)
	}

	for i, k := range brks {
		brks[i] = strings.Replace(strings.ToLower(k), "\n", "......", -1)
	}

	fmt.Println(brks)

	return strings.Join(brks, " ")
}
