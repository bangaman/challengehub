package register

import (
	"encoding/json"
	"net/http"
	"html/template"
	"hasuragraphql"
	"strings"
	"strconv"
	"fmt"
)

type LoginStruct struct {
	Username string
	Email string
	Usersid string
	Sucess  bool
	Error template.HTML
}


func (d *LoginStruct)  LoginVerify(r *http.Request){
	if strings.ToLower(r.Method) == "post"{
		d.Username = strings.TrimSpace(r.FormValue("username"))
		password := strings.TrimSpace(r.FormValue("password"))

		//check for user in database
		result := hasuragraphql.HasuraGraphQlApi(hasuragraphql.FindUser(d.Username, password))
		var x hasuragraphql.RegisterStruct
		fmt.Println(result)
		json.Unmarshal([]byte(result), &x)

		fmt.Println(x)

		//check user cant be found in database
		if len(x.Data.Ragister) > 0 {
			for _, i := range x.Data.Ragister {
				d.Username = i.Username
				d.Email = i.Email
				d.Usersid = strconv.Itoa(i.Usersid)
			}
			fmt.Println(d.Usersid, " ID")
			d.Sucess = true
		}else{
			d.Sucess = false
			d.Error = template.HTML("<span style='color:red;'><i>Invalid Credentials</i></span>")
		}

	}
}