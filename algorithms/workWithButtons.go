package algorithms

import (
	. "LearningEnglish/data"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddNewWord(c *gin.Context) {
	user, err := GetUserFromBD(c)
	if err != nil {
		panic(err)
	}
	userId, err := GetUserId(c)
	if err != nil {
		panic(err)
	}

	word := c.PostForm("word")
	translate := c.PostForm("translate")
	if word != "" || translate != "" {
		id := user.LastWordId + 1
		user.LastWordId++

		user.Dictionary = append(user.Dictionary, WordTranslate{
			Word:            word,
			Translate:       translate,
			WordId:          id,
			FirstAlgorithm:  true,
			SecondAlgorithm: true,
			ThirdAlgorithm:  true,
			WantToLearn:     true,
			// FourthAlgorithm: true,
		})
		user.WordsForLearning = append(user.WordsForLearning, WordTranslate{
			Word:            word,
			Translate:       translate,
			WordId:          id,
			FirstAlgorithm:  true,
			SecondAlgorithm: true,
			ThirdAlgorithm:  true,
			WantToLearn:     true,
			// FourthAlgorithm: true,
		})
		err = Collection.Update(Obj{"_id": userId}, user)
		if err != nil {
			c.String(http.StatusInternalServerError, "AddNewWord: fail while trying to add new word to dictionary: ", err)
		}
	}

	c.Status(http.StatusOK)
}

func DeleteWord(c *gin.Context) {
	user, err := GetUserFromBD(c)
	if err != nil {
		panic(err)
	}
	userId, err := GetUserId(c)
	if err != nil {
		panic(err)
	}

	idInt, _ := strconv.Atoi(c.PostForm("id"))
	idToDelete := uint64(idInt)
	for i, v := range user.Dictionary {
		if v.WordId == idToDelete {
			user.Dictionary = append(user.Dictionary[:i], user.Dictionary[i+1:]...)
			for j, val := range user.WordsForLearning {
				if v.WordId == val.WordId {
					user.WordsForLearning = append(user.WordsForLearning[:j], user.WordsForLearning[j+1:]...)
				}
			}
			err = Collection.Update(Obj{"_id": userId}, user)
			if err != nil {
				c.String(http.StatusInternalServerError, "DeleteWord: fail while trying to delete word: ", err)
			}
		}
	}

	c.Status(http.StatusOK)
}

func AddToLearnList(c *gin.Context) {
	user, err := GetUserFromBD(c)
	if err != nil {
		panic(err)
	}
	userId, err := GetUserId(c)
	if err != nil {
		panic(err)
	}

	idInt, _ := strconv.Atoi(c.Param("id"))
	idToLearn := uint64(idInt)

	if len(user.WordsForLearning) > 0 {
		index := GetWordIndex(idToLearn, user.Dictionary)
		word := user.Dictionary[index]
		word.FirstAlgorithm = true
		word.SecondAlgorithm = true
		word.ThirdAlgorithm = true
		user.Dictionary[index].WantToLearn = true
		user.WordsForLearning = append(user.WordsForLearning, word)
	} else {
		index := GetWordIndex(idToLearn, user.Dictionary)
		word := user.Dictionary[index]
		word.FirstAlgorithm = true
		word.SecondAlgorithm = true
		word.ThirdAlgorithm = true
		word.WantToLearn = true
		user.WordsForLearning = append(user.WordsForLearning, word)
		user.Dictionary[index].WantToLearn = true
	}
	err = Collection.Update(Obj{"_id": userId}, user)
	if err != nil {
		c.String(http.StatusInternalServerError, "AddToLearnList: fail while trying to add word to learn list: ", err)
	}
	c.Status(http.StatusOK)
}

func ShowLearnList(c *gin.Context) {
	user, err := GetUserFromBD(c)
	if err != nil {
		panic(err)
	}

	c.HTML(http.StatusOK, "LearnList.html", Obj{
		"user": user,
	})
}

func Learn(c *gin.Context) {
	user, err := GetUserFromBD(c)
	if err != nil {
		panic(err)
	}

	c.HTML(http.StatusOK, "algorithms.html", Obj{
		"user": user,
	})
}
