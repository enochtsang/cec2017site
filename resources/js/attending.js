$(document).ready(function(){

    $(".delegate-package-label").click(function() {
        $("#delegate-package-pdf").slideToggle();
        if($(".delegate-package-label .dropdown-arrow").css("transform") == 'none'){
            $(".delegate-package-label .dropdown-arrow").css("transform", "rotate(90deg)");
        } else {
            $(".delegate-package-label .dropdown-arrow").css("transform", "");
        }
    });

    $(".official-rulebook-label").click(function() {
        $("#official-rulebook-pdf").slideToggle();
        if($(".official-rulebook-label .dropdown-arrow").css("transform") == 'none'){
            $(".official-rulebook-label .dropdown-arrow").css("transform", "rotate(90deg)");
        } else {
            $(".official-rulebook-label .dropdown-arrow").css("transform", "");
        }
    });

});
