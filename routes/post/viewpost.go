package post

import {
	"encoding/json"
	"hasuragraphql"
	"fmt"
}

func ShowView() {
	// result := HasuraGraphQlApi(FindPost())
	var x PostStruct
	json.Unmarshal([]byte(hasuragraphql.Posts), &x)

	fmt.Println(x)
}