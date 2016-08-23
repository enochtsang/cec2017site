$(document).ready(function(){
	var currentSection = $(".competitions-general")
	if(window.location.search.slice(1)) {
		var currentSection = $("." + window.location.search.slice(1));
		if(!currentSection.length) {
			currentSection = $(".competitions-general")
			console.log("doesn't exist");
		}
	}

	currentSection.fadeIn();

	$(".competitions-nav li").click(function() {
		var newSection = $("." + $(this).attr("value"));
		currentSection.fadeOut(function() {
			newSection.fadeIn();
			currentSection = newSection;
		});
	})
});
