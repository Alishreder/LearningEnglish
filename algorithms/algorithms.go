package algorithms

import (
	. "dictionaryProject/data"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetWordIndex(id uint64, dictionary []WordTranslate) (index int) {
	for i, v := range dictionary {
		if v.WordId == id {
			return i
		}
	}
	return
}

func CheckFirstAlg(c *gin.Context) {
	user, err := GetUserFromBD(c)
	if err != nil {
		panic(err)
	}
	userId, err := GetUserId(c)
	if err != nil {
		panic(err)
	}

	idInt, _ := strconv.Atoi(c.PostForm("id"))
	id := uint64(idInt)
	translate := c.PostForm("translate")

	userTry := User{}
	count, err := Collection.Find(Obj{"_id": userId, "dictionary.word_id": id}).Count()
	if err != nil {
		panic(err)
	}
	if count > 0 {
		_ = Collection.Find(Obj{"_id": userId, "dictionary.word_id": id}).One(&userTry)
		if user.Dictionary[GetWordIndex(id, userTry.Dictionary)].Translate == translate {
			c.Status(http.StatusOK)
			user.Dictionary[GetWordIndex(id, userTry.Dictionary)].FirstAlgorithm = false
			user.WordsForLearning[GetWordIndex(id, userTry.WordsForLearning)].FirstAlgorithm = false

			err := Collection.Update(Obj{"_id": userId}, user)
			if err != nil {
				c.String(http.StatusInternalServerError, "CheckFirstAlg: fail while trying to update words parameters: ", err)
			}
		} else {
			c.Status(http.StatusBadRequest)
		}
	}
	// c.Status(http.StatusOK)
}

func CheckSecondAlg(c *gin.Context) {
	user, err := GetUserFromBD(c)
	if err != nil {
		panic(err)
	}
	userId, err := GetUserId(c)
	if err != nil {
		panic(err)
	}

	idInt, _ := strconv.Atoi(c.PostForm("id"))
	id := uint64(idInt)
	word := c.PostForm("word")

	userTry := User{}
	count, err := Collection.Find(Obj{"_id": userId, "dictionary.word_id": id}).Count()
	if err != nil {
		panic(err)
	}
	if count > 0 {
		_ = Collection.Find(Obj{"_id": userId, "dictionary.word_id": id}).One(&userTry)
		if userTry.Dictionary[GetWordIndex(id, userTry.Dictionary)].Word == word {
			c.Status(http.StatusOK)
			user.Dictionary[GetWordIndex(id, userTry.Dictionary)].SecondAlgorithm = false
			user.WordsForLearning[GetWordIndex(id, userTry.WordsForLearning)].SecondAlgorithm = false

			err := Collection.Update(Obj{"_id": userId}, user)
			if err != nil {
				c.String(http.StatusInternalServerError, "CheckSecondAlg: fail while trying to update words parameters: ", err)
			}
		} else {
			c.Status(http.StatusBadRequest)
		}
	}
	c.Status(http.StatusOK)
}

func CheckThirdAlg(c *gin.Context) {
	user, err := GetUserFromBD(c)
	if err != nil {
		panic(err)
	}
	userId, err := GetUserId(c)
	if err != nil {
		panic(err)
	}

	idInt, _ := strconv.Atoi(c.PostForm("id"))
	id := uint64(idInt)
	sentence := c.PostForm("sentence")

	userTry := User{}
	count, err := Collection.Find(Obj{"_id": userId, "dictionary.word_id": id}).Count()
	if err != nil {
		panic(err)
	}
	if count > 0 {
		_ = Collection.Find(Obj{"_id": userId, "dictionary.word_id": id}).One(&userTry)
		if sentence != "" {
			user.Dictionary[GetWordIndex(id, userTry.Dictionary)].ThirdAlgorithm = false
			index := GetWordIndex(id, userTry.WordsForLearning)
			user.WordsForLearning[index].ThirdAlgorithm = false
			word := user.WordsForLearning[GetWordIndex(id, userTry.WordsForLearning)]
			if !word.FirstAlgorithm && !word.SecondAlgorithm && !word.ThirdAlgorithm {
				user.WordsForLearning = append(user.WordsForLearning[:index], user.WordsForLearning[index+1:]...)
				user.Dictionary[GetWordIndex(id, userTry.Dictionary)].WantToLearn = false
			}

			err := Collection.Update(Obj{"_id": userId}, user)
			if err != nil {
				c.String(http.StatusInternalServerError, "CheckThirdAlg: fail while trying to update words parameters: ", err)
			}
		} else {
			c.Status(http.StatusBadRequest)
		}
	}
	// c.Status(http.StatusOK)
}
