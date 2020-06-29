(function ($) {
    $(document).ready(function () {
        $(".learnWords").click(function () {
            $.ajax({
                url: "/learn",
                type: "GET",
                success: function () {
                    document.location.href = "http://localhost:8080/learn"
                },
                error: function () {
                    alert("err");
                }
            })
        })

        $(".checkFirstAlg").click(function () {
            const id = $(this).attr("id");
            let translate = $("#"+id).serializeArray()[0].value
            $.ajax({
                url: "/checkFirstAlg",
                data: {id: id, translate: translate},
                type: "POST",
                success: function () {
                    alert("ok");
                },
                error: function () {
                    alert("wrong translate")
                }
            })
        })

        $(".checkSecondAlg").click(function () {
            const id = $(this).attr("id");
            let word = $("#"+id).serializeArray()[0].value
            $.ajax({
                url: "/checkSecondAlg",
                data: {id: id, word: word},
                type: "POST",
                success: function () {
                    alert("ok");
                },
                error: function () {
                    alert("wrong word"+ word)
                }
            })
        })

        $(".checkThirdAlg").click(function () {
            const id = $(this).attr("id");
            let sentence = $("#"+id).serializeArray()[0].value
            $.ajax({
                url: "/checkThirdAlg",
                data: {id: id, sentence: sentence},
                type: "POST",
                success: function () {
                    alert("ok");
                },
                error: function () {
                    alert("you should write something"+ sentence)
                }
            })
            $(".third").show()
        })

        $("#first").click(function () {
            $(".first").show()
            $(".second").hide()
            $(".third").hide()
        })
        $("#second").click(function () {
            $(".first").hide()
            $(".second").show()
            $(".third").hide()
        })
        $("#third").click(function () {
            $(".first").hide()
            $(".second").hide()
            $(".third").show()
        })
    })
})(jQuery);