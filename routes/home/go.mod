module home

go 1.17

replace templates => ../../templates

require (
	hasuragraphql v0.0.0-00010101000000-000000000000
	templates v0.0.0-00010101000000-000000000000
)

require (
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/gorilla/sessions v1.2.1 // indirect
)

replace hasuragraphql => ../../hasuragraphql

replace post => ./post
