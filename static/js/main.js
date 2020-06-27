(function ($) {
    $(document).ready(function () {
        $(".addWord").click(function () {
            $.ajax({
                url: "/addNewWord",
                type: "POST",
                success: function () {
                    document.location.href = "http://localhost:8080/home"
                },
                error: function () {
                    alert("error")
                }
            }).done(function () {
                console.log("successfully added");
            }).fail(function (response) {
                console.log(response.responseText)
            });
        })

        $(".deleteWord").click(function () {
            const id = $(this).attr("id")
            $.ajax({
                url: "/deleteWord/" + id,
                type: "POST",
                success: function () {
                    document.location.href = "http://localhost:8080/home"
                },
                error: function () {
                    alert("err");
                }
            })
        });

        $(".addToLearnList").click(function () {
            const id = $(this).attr("id")
            $.ajax({
                url: "/addToLearnList/" + id,
                type: "POST",
                success: function () {
                    document.location.href = "http://localhost:8080/home"
                },
                error: function () {
                    alert("err");
                }
            })

        })
    })
})(jQuery);

