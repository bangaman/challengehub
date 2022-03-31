package accept

import (
	"fmt"
)

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


func FindPublish(id string) map[string]string {
	jsonData := map[string]string {
        "query": fmt.Sprintf(
			`{
				publish(where: {usersid: {_eq: "%s"}}) {
					publishid
				  }
			}	
            `,
		id),
    }

	return jsonData
}



func InsertPublish(story, code, usersid, postid string) map[string]string {
	jsonInsert := map[string]string{
		"query": fmt.Sprintf(
			`
			mutation MyMutation {
				insert_publish_one(object: {story: "%s", code: "%s", usersid:"%s", postid:"%s"})
				{
					publishid
				}
			  }			  
			`,
		story, code, usersid, postid),
          
    }

	return jsonInsert
} 


func DeletePublish(publishid string) map[string]string {
	jsonInsert := map[string]string{
		"query": fmt.Sprintf(
			`
			mutation MyMutation {
				delete_publish_by_pk(publishid: %s) {
					publishid
				  }
			  }			  
			`,
		publishid),
          
    }

	return jsonInsert
} 
