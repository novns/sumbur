$(".articles").on("click", ".article-edit", function () {
    var id = $(this).attr("id").split("-")[2];

    if ($("#article-form-" + id).length) {
        $("#article-form-" + id).remove();
    } else {
        $.get("/edit/" + id, function (data) {
            $("#article-" + id).html(data);
            $((id == 0 ? "#url-" : "#title-") + id).focus();
        });
    }
});


$(".articles").on("click", ".cancel-edit", function () {
    var id = $(this).attr("id").split("-")[1];

    $("#article-form-" + id).remove();
});


$(".articles").on("keyup", ".article-form", function (event) {
    if (event.keyCode == 27)
        $(this).remove();
});


$(".articles").on("submit", "form.article-form", function () {
    var id = $(this).attr("id").split("-")[2];

    $("#article-form-" + id + " :input").each(function () {
        $(this).val($.trim($(this).val()));
    });

    $.post("/edit/" + id,
        $("#article-form-" + id).serialize(),
        function (data) {
            if (id == 0) {
                location.reload();
            } else {
                $("#article-" + id).html(data);
                $("#title-" + id).focus()

                $.get("/tags/" + $("#tags").data("stag"), function (data) {
                    $("#tags").replaceWith(data);
                });
            }
        }
    );

    return false;
});
