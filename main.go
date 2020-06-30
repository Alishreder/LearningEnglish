package main

import (
	. "dictionaryProject/algorithms"
	. "dictionaryProject/data"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ai "github.com/night-codes/mgo-ai"
	"gopkg.in/mgo.v2"
	"log"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

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
	router.POST("/deleteWord", DeleteWord)
	router.POST("/addToLearnList/:id", AddToLearnList)
	router.GET("/showLearnList", ShowLearnList)

	router.GET("/learn", Learn)

	router.POST("/checkFirstAlg", CheckFirstAlg)
	router.POST("/checkSecondAlg", CheckSecondAlg)
	router.POST("/checkThirdAlg", CheckThirdAlg)

	router.GET("/showUsersList", ShowUsersList)
	router.GET("/showUsersDictionary", ShowUsersDictionary)


	err = router.Run(":8080")
	if err != nil {
		fmt.Println("Cant listen and serve on 0.0.0.0:8080")
		return
	}
}

