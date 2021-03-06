package main

import (
	"net/http"
	"os"
	"log"
	"encoding/json"
	"fmt"
	// "hasuragraphql"
	"home"
	"register"
	"ajax"
	"accept"
	"logout"
)


func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "relative/path/to/favicon.ico")
}

func main(){

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/", home.Home)
	http.HandleFunc("/profile/", home.Profile)
	http.HandleFunc("/make/challenge/", home.Challenge)
	http.HandleFunc("/story/", home.Story)
	http.HandleFunc("/views/", home.View)
	http.HandleFunc("/register/", register.Register)
	http.HandleFunc("/login/", register.Login)
	http.HandleFunc("/logout/", logout.Logout)
	http.HandleFunc("/ajax/react/", ajax.React)
	http.HandleFunc("/in/accept/", accept.Accept)
	http.HandleFunc("/publish/accept/", accept.Publish)
	log.Fatal(http.ListenAndServe(Getenver(), nil))
}


func Getenver() string {
	var port = os.Getenv("PORT")

	if port == ""{
		port = "8080"
		fmt.Println("INFO: no port detected")
	}else{
		fmt.Println("HEROKU PORT FOUND "+port)
	}

	return ":"+port
}

type AutoGenerated struct {
	Data struct {
		Movies []struct {
			URL   string `json:"url"`
			Title string `json:"title"`
		} `json:"movies"`
	} `json:"data"`
}

func Change () {
	var x AutoGenerated
	data := `{"data":{"movies":[{"url":"https://riseofgraphqlwarriorpt1.com/","title":"Rise of GraphQL Warrior Pt1"}, {"url":"https://riseofgraphqlwarriorpt1.com/","title":"Rise of GraphQL Warrior Pt1"}, {"url":"the bango","title":"homer"}, {"url":"basho","title":"come home"}, {"url":"grace boy","title":"stay lead"}]}}`
	json.Unmarshal([]byte(data), &x)
    // fmt.Println(err)
	fmt.Println(x)
}

func Parse(text string){

	// vals := "@image-1728592637.jpg#10==shadda@image-206769599.jpg#5000==wando"
}
