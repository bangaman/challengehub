package home



func FindPublish()map[string]string {
	jsonData := map[string]string {
		"query" : `{
			publish {
			   story
			   code
			   usersid
			   postid
			   publishid
			  ragister_publish {
				username
			  }
			}
		  }`,
	}

	return jsonData
	  
}