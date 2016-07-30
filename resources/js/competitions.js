$(document).ready(function(){
    var currentSection = $("." + window.location.search.slice(1));
    currentSection.fadeIn();

    $(".competitions-nav li").click(function() {
        var newSection = $("." + $(this).attr("value"));
        currentSection.fadeOut(function() {
            newSection.fadeIn();
            currentSection = newSection;
        });
    })
});
