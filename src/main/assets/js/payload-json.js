$("#user-form").on("submit", function (e) {
    e.preventDefault()

    var $self = $(this)
    var payload = JSON.stringify({
        name: $('[name="name"]').val(),
        age: parseInt($('[name="age"]').val(), 10),
        gender: $('[name="gender"]').val()
    });

    $.ajax({
        url: $self.attr("action"),
        type: $self.attr("method"),
        data: payload,
        contentType: 'application/json',
    }).then(function (res) {
        $(".message").text(res);
    }).catch(function (a) {
        alert("ERROR: " + a.responseText);
    });

});