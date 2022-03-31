package home

import (
	"fmt"
	"hasuragraphql"
	"encoding/json"
	"strings"
	"strconv"

)


func ChallengeAcceptCheck(usersid, postid string ) (bool, string){
	// FindAccept
	findall := hasuragraphql.HasuraGraphQlApi(hasuragraphql.FindAccept(usersid, postid))

	var x hasuragraphql.FindAcceptJson
	json.Unmarshal([]byte(findall), &x)

	res := []string{}
	var verfys bool

	if len(x.Data.Accept) > 0 {
		for _, item := range x.Data.Accept {

			if postid == item.Post && item.Accepter == usersid {
				verfys = true
				res = append(res, "<div id='reaction-join'><button>Welcome Dude</button></div>")
			}else{
				verfys = false
				res = append(res, "<div id='reaction-join'><button>Accept Challenge</button></div>")
			}
	
			fmt.Println(verfys, " CHECKER")
		}
	}else{
		verfys = false
		res = append(res, "<div id='reaction-join'><button>Accept Challenge</button></div>")
	}

	return verfys, strings.Join(res, " ")

}


func FindAccept(usersid string) map[int]bool {
	// FindAccept
	findall := hasuragraphql.HasuraGraphQlApi(hasuragraphql.FindAllAccept())

	var x hasuragraphql.FindAcceptJson
	json.Unmarshal([]byte(findall), &x)

	mapper := make(map[int]bool)

	if len(x.Data.Accept) > 0 {
		for _, item := range x.Data.Accept {

			if item.Accepter == usersid {
				res, _ :=strconv.Atoi(item.Post)
				mapper[res] = true
			}
		}
	}

	return mapper
}