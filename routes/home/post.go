package home

import (
	"encoding/json"
	"hasuragraphql"
	"html/template"
	"fmt"
	"strings"
	"strconv"
	// "fmt"
)

type HomeStruct struct {
	Details template.HTML
	Username string
	User string
	RelatedPost template.HTML
}

func (l *HomeStruct) HomeDetails (name string){
	l.User = name[:1]
	l.Username = name
}

func ShowView(name, usersid string) string {
	result := hasuragraphql.HasuraGraphQlApi(hasuragraphql.FindPost())
	var x hasuragraphql.PostStruct
	json.Unmarshal([]byte(result), &x)

	postresult := []string{}

	if len(x.Data.Challenge) > 0 {
		for _, item := range x.Data.Challenge {


			// fmt.Println(item.RagisterDetails.Username)

			if name == "loggedin" {
				postresult  = append(postresult, "<a style='text-decoration:none;color:black;' href='/story/?default="+strconv.Itoa(item.ID)+"'/>")
			}
			postresult  = append(postresult, "<div class='challenge-show'><div class='child-challenge-show'>")
			postresult  = append(postresult, "<div class='top-details'>")

			//details
			postresult  = append(postresult, "<div id='pic'><button>"+item.RagisterDetails.Username[0:1]+"</button></div>")
			postresult  = append(postresult, "<div id='name'><span>"+item.RagisterDetails.Username+"</span></div>")
			postresult  = append(postresult, "<div id='date'><span>"+hasuragraphql.GetDate(item.Date)+"</span></div>")
			//details

			//top
			postresult  = append(postresult, "</div>")

			postresult  = append(postresult, "<div class='body'>")
			postresult  = append(postresult, "<div id='question'><h3>"+item.Question+"</h3></div>")
			postresult  = append(postresult, "<div id='answer'>"+item.Answer+"</div>")
			postresult  = append(postresult, "</div>")

			//do it in a different way challenge

			postresult  = append(postresult, "<div class='challenge-title'>")
			postresult  = append(postresult, "<div id='challenge-title'><span><img src='/assets/icons/smile.svg' /> <i>How to do it in a different way challenge</i></span></div>")
			postresult  = append(postresult, "</div>")

			postresult  = append(postresult, "<div class='reaction'>")
			//reaction
			postresult  = append(postresult, "<div class='reaction-btn'>")
			postresult  = append(postresult, "<div id='reaction-btn'>"+GetLike(usersid, strconv.Itoa(item.ID))+"</div>")
			postresult  = append(postresult, "<div id='reaction-btn'><button><img src='/assets/icons/bookmarkbl.svg' /></button></div>")
			postresult  = append(postresult, "<div id='reaction-btn'><button style='border:1px solid #ccc;height:25px;'>Public</button></div>")
			postresult  = append(postresult, "<div id='reaction-btn'><button style='border:1px solid #ccc;height:25px;'><img style='height:18px;width:18px;' src='/assets/icons/gift.svg' /></button></div>")
			postresult  = append(postresult, "</div>")

			//join
			postresult  = append(postresult, "<div id='reaction-join'><button>Accept Challenge</button></div>")


			postresult  = append(postresult, "</div>")

			postresult  = append(postresult, "</div></div>")

			if name == "loggedin" {
				postresult  = append(postresult, "</a>")
			}

		}
	}

	return strings.Join(postresult, "")
}



