module ajax

go 1.17

replace hasuragraphql => ../../hasuragraphql

require (
	hasuragraphql v0.0.0-00010101000000-000000000000
	home v0.0.0-00010101000000-000000000000
)

require (
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/gorilla/sessions v1.2.1 // indirect
	templates v0.0.0-00010101000000-000000000000 // indirect
)

replace home => ../home

replace templates => ../../templates
