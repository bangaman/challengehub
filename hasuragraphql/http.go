package hasuragraphql


import (
	"bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
)



func HasuraGraphQlApi(Query map[string]string) string {
	//marshal json
	jsonValue, _ := json.Marshal(Query)
    request, err := http.NewRequest("POST", "https://addprojectnamehere.hasura.app/v1/graphql", bytes.NewBuffer(jsonValue))

	// Add Authorization header
	request.Header.Add("x-hasura-admin-secret", os.Getenv("HASURA_SECRET"))
    client := &http.Client{Timeout: time.Second * 10}

    response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		// return
	  }

    defer response.Body.Close()
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    }

	//get data as json 
    data, _ := ioutil.ReadAll(response.Body)
    
	return string(data)
}