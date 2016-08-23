$(document).ready(function(){
    $(document).on('click', 'input[type=submit]', function() {
        $('input[type=submit]').fadeOut(function() {
            $('.contact-sent').fadeIn();
        });
        return false;
    })
});
