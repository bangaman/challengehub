package hasuragraphql


import (
	"github.com/gorilla/sessions"
)


var (
	
	Store = sessions.NewCookieStore([]byte("hasurs-graphql-session-xynames-show"))
)