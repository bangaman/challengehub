package hasuragraphql

import (
	"fmt"
	// "strconv"
)

func FindUser(a, b string) map[string]string {
	jsonData := map[string]string {
        "query": fmt.Sprintf(
			`{
				ragister(where: {username: {_eq: "%s"}, _and: {}, password: {_eq: "%s"}}) {
					email
					regdate
					username
					usersid
				  }
			}	
            `,
		a, b),
    }

	return jsonData
}


func FindAllReact() map[string]string {
	jsonData := map[string]string {
        "query": `{
			react {
				post
				liker
				reactid
			}
		}`,
    }

	return jsonData
}



func FindReact(a, b string) map[string]string {
	jsonData := map[string]string {
        "query": fmt.Sprintf(
			`{
				react(where: {liker: {_eq: "%s"}, _and: {}, post: {_eq: "%s"}}) {
					liker
					post
				  }
			}	
            `,
		a, b),
    }

	return jsonData
}


func InsertReact(a, b string) map[string]string {
	jsonInsert := map[string]string{
		"query": fmt.Sprintf(
			`
			mutation MyMutation {
				insert_react_one(object: {liker: "%s", post: "%s"})
				{
					reactid
				}
			  }			  
			`,
		a,b),
          
    }

	return jsonInsert
} 


func DeleteReact(a, b string) map[string]string {
	jsonInsert := map[string]string{
		"query": fmt.Sprintf(
			`
			mutation MyMutation {
				delete_react(where: {liker: {_eq: "%s"}, _and: {post: {_eq: "%s"}}}) {
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


func InsertUser(a, b, c string) map[string]string {
	jsonInsert := map[string]string{
		"query": fmt.Sprintf(
			`
			mutation MyMutation {
				insert_ragister_one(object: {username: "%s", email: "%s", password: "%s", profilepic: "nil"})
				{
					usersid
				}
			  }			  
			`,
		a,b,c),
          
    }

	return jsonInsert
} 


func FindPost ()map[string]string {
	jsonData := map[string]string {
		"query" : `{
			challenge {
			  question
			  language
			  visibility
			  code
			  answer
			  id
			  usersid
			  date
			  ragister_details {
				username
			  }
			}
		  }`,
	}

	return jsonData
	  
}


func InsertPost(a, b, c, d, e, f string) map[string]string {
	// res, _ := strconv.Atoi(usersid)
	jsonInsert := map[string]string{
		"query": fmt.Sprintf(
			`
			mutation MyMutation {
				insert_challenge_one(object: {answer: "%s", visibility: "%s", usersid: "%s", code: "%s", language: "%s", question: "%s"})
				{
				  id
				}
			  }	  
			`,
			a, b, c, d, e, f ),
          
    }

	return jsonInsert
} 





// func FindPostById ( text string)map[string]string {
// 	jsonData := map[string]string {
//         "query": fmt.Sprintf(
// 			`{
// 				challenge(where: {username: {_eq: "%s"}}) {
// 					email
// 					regdate
// 					username
// 					usersid
// 				  }
// 			}	
//             `,
// 		a, b),
//     }
// 	jsonData := map[string]string {
// 		"query" : `{
// 			challenge {
// 			  duration
// 			  medal
// 			  question
// 			  answer
// 			  id
// 			  date
// 			  ragister_details {
// 				username
// 			  }
// 			}
// 		  }`,
// 	}

// 	return jsonData
	  
// }



func AllUsersQuery ()map[string]string {
	jsonData := map[string]string {
		"query" : `{
			ragister {
				username
				email
			}
		  }`,
	}

	return jsonData
	  
}