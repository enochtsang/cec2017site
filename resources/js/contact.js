$(document).ready(function(){

    $(document).on('click', 'input[type=submit]', function() {
        var goodInput = true;
        var inputs = [
            $(".contact-form input[name=name]"),
            $(".contact-form input[name=organization]"),
            $(".contact-form input[name=return-email]"),
            $(".contact-form textarea[name=message]")
        ];

        //Check honeyput
        if($(".contact-form input[name=LeaveBlank]").val() !== "") {
            goodInput = false;
        }

        // Check blank fields
        for(var i = 0; i < inputs.length; ++i) {
            if(inputs[i].val() === "") {
                inputs[i].css("background-color", "#FCC");
                $('.contact-error').fadeOut(function() {
                    $('.contact-error').html('Fill in all fields').fadeIn();
                });
                goodInput = false;
            } else {
                inputs[i].css("background-color", "white");
            }
        }

        // Check valid email
        if(goodInput) {
            if(!validateEmail(inputs[2].val())) {
                inputs[2].css("background-color", "#FCC");
                $('.contact-error').fadeOut(function() {
                    $('.contact-error').html('Invalid email address').fadeIn();
                });
                goodInput = false;
            } else {
                inputs[2].css("background-color", "white");
            }
        }

        // Check captcha filled out
        var captchaResponse = grecaptcha.getResponse()
        console.log(captchaResponse)
        if(captchaResponse.length == 0) {
            goodInput = false;
            $('.contact-error').fadeOut(function() {
                $('.contact-error').html('Captcha not completed').fadeIn();
            });
        }

        if(!goodInput) {
            return false;
        }

        for(var i = 0; i < inputs.length; ++i) {
            inputs[i].prop('disabled', true);
            inputs[i].css('background-color', '#EEE');
            inputs[i].css('color', '#888');
        }
        $.post('/contact', {
            name: inputs[0].val(),
            organization: inputs[1].val(),
            returnEmail: inputs[2].val(),
            message: inputs[3].val(),
            captcha: captchaResponse
        });

        $('.contact-form .submit-area').fadeOut(function() {
            $('.contact-sent').fadeIn();
        });
        return false;
    })
});

function validateEmail(email) {
    var re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    return re.test(email);
}

// render google recaptcha
var onloadCallback = function() {
    grecaptcha.render('g-recaptcha', {
        'sitekey' : '6Lce7SYTAAAAAIrX7phaDRlvkkJDzmzHWcGhS1GN',
        // `data-size`:
        'data-size' : `compact`
    });
};
