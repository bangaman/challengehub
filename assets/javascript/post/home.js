document.addEventListener("keyup", function(event){
   if(event.target.id === "writexdetails") {
      if(event.target.value.length > 0 && document.getElementById("codexdetails").value.length > 0) {
         document.getElementById("publish").style.display = "block"
         
      }else{
         document.getElementById("publish").style.display = "none"
      }
   }else if(event.target.id === "codexdetails") {
       if(event.target.value.length > 0 && document.getElementById("writexdetails").value.length > 0) {
         document.getElementById("publish").style.display = "block"
         
      }else{
         document.getElementById("publish").style.display = "none"
      }

      
   }else if(event.target.id === "add-comment") {
       if(event.target.value.length > 0) {
         document.getElementById("submit-comment").style.display = "block"
         
      }
  }
})
//trash


$(document).ready(function(){
    $(document).click(function(event){
        if(event.target.id === "trash") {
         $("#main-delete").toggle()
     }
   
})})
