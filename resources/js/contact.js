$(document).ready(function(){
    $(document).on('submit', '.contact-form', function() {
        $('.contact-form').fadeOut(function() {
            $('.contact-sent').delay(500).fadeIn();
        });
        return false;
    })
});
