$(document).ready(function(){
    // Fade in drop downs on nav-bar
    $(".nav-bar li").hover(function() {
        $(this).children(".nav-dropdown").fadeIn("fast");
    }, function() {
        $(this).children(".nav-dropdown").fadeOut("fast");
    });

    // Highlight current tab in nav-bar
    var loc = window.location.pathname;
    var dir = loc.substring(1, loc.length) + "/";
    dir = dir.substring(0, dir.indexOf("/"));
    $(".nav-bar a[href=\"/" + dir + "\"]").css("color", "#F20");

    $(".mobile-nav-controller").click(function() {
        $(".mobile-nav-bar ul").slideToggle();
        if($(".mobile-nav-controller img").css("transform") == 'none'){
            $(".mobile-nav-controller img").css("transform", "rotate(180deg)");
        } else {
            $(".mobile-nav-controller img").css("transform", "");
        }
    });
});
