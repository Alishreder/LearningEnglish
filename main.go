package main

import (
	. "dictionaryProject/algorithms"
	. "dictionaryProject/data"
	"fmt"
	"github.com/gin-gonic/gin"
	ai "github.com/night-codes/mgo-ai"
	"gopkg.in/mgo.v2"
)

// func addUser() {
// 	user := User{Login: "lol@gmail.com", Pass: "1"}
// 	Users[user] = struct{}{}
// }

func main() {
	var err error
	Session, err = mgo.Dial("mongodb://127.0.0.1")
	if err != nil {
		panic(err)
	}
	DB = Session.DB("Users")
	Collection = DB.C("user")
	ai.Connect(Collection)

	router := gin.Default()
	router.GET("/", Default)
	// addUser()

	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/registration", RegistrationGet)
	router.GET("/authorization", AuthorizationGet)

	router.POST("/registration", RegistrationPost)
	router.POST("/authorization", AuthorizationPost)

	router.GET("/home", Home)

	router.POST("/addNewWord", AddNewWord)
	router.POST("/deleteWord/:id", DeleteWord)
	router.POST("/addToLearnList/:id", AddToLearnList)




	err = router.Run(":8080")
	if err != nil {
		fmt.Println("Cant listen and serve on 0.0.0.0:8080")
		return
	}
}