func SinglePost(id, usersid string) string {
	result := hasuragraphql.HasuraGraphQlApi(hasuragraphql.FindPost())
	var x hasuragraphql.PostStruct
	json.Unmarshal([]byte(result), &x)

	postresult := []string{}

	if len(x.Data.Challenge) > 0 {
		for _, item := range x.Data.Challenge {


			// fmt.Println(item.RagisterDetails.Username)

			if id == strconv.Itoa(item.ID) {
				postresult  = append(postresult, "<div class='challenge-show'><div class='child-challenge-show'>")
			    postresult  = append(postresult, "<div class='top-details'>")

			    //details
			    postresult  = append(postresult, "<div id='pic'><button>"+item.RagisterDetails.Username[0:1]+"</button></div>")
			    postresult  = append(postresult, "<div id='name'><span>"+item.RagisterDetails.Username+"</span></div>")
			    postresult  = append(postresult, "<div id='date'><span>"+hasuragraphql.GetDate(item.Date)+"</span></div>")
			    //details

			    //top
			    postresult  = append(postresult, "</div>")

			    postresult  = append(postresult, "<div class='body'>")
			    postresult  = append(postresult, "<div id='question'><h3>"+item.Question+"</h3></div>")
			    postresult  = append(postresult, "<div id='answer'>"+item.Answer+"</div>")
			    postresult  = append(postresult, "</div>")

				postresult  = append(postresult, "<div style='padding-top:10px;' class='code'>")
				postresult  = append(postresult, MoldVariable(item.Code))
				postresult  = append(postresult, "</div>")

			    //do it in a different way challenge

			    postresult  = append(postresult, "<div class='challenge-title'>")
			    postresult  = append(postresult, "<div id='challenge-title'><span><img src='/assets/icons/smile.svg' /> <i>How to do it in a different way challenge</i></span></div>")
			    postresult  = append(postresult, "</div>")

			    postresult  = append(postresult, "<div class='reaction'>")
			    //reaction
			    postresult  = append(postresult, "<div class='reaction-btn'>")
			    postresult  = append(postresult, "<div id='reaction-btn'><span id='like"+strconv.Itoa(item.ID)+"'>"+GetLike(usersid, strconv.Itoa(item.ID))+"</span></div>")
			    postresult  = append(postresult, "<div id='reaction-btn'><button><img src='/assets/icons/bookmarkbl.svg' /></button></div>")
			    postresult  = append(postresult, "<div id='reaction-btn'><button style='border:1px solid #ccc;height:25px;'>Public</button></div>")
			    postresult  = append(postresult, "<div id='reaction-btn'><button style='border:1px solid #ccc;height:25px;'><img style='height:18px;width:18px;' src='/assets/icons/gift.svg' /></button></div>")
			    postresult  = append(postresult, "</div>")

			    //join
				// /in/accept/

				fmt.Println(strconv.Itoa(item.Usersid), " POST")

				if strconv.Itoa(item.Usersid) != usersid {
					postresult  = append(postresult, "<form method='post' action=' /in/accept/'><input style='display:none;' type='text' name='postid' value='"+strconv.Itoa(item.ID)+"'/>")

				    verfys, res := ChallengeAcceptCheck(usersid, strconv.Itoa(item.ID))
			        postresult  = append(postresult, res)
				    postresult  = append(postresult, "</form>")

			        postresult  = append(postresult, "</div>")

					verfyspublish, publishres := AllPublish(strconv.Itoa(item.ID), usersid, item.Question)

				    if verfyspublish == false {
						if verfys == true {
							postresult  = append(postresult, "<form method='post' action='/publish/accept/'>")
					        postresult  = append(postresult, "<div class='start-write'><div id='start-write'><span><img src='/assets/icons/pen.svg' /> Start Writing</span></div><div class='child-start-write'>")
					        postresult  = append(postresult, "<div id='text-area'><textarea id='writexdetails' name='story' placeholder='Write ups'></textarea></div>")
					        postresult  = append(postresult, "<div id='code-text'><input type='text' name='code' id='codexdetails' placeholder='Code goes here'/></div>")
					        postresult  = append(postresult, "<input style='display:none;' type='text' name='postid' value='"+strconv.Itoa(item.ID)+"'/>")
					        postresult  = append(postresult, "<div style='display:none;' id='publish'><button>Publish Challenge</button></div>")
					        postresult  = append(postresult, "</div>")
					        postresult  = append(postresult, "</form>")
						}
				    }
					
					postresult = append(postresult, publishres)
				}

				//reaction 
				postresult = append(postresult, "</div>")

				
				// postresult  = append(postresult, "<div class='comments'><button><img src='/assets/icons/comments.svg' /> <span><i>View/Add Comments</i></span> </button><div class='child-comment'>")
				// postresult  = append(postresult, "<div id='add'><textarea name='comment' placeholder='Add a Comment' id='add-comment'></textarea><button style='display:none;' id='submit-comment'>Comment</button></div>")

				postresult  = append(postresult, "</div></div>")

			    postresult  = append(postresult, "</div></div>")
			}
		}
	}
	return strings.Join(postresult, "")
}




