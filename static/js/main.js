(function ($) {
    $(document).ready(function () {
        $(".addWord").click(function () {
            $.ajax({
                url: "/home/addNewWord",
                type: "POST",
                success: function () {
                    document.location.href = "http://localhost:8080/home";
                },
                error: function () {
                    alert("error");
                }
            }).done(function () {
                console.log("successfully added");
            }).fail(function (response) {
                console.log(response.responseText);
            });
        });

        $(".deleteWord").click(function () {
            const id = $(this).attr("id");
            $.ajax({
                url: "/home/deleteWord",
                type: "POST",
                data: {id: id},
                success: function () {
                    location.reload();
                },
                error: function () {
                    alert("err");
                }
            });
        });

        $(".addToLearnList").click(function () {
            const id = $(this).attr("id");
            $.ajax({
                url: "/home/addToLearnList/" + id,
                type: "POST",
                success: function () {
                    location.reload();
                },
                error: function (data) {
                    alert("err"+data);
                }
            });
        });

        $("#showLearnList").click(function () {
            $.ajax({
                url: "/home/showLearnList",
                type: "GET",
                success: function () {
                    document.location.href = "http://localhost:8080/home/showLearnList";
                },
                error: function () {
                    alert("err");
                }
            });
        });

        $("#showHomePage").click(function () {
            $.ajax({
                url: "/home",
                type: "GET",
                success: function () {
                    document.location.href = "http://localhost:8080/home";
                },
                error: function () {
                    alert("err");
                }
            });
        });

        $("#showListOfUsers").click(function () {
            $.ajax({
                url: "/showUsersList",
                type: "GET",
                success: function () {
                    document.location.href = "http://localhost:8080/showUsersList";
                },
                error: function () {
                    alert("err");
                }
            });
        });

        $(".showUsersDictionary").click(function () {
            let id = $(this).attr("id");
            $.ajax({
                url: "/home/showUsersDictionary?id="+id,
                type: "GET",
                success: function () {
                    document.location.href = "http://localhost:8080/home/showUsersDictionary?id="+id;
                },
                error: function () {
                    alert("err");
                }
            });
        });
    })
})(jQuery);

