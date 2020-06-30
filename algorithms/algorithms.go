package algorithms

import (
	. "dictionaryProject/data"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
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
			// FourthAlgorithm: true,
		})
		user.WordsForLearning = append(user.WordsForLearning, WordTranslate{
			Word:            word,
			Translate:       translate,
			WordId:          id,
			FirstAlgorithm:  true,
			SecondAlgorithm: true,
			ThirdAlgorithm:  true,
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
		if user.Dictionary[index] != user.WordsForLearning[GetWordIndex(idToLearn, user.WordsForLearning)] {
			user.WordsForLearning = append(user.WordsForLearning, word)
		}
	} else {
		index := GetWordIndex(idToLearn, user.Dictionary)
		word := user.Dictionary[index]
		word.FirstAlgorithm = true
		word.SecondAlgorithm = true
		word.ThirdAlgorithm = true
		user.WordsForLearning = append(user.WordsForLearning, word)
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
		"list": user.WordsForLearning,
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
			c.Status(http.StatusOK)
			user.Dictionary[GetWordIndex(id, userTry.Dictionary)].ThirdAlgorithm = false
			index := GetWordIndex(id, userTry.WordsForLearning)
			user.WordsForLearning[index].ThirdAlgorithm = false
			word := user.WordsForLearning[GetWordIndex(id, userTry.WordsForLearning)]
			if !word.FirstAlgorithm && !word.SecondAlgorithm && !word.ThirdAlgorithm {
				user.WordsForLearning = append(user.WordsForLearning[:index], user.WordsForLearning[index+1:]...)
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

func IsAdmin(user User) bool {

	adminPass, _ := os.LookupEnv("ADMIN_PASSWORD")
	adminEmail, _ := os.LookupEnv("ADMIN_EMAIL")
	if user.Pass == adminPass &&
		user.Email == adminEmail {
		return true
	}

	return false
}

func GetAllUsersFromBD() (users []User, err error) {
	err = Collection.Find(Obj{}).All(&users)
	users = append(users[1:])
	return users, err
}

func ShowUsersList(c *gin.Context) {
	user, err := GetUserFromBD(c)
	if err != nil {
		panic(err)
	}
	if IsAdmin(user) {
		users, err := GetAllUsersFromBD()
		if err != nil {
			c.Status(http.StatusNotFound)
		} else {
			c.HTML(http.StatusOK, "usersList.html", Obj{
				"users": users,
			})
		}
	}
}

func ShowUsersDictionary(c *gin.Context) {
	idInt, _ := strconv.Atoi(c.Query("id"))
	id := uint64(idInt)
	user := User{}
	count, err := Collection.Find(Obj{"_id": id}).Count()
	if err != nil {
		panic(err)
	}
	if count > 0 {
		_ = Collection.Find(Obj{"_id": id}).One(&user)
		c.HTML(http.StatusOK, "usersDictionary.html", Obj{
			"user": user,
		})
	} else {
		c.Status(http.StatusBadRequest)
	}

}
