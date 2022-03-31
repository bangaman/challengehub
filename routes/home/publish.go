package home

import (
	"encoding/json"
	"hasuragraphql"
	"fmt"
	"strings"
	"strconv"
	// "fmt"
)

type PublishStruct struct {
	Data struct {
		Publish []struct {
			Story           string `json:"story"`
			Code            string `json:"code"`
			Usersid         int    `json:"usersid"`
			Postid          string `json:"postid"`
			Publishid       int `json:"publishid"`
			RagisterPublish struct {
				Username string `json:"username"`
			} `json:"ragister_publish"`
		} `json:"publish"`
	} `json:"data"`
}

func AllPublish(postid, id, heading string)  (bool, string) {
	result := hasuragraphql.HasuraGraphQlApi(FindPublish())
	var x PublishStruct
	json.Unmarshal([]byte(result), &x)

	publishout := []string{}
	checkuser := []string{}

	for _, x := range x.Data.Publish {
		if x.Postid == postid {

			if strconv.Itoa(x.Usersid) == id {
				checkuser = append(checkuser, id)

			}
			fmt.Println(strconv.Itoa(x.Usersid), id)
			publishout = append(publishout, "<div class='published'><div class='child-published'>")

			if strconv.Itoa(x.Usersid) == id {
				publishout = append(publishout, "<div style='padding-bottom:10px;display:none;' id='main-delete' class='main-delete' ><div class='delete'>")
			    publishout = append(publishout, "<form action='/publish/accept/' method='post'><div style='display:none;' id='all'><input type='text' name='delete' value='"+strconv.Itoa(x.Publishid)+"'> <input type='text' name='postid' value='"+postid+"'/> </div><span> Continue with delete <button>Continue</button></span></form>")
			    publishout = append(publishout, "</div></div>")
			}

			publishout = append(publishout, "<div class='published-top'>")

			//details
			publishout = append(publishout, "<div id='pic'><button>"+x.RagisterPublish.Username[0:1]+"</button></div>")
			publishout = append(publishout, "<div id='name'><span>"+x.RagisterPublish.Username+"</span> <div id='trash'><img style='pointer-events:none;' src='/assets/icons/trash.svg' /></div></div>")
			// publishout = append(publishout, "<div id='date'><span>"+item.Date+"</span></div>")
			//details
			publishout = append(publishout, "</div>")

			//heading
			publishout = append(publishout, "<div class='heading'><button><i><b>TAG</b> : "+heading+"</i></button></div>")
			//heading

			//general 
			publishout = append(publishout, "<div class='general-wrap'>")

			//open f - general wrap
			publishout = append(publishout, "<div class='f-general-wrap'>")

			//story
			publishout = append(publishout, "<div class='story'><div id='story'><span>"+x.Story+"</span></div></div>")
			//story

			publishout = append(publishout, "<div style='padding-top:10px;' class='code'>")
			publishout = append(publishout, MoldVariable(x.Code))
			publishout = append(publishout, "</div>")

			//close f-general wrap
			publishout = append(publishout, "</div>")

			//open t - general wrap
			publishout = append(publishout, "<div class='f-general-wrap'>")

			//close t-general wrap
			publishout = append(publishout, "</div>")

			//close general
			publishout = append(publishout, "</div>")


			publishout = append(publishout, "</div></div>")
		}
	}

	var checker bool 

	if len(checkuser) > 0 {
		checker = true
	}else{
		checker = false
	}

	fmt.Println(checkuser)

	return checker, strings.Join(publishout, " ")
	
}