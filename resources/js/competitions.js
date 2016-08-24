$(document).ready(function(){
	var currentSection = $(".competitions-general")
	if(window.location.search.slice(1)) {
		var currentSection = $("." + window.location.search.slice(1));
		if(!currentSection.length) {
			currentSection = $(".competitions-general")
			console.log("doesn't exist");
		}
	}

	var maxHeight = Math.max.apply(null, $(".competition-info").map(function() {
		return $(this).height();
	}).get());
	$(".competition-info-container").css("min-height", maxHeight);

	currentSection.css("opacity", "1");

	$(".competitions-nav li").click(function() {
		var newSection = $("." + $(this).attr("value"));
		currentSection.css("opacity", "0");
		setTimeout(function() {
			newSection.css("opacity", "1");
			currentSection = newSection;
		}, 400);
	});
});
