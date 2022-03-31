package hasuragraphql


import (
	// 
)


type RegisterStruct struct {
	Data struct {
		Ragister []struct {
			Email    string `json:"email"`
			Regdate  string `json:"regdate"`
			Username string `json:"username"`
			Usersid  int    `json:"usersid"`
		} `json:"ragister"`
	} `json:"data"`
}

type PostStruct struct {
	Data struct {
		Challenge []struct {
			Code  string `json:"code"`
			Language string `json:"language"`
			Visibility string `json:"visibility"`
			Question string `json:"question"`
			Answer          string `json:"answer"`
			ID              int    `json:"id"`
			Date string `json:"date"`
			Usersid int `json:"usersid"`
			RagisterDetails struct {
				Username string `json:"username"`
			} `json:"ragister_details"`
		} `json:"challenge"`
	} `json:"data"`
}