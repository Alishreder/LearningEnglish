package main

import (
	. "LearningEnglish/algorithms"
	. "LearningEnglish/data"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ai "github.com/night-codes/mgo-ai"
	"gopkg.in/mgo.v2"
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
	group := router.Group("/home")
	group.POST("/addNewWord", AddNewWord)
	group.POST("/deleteWord", DeleteWord)
	group.POST("/addToLearnList/:id", AddToLearnList)
	group.GET("/showLearnList", ShowLearnList)

	group.GET("/learn", Learn)

	group.POST("/checkFirstAlg", CheckFirstAlg)
	group.POST("/checkSecondAlg", CheckSecondAlg)
	group.POST("/checkThirdAlg", CheckThirdAlg)

	router.GET("/showUsersList", ShowUsersList)
	group.GET("/showUsersDictionary", ShowUsersDictionary)

	err = router.Run(":8080")
	if err != nil {
		fmt.Println("Cant listen and serve on 0.0.0.0:8080")
		return
	}
}
