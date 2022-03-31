package hasuragraphql

import (
	"fmt"
)


func FindAllAccept() map[string]string {
	jsonData := map[string]string {
        "query":
			`{
				accept{
					accepter
					post
				}
			}	

            `,
    }

	return jsonData
}

func FindAccept(a, b string) map[string]string {
	jsonData := map[string]string {
        "query": fmt.Sprintf(
			`{
				accept(where: {accepter: {_eq: "%s"}, _and: {}, post: {_eq: "%s"}}) {
					accepter
					post
				  }
			}	
            `,
		a, b),
    }

	return jsonData
}


func InsertAccept(a, b string) map[string]string {
	jsonInsert := map[string]string{
		"query": fmt.Sprintf(
			`
			mutation MyMutation {
				insert_accept_one(object: {accepter: "%s", post: "%s"})
				{
					accepterid
				}
			  }			  
			`,
		a,b),
          
    }

	return jsonInsert
} 


func DeleteAccept(a, b string) map[string]string {
	jsonInsert := map[string]string{
		"query": fmt.Sprintf(
			`
			mutation MyMutation {
				delete_accept(where: {accepter: {_eq: "%s"}, _and: {post: {_eq: "%s"}}}) {
					returning {
						post
					}
				}
			  }			  
			`,
		a,b),
          
    }

	return jsonInsert
} 
