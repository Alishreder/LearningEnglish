package algorithms

import (
	. "LearningEnglish/data"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserFromBD(c *gin.Context) (User, error) {
	user := User{}
	token, err := c.Cookie("token")
	if err == nil {
		claims, err := extractClaims(token)
		if !err {
			fmt.Println("err:", err)
		} else {
			_ = Collection.Find(Obj{"_id": claims.UserID}).One(&user)
			return user, nil
		}
	}
	return user, err
}

func GetUserId(c *gin.Context) (uint64, error) {
	token, errf := c.Cookie("token")
	if errf == nil {
		claims, err := extractClaims(token)
		if !err {
			panic(err)
		} else {
			count, _ := Collection.Find(Obj{"_id": claims.UserID}).Count()
			if count > 0 {
				return claims.UserID, nil
			} else {
				errf = fmt.Errorf("can`t find user with id %v", claims.UserID)
			}
		}
	}
	return 0, errf
}

func Home(c *gin.Context) {
	user, err := GetUserFromBD(c)
	if err != nil {
		panic(err)
	}
	c.HTML(http.StatusOK, "main.html", Obj{
		"User": user,
	})
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
				"users":    users,
				"cur_user": user,
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

func Default(c *gin.Context) {
	c.HTML(http.StatusOK, "authorization.html", nil)
}
