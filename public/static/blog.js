$(".articles").on("click", ".article-edit", function () {
    var id = $(this).attr("id").split("-")[2];

    if ($("#article-form-" + id).length) {
        $("#article-form-" + id).remove();
    } else {
        $.get("/edit/" + id, function (data) {
            $("#article-" + id).html(data);
            $("#title-" + id).focus()
        });
    }
});


$(".articles").on("click", ".cancel-edit", function () {
    var id = $(this).attr("id").split("-")[1];

    $("#article-form-" + id).remove();
});