func ShowRelated(name, usersid, postid string) string {
	result := hasuragraphql.HasuraGraphQlApi(hasuragraphql.FindPost())
	var x hasuragraphql.PostStruct
	json.Unmarshal([]byte(result), &x)

	postresult := []string{}

	if len(x.Data.Challenge) > 0 {
		for _, item := range x.Data.Challenge {


			// fmt.Println(item.RagisterDetails.Username)

			if strconv.Itoa(item.ID) != postid {
				if name == "loggedin" {
					postresult  = append(postresult, "<a style='text-decoration:none;color:black;' href='/story/?default="+strconv.Itoa(item.ID)+"'/>")
				}
				postresult  = append(postresult, "<div class='challenge-show'><div class='child-challenge-show'>")
				postresult  = append(postresult, "<div class='top-details'>")
	
				//details
				postresult  = append(postresult, "<div id='pic'><button>"+item.RagisterDetails.Username[0:1]+"</button></div>")
				postresult  = append(postresult, "<div id='name'><span>"+item.RagisterDetails.Username+"</span></div>")
				postresult  = append(postresult, "<div id='date'><span>"+hasuragraphql.GetDate(item.Date)+"</span></div>")
				//details
	
				//top
				postresult  = append(postresult, "</div>")
	
				postresult  = append(postresult, "<div class='body'>")
				postresult  = append(postresult, "<div id='question'><h3>"+item.Question+"</h3></div>")
				postresult  = append(postresult, "<div id='answer'>"+item.Answer+"</div>")
				postresult  = append(postresult, "</div>")
	
				//do it in a different way challenge
	
				postresult  = append(postresult, "<div class='challenge-title'>")
				postresult  = append(postresult, "<div id='challenge-title'><span><img src='/assets/icons/smile.svg' /> <i>How to do it in a different way challenge</i></span></div>")
				postresult  = append(postresult, "</div>")
	
				postresult  = append(postresult, "<div class='reaction'>")
				//reaction
				postresult  = append(postresult, "<div class='reaction-btn'>")
				postresult  = append(postresult, "<div id='reaction-btn'>"+GetLike(usersid, strconv.Itoa(item.ID))+"</div>")
				postresult  = append(postresult, "<div id='reaction-btn'><button><img src='/assets/icons/bookmarkbl.svg' /></button></div>")
				postresult  = append(postresult, "<div id='reaction-btn'><button style='border:1px solid #ccc;height:25px;'>Public</button></div>")
				postresult  = append(postresult, "<div id='reaction-btn'><button style='border:1px solid #ccc;height:25px;'><img style='height:18px;width:18px;' src='/assets/icons/gift.svg' /></button></div>")
				postresult  = append(postresult, "</div>")
	
				//join
				postresult  = append(postresult, "<div id='reaction-join'><button>Accept Challenge</button></div>")
	
	
				postresult  = append(postresult, "</div>")
	
				postresult  = append(postresult, "</div></div>")
	
				if name == "loggedin" {
					postresult  = append(postresult, "</a>")
				}
			}

		}
	}

	return "<div class='related-post'><h2>Related Challenges   <img src='/assets/icons/related.svg' /></h2></div>"+strings.Join(postresult, "")
}





func UserChallenge(name, usersid string) string {
	result := hasuragraphql.HasuraGraphQlApi(hasuragraphql.FindPost())
	var x hasuragraphql.PostStruct
	json.Unmarshal([]byte(result), &x)

	postresult := []string{}

	if len(x.Data.Challenge) > 0 {
		for _, item := range x.Data.Challenge {


			// fmt.Println(item.RagisterDetails.Username)

			if strconv.Itoa(item.Usersid) == usersid {
				if name == "loggedin" {
					postresult  = append(postresult, "<a style='text-decoration:none;color:black;' href='/story/?default="+strconv.Itoa(item.ID)+"'/>")
				}
				postresult  = append(postresult, "<div class='challenge-show'><div class='child-challenge-show'>")
				postresult  = append(postresult, "<div class='top-details'>")
	
				//details
				postresult  = append(postresult, "<div id='pic'><button>"+item.RagisterDetails.Username[0:1]+"</button></div>")
				postresult  = append(postresult, "<div id='name'><span>"+item.RagisterDetails.Username+"</span></div>")
				postresult  = append(postresult, "<div id='date'><span>"+hasuragraphql.GetDate(item.Date)+"</span></div>")
				//details
	
				//top
				postresult  = append(postresult, "</div>")
	
				postresult  = append(postresult, "<div class='body'>")
				postresult  = append(postresult, "<div id='question'><h3>"+item.Question+"</h3></div>")
				postresult  = append(postresult, "<div id='answer'>"+item.Answer+"</div>")
				postresult  = append(postresult, "</div>")
	
				//do it in a different way challenge
	
				postresult  = append(postresult, "<div class='challenge-title'>")
				postresult  = append(postresult, "<div id='challenge-title'><span><img src='/assets/icons/smile.svg' /> <i>How to do it in a different way challenge</i></span></div>")
				postresult  = append(postresult, "</div>")
	
				postresult  = append(postresult, "<div class='reaction'>")
				//reaction
				postresult  = append(postresult, "<div class='reaction-btn'>")
				postresult  = append(postresult, "<div id='reaction-btn'>"+GetLike(usersid, strconv.Itoa(item.ID))+"</div>")
				postresult  = append(postresult, "<div id='reaction-btn'><button><img src='/assets/icons/bookmarkbl.svg' /></button></div>")
				postresult  = append(postresult, "<div id='reaction-btn'><button style='border:1px solid #ccc;height:25px;'>Public</button></div>")
				postresult  = append(postresult, "<div id='reaction-btn'><button style='border:1px solid #ccc;height:25px;'><img style='height:18px;width:18px;' src='/assets/icons/gift.svg' /></button></div>")
				postresult  = append(postresult, "</div>")
	
				//join
				postresult  = append(postresult, "<div id='reaction-join'><button>Accept Challenge</button></div>")
	
	
				postresult  = append(postresult, "</div>")
	
				postresult  = append(postresult, "</div></div>")
	
				if name == "loggedin" {
					postresult  = append(postresult, "</a>")
				}
			}

		}
	}

	if len(postresult) > 0 {
		return "<div class='related-post'><h2>Challenges   <img src='/assets/icons/related.svg' /></h2></div>"+strings.Join(postresult, "")
	}else{
		return "<div class='related-post'><h2>No Challenge To Show <img src='/assets/icons/related.svg' /></h3></div><div class='start-new'><a href='/make/challenge/' style='text-decoration: none;color:black;'><button>Start a New Challenge</button></a></div>"
	}
}


