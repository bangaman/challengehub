package templates


import (
    "html/template"
    "net/http"
)


func Render(w http.ResponseWriter, filename string, data interface{}){
	t, err := template.ParseFiles(filename)

	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	err = t.Execute(w, data)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

	// return t
}