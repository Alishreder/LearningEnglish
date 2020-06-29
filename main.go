package main

import (
	. "dictionaryProject/algorithms"
	. "dictionaryProject/data"
	"fmt"
	"github.com/gin-gonic/gin"
	ai "github.com/night-codes/mgo-ai"
	"gopkg.in/mgo.v2"
)

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
	router.GET("/showLearnList", ShowLearnList)

	router.GET("/learn", Learn)

	router.POST("/checkFirstAlg", CheckFirstAlg)
	router.POST("/checkSecondAlg", CheckSecondAlg)
	router.POST("/checkThirdAlg", CheckThirdAlg)

	err = router.Run(":8080")
	if err != nil {
		fmt.Println("Cant listen and serve on 0.0.0.0:8080")
		return
	}
}

