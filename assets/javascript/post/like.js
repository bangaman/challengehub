
$(document).ready(function(){
    $(document).click(function(event){
        if(event.target.id === "btn-react"){
            $.post("/ajax/react/",
            {
                postid: event.target.getAttribute('data-post'),
             
            },
            function(data,status){
                if(status === "success"){
                    $("#like"+event.target.getAttribute('data-post')).html(data)
                    //alert(data)
                }
            });


           
        }else if(event.target.id === "suggestion-button"){
            $.post("/ajax/sub/",
            {
                spaceid: event.target.getAttribute('data-space'),
                owner: event.target.getAttribute('data-postuser')
             
            },
            function(data,status){
                if(status === "success"){
                    $("#span-suggestion"+event.target.getAttribute('data-space')).html(data)
                }
            });
            
            
        }
    })
})

