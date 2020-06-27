package algorithms

import (
	. "dictionaryProject/data"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddNewWord(c *gin.Context) {
	word := c.PostForm("word")
	translate := c.PostForm("translate")
	if word != "" || translate != "" {
		id := CurrentUser.LastWordId + 1
		CurrentUser.LastWordId++
		CurrentUser.Dictionary = append(CurrentUser.Dictionary, WordTranslate{Word: word, Translate: translate, WordId: id, WantToLearn: true})
		err := Collection.Update(Obj{"_id": CurrId}, CurrentUser)
		if err != nil {
			c.String(http.StatusInternalServerError, "AddNewWord: fail while trying to add new word to dictionary: ", err)
		}
		fmt.Println(CurrentUser, CurrId)
	}

	c.Status(http.StatusOK)
}

func DeleteWord(c *gin.Context) {
	idInt, _ := strconv.Atoi(c.Param("id"))
	idToDelete := uint64(idInt)
	for i, v := range CurrentUser.Dictionary {
		if v.WordId == idToDelete {
			CurrentUser.Dictionary = append(CurrentUser.Dictionary[:i], CurrentUser.Dictionary[i+1:]...)
			err := Collection.Update(Obj{"_id": CurrId}, CurrentUser)
			if err != nil {
				c.String(http.StatusInternalServerError, "DeleteWord: fail while trying to delete word: ", err)
			}
		}
	}

	c.Status(http.StatusOK)
}

func AddToLearnList(c *gin.Context) {
	idInt, _ := strconv.Atoi(c.Param("id"))
	idToLearn := uint64(idInt)
	for i, v := range CurrentUser.Dictionary {
		if v.WordId == idToLearn {
			CurrentUser.Dictionary[i].WantToLearn = true
		}
	}

	c.Status(http.StatusOK)
}

// func Learn(c *gin.Context) {
// 	CurrentUser
// }


/*

It was in main

var MyCache = CacheType{
	Cache: make(map[uint64]Word),
}
MyCache.AddWordToCache("слово", "word")
MyCache.AddWordToCache("имя", "name")
MyCache.AddWordToCache("фамилия", "surname")

for _, v := range MyCache.Cache {
fmt.Printf("%+v\n", v)
}

*/

// func FromEngToRus(dictionary []WordTranslate) {
// 	for i := 0; i < len(dictionary); i++ {
//
// 	}
// }