// FindAccept(usersid string)



func AcceptedChallenge(name, usersid string) string {
	result := hasuragraphql.HasuraGraphQlApi(hasuragraphql.FindPost())
	var x hasuragraphql.PostStruct
	json.Unmarshal([]byte(result), &x)

	postresult := []string{}

	accepts := FindAccept(usersid)

	if len(x.Data.Challenge) > 0 {
		for _, item := range x.Data.Challenge {


			// fmt.Println(item.RagisterDetails.Username)

			if accepts[item.ID] == true{
				if name == "loggedin" {
					postresult  = append(postresult, "<a style='text-decoration:none;color:black;' href='/story/?default="+strconv.Itoa(item.ID)+"'/>")
				}
				postresult  = append(postresult, "<div class='challenge-show'><div class='child-challenge-show'>")
				postresult  = append(postresult, "<div class='top-details'>")
	
				//details
				postresult  = append(postresult, "<div id='pic'><button>"+item.RagisterDetails.Username[0:1]+"</button></div>")
				postresult  = append(postresult, "<div id='name'><span>"+item.RagisterDetails.Username+"</span></div>")
				postresult  = append(postresult, "<div id='date'><span>"+hasuragraphql.GetDate(item.Date)+"</span></div>")
				//details
	
				//top
				postresult  = append(postresult, "</div>")
	
				postresult  = append(postresult, "<div class='body'>")
				postresult  = append(postresult, "<div id='question'><h3>"+item.Question+"</h3></div>")
				postresult  = append(postresult, "<div id='answer'>"+item.Answer+"</div>")
				postresult  = append(postresult, "</div>")
	
				//do it in a different way challenge
	
				postresult  = append(postresult, "<div class='challenge-title'>")
				postresult  = append(postresult, "<div id='challenge-title'><span><img src='/assets/icons/smile.svg' /> <i>How to do it in a different way challenge</i></span></div>")
				postresult  = append(postresult, "</div>")
	
				postresult  = append(postresult, "<div class='reaction'>")
				//reaction
				postresult  = append(postresult, "<div class='reaction-btn'>")
				postresult  = append(postresult, "<div id='reaction-btn'>"+GetLike(usersid, strconv.Itoa(item.ID))+"</div>")
				postresult  = append(postresult, "<div id='reaction-btn'><button><img src='/assets/icons/bookmarkbl.svg' /></button></div>")
				postresult  = append(postresult, "<div id='reaction-btn'><button style='border:1px solid #ccc;height:25px;'>Public</button></div>")
				postresult  = append(postresult, "<div id='reaction-btn'><button style='border:1px solid #ccc;height:25px;'><img style='height:18px;width:18px;' src='/assets/icons/gift.svg' /></button></div>")
				postresult  = append(postresult, "</div>")
	
				//join
				postresult  = append(postresult, "<div id='reaction-join'><button>Accept Challenge</button></div>")
	
	
				postresult  = append(postresult, "</div>")
	
				postresult  = append(postresult, "</div></div>")
	
				if name == "loggedin" {
					postresult  = append(postresult, "</a>")
				}
			}

		}
	}

	if len(postresult) > 0 {
		return "<div class='related-post'><h2>Accepted Challenge   <img src='/assets/icons/related.svg' /></h2></div>"+strings.Join(postresult, "")
	}else{
		return "<div class='related-post'><h2>No Challenge Accepted <img src='/assets/icons/related.svg' /></h3></div><div class='start-new'><a href='/' style='text-decoration: none;color:black;'><button>Explore Challenges</button></a></div>"
	}
}