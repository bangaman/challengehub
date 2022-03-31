module main

go 1.17

replace home => ./routes/home

replace templates => ./templates

require (
	accept v0.0.0-00010101000000-000000000000
	ajax v0.0.0-00010101000000-000000000000
	home v0.0.0-00010101000000-000000000000
	logout v0.0.0-00010101000000-000000000000
	register v0.0.0-00010101000000-000000000000
)

require (
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/gorilla/sessions v1.2.1 // indirect
	hasuragraphql v0.0.0-00010101000000-000000000000 // indirect
	templates v0.0.0-00010101000000-000000000000 // indirect
)

replace hasuragraphql => ./hasuragraphql

replace register => ./routes/register

replace ajax => ./routes/ajax

replace accept => ./routes/accept

replace logout => ./routes/logout
