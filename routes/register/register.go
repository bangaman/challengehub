package register

import (
	"encoding/json"
	"net/http"
	"html/template"
	"hasuragraphql"
	"strings"
	"fmt"
)

type RegisterStruct struct {
	Username string
	UserErr template.HTML
	Email string
	EmailErr template.HTML
	Password string
	Usersid string
	Sucess  bool
	Error template.HTML
}

type CheckIfRegistered struct {
	Data struct {
		InsertRagisterOne struct {
			Usersid int `json:"usersid"`
		} `json:"insert_ragister_one"`
	} `json:"data"`
}


func (d *RegisterStruct)  RegisterVerify(r *http.Request){
	if strings.ToLower(r.Method) == "post"{
		d.Username = strings.TrimSpace(r.FormValue("username"))
		d.Email = strings.TrimSpace(r.FormValue("email"))
		d.Password = strings.TrimSpace(r.FormValue("password"))

		//erro slice
		catcherr := []string{}

		//check for user in database
		result := hasuragraphql.HasuraGraphQlApi(hasuragraphql.AllUsersQuery())
		var x hasuragraphql.RegisterStruct
		json.Unmarshal([]byte(result), &x)

		fmt.Println(x)
		// fmt.Println(result)

		//check user cant be found in database
		if len(x.Data.Ragister) > 0 {

			for _, item := range x.Data.Ragister {
				
				if item.Username == d.Username {
					d.UserErr = template.HTML("<span style='color:red;'><i>Username Taken Already</i></span>")
					catcherr = append(catcherr, "error")
				}

				if item.Email == d.Email {
					d.EmailErr = template.HTML("<div style='padding-top:10px;'><span style='color:red;'><i>Email Taken Already</i></span></div>")
					catcherr = append(catcherr, "error")
				}

			}

		}

		if len(catcherr) > 0 {
			d.Sucess = false
			fmt.Println(catcherr)
		}else{
			insert := hasuragraphql.HasuraGraphQlApi(hasuragraphql.InsertUser(d.Username, d.Email, d.Password))

			//unmarshal returned json
			var i CheckIfRegistered
			json.Unmarshal([]byte(insert), &i)

			//check if user id is returned, user is inserted
			if i.Data.InsertRagisterOne.Usersid > 0 {
				d.Sucess = true
				fmt.Println("inserted already")
			}
		}

	}
}