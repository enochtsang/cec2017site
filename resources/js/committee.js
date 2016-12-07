$(document).ready(function(){
	var currentDirector = "christine-cao";
    
    $(".directors-mini img").click(function() {
        nextDirector = $(this).attr("class");
        console.log("currentDirector is " + currentDirector);
        $(".directors-splash ." + currentDirector).fadeOut(function() {
            $(".directors-splash ." + nextDirector).css("display", "flex").hide().fadeIn("slow");
        });
        currentDirector = nextDirector;
        $("html, body").animate({ scrollTop: 0 }, "slow");
    });
});
