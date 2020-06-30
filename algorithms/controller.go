package algorithms

import (
	. "dictionaryProject/data"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	ai "github.com/night-codes/mgo-ai"
	"log"
	"net/http"
	"os"
	"time"
)

func createToken(authData User) string {
	jwtToken, err := generateToken(authData.Id)
	if err != nil {
		log.Println(err)
		return err.Error()
	}

	return jwtToken
}

func generateToken(id uint64) (token string, err error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, AccessTokenClaims{
		UserID:         id,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix()},
	})

	secretWord, _ := os.LookupEnv("SECRET_WORD")

	token, err = at.SignedString([]byte(secretWord))
	if err != nil {
		return "", err
	}

	return token, nil
}

func extractClaims(tokenStr string) (*AccessTokenClaims, bool) {
	token, _ := jwt.ParseWithClaims(tokenStr, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	if claims, ok := token.Claims.(*AccessTokenClaims); ok {
		return claims, true
	} else {
		return nil, false
	}
}

func RegistrationPost(c *gin.Context) {
	email := c.PostForm("email")
	pass := c.PostForm("pass")

	count, err := Collection.Find(Obj{"login": email}).Count()
	if err != nil {
		panic(err)
	}
	if count == 0 {
		user := User{
			Id:    ai.Next("user"),
			Email: email,
			Pass:  pass,
		}
		token := createToken(user)
		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "token",
			Value:    token,
			MaxAge:   86400,
			Path:     "/",
			Domain:   "localhost",
			SameSite: http.SameSiteStrictMode,
			Secure:   false,
			HttpOnly: true,
		})

		err = Collection.Insert(user)
		if err != nil {
			panic(err)
		}
		c.HTML(http.StatusOK, "main.html", nil)
	} else {
		c.HTML(http.StatusOK, "registration.html", nil) // Написать вывод ошибки !!!!!!!!!!!!!!!
	}
}

func AuthorizationPost(c *gin.Context) {
	email := c.PostForm("email")
	pass := c.PostForm("pass")

	count, err := Collection.Find(Obj{"login": email, "pass": pass}).Count()
	if err != nil {
		panic(err)
	}
	if count > 0 {
		user := User{}
		_ = Collection.Find(Obj{"login": email, "pass": pass}).One(&user)
		token := createToken(user)
		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "token",
			Value:    token,
			MaxAge:   86400,
			Path:     "/",
			Domain:   "localhost",
			SameSite: http.SameSiteStrictMode,
			Secure:   false,
			HttpOnly: true,
		})
		c.HTML(http.StatusOK, "main.html", Obj{
			"User": user,
		})

	} else {
		c.HTML(http.StatusOK, "authorization.html", nil) // Написать вывод ошибки !!!!!!!!!!!!!!!!
	}
	// c.String(http.StatusOK, "user")
}

func RegistrationGet(c *gin.Context) {
	c.HTML(http.StatusOK, "registration.html", nil)
}

func AuthorizationGet(c *gin.Context) {
	c.HTML(http.StatusOK, "authorization.html", nil)
}

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

func Default(c *gin.Context) {
	c.HTML(http.StatusOK, "authorization.html", nil)
}
