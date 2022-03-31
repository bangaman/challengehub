package home

import (
	"encoding/json"
	"hasuragraphql"
	"fmt"
	"strconv"
)

type LikeStruct struct {
	Data struct {
		React []struct {
			Post  string `json:"post"`
			Liker string `json:"liker"`
			Reactid int `json:"reactid"`
		} `json:"react"`
	} `json:"data"`
}

func GetLike(id, postid string) string {
	findall := hasuragraphql.HasuraGraphQlApi(hasuragraphql.FindAllReact())

	var x LikeStruct
	json.Unmarshal([]byte(findall), &x)

	var getliker bool
	countlike := []string{}

	countall := []string{}
	for _, item := range x.Data.React {

		fmt.Println(item.Reactid, "HHH")


		if item.Post == postid {
			countall = append(countall, item.Post)
		}

		if item.Liker == id && item.Post == postid{
			getliker = true
			countlike = append(countlike, item.Post)
		}
	}

	var showcountall string

	if len(countall) > 0 {
		showcountall = strconv.Itoa(len(countall))
	}else{
		showcountall = ""
	}
	
	if len(countlike) > 0 {
		if getliker == true {
			return "<button data-post='"+postid+"' id='btn-react'><img style='pointer-events:none;' src='/assets/icons/heartyl.svg' /> "+showcountall+"</button>"
		}else{
			return "<button data-post='"+postid+"' id='btn-react'><img style='pointer-events:none;' src='/assets/icons/heartbl.svg' /> "+showcountall+"</button>"
		}
	}else{
		return "<button data-post='"+postid+"' id='btn-react'><img style='pointer-events:none;' src='/assets/icons/heartbl.svg' /> "+showcountall+"</button>"
	}

	return ""
}