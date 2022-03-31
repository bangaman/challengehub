package home

import (
	"html/template"
)

type ProfileStruct struct{
	Details template.HTML
	Accepts template.HTML
	Username string
	User template.HTML
}


func (l *ProfileStruct) ProfileDetails (name string){
	l.User = template.HTML("<button>"+name[:1]+"</button>")
	l.Username = name
}