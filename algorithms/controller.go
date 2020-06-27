package algorithms

import (
	. "dictionaryProject/data"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	ai "github.com/night-codes/mgo-ai"
	"log"
	"net/http"
	"time"
)

func createToken(authData User) string {
	jwtToken, err := generateToken()
	if err != nil {
		log.Println(err)
		return err.Error()
	}

	Tokens[jwtToken] = authData

	return jwtToken
}

func generateToken() (token string, err error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err = at.SignedString([]byte(SecretWord))
	if err != nil {
		return "", err
	}

	return token, nil
}

func RegistrationPost(c *gin.Context) {
	email := c.PostForm("email")
	pass := c.PostForm("pass")

	count, err := Collection.Find(Obj{"login": email}).Count()
	if err != nil {
		panic(err)
	}
	if count == 0 {
		CurrentUser = User{
			Id:    ai.Next("user"),
			Email: email,
			Pass:  pass,
		}
		CurrId = CurrentUser.Id

		err = Collection.Insert(CurrentUser)
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
		_ = Collection.Find(Obj{"login": email, "pass": pass}).One(&CurrentUser)
		CurrId = CurrentUser.Id
		c.HTML(http.StatusOK, "main.html", Obj{
			"User": CurrentUser,
		})
	} else {
		c.HTML(http.StatusOK, "authorization.html", nil)  // Написать вывод ошибки !!!!!!!!!!!!!!!!
	}
	user := User{
		Email: email,
		Pass:  pass,
	}
	token := createToken(user)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    token,
		MaxAge:   60,
		Path:     "/",
		Domain:   "localhost",
		SameSite: http.SameSiteStrictMode,
		Secure:   false,
		HttpOnly: true,
	})

	// c.String(http.StatusOK, "user")
}

func RegistrationGet(c *gin.Context) {
	c.HTML(http.StatusOK, "registration.html", nil)
}

func AuthorizationGet(c *gin.Context) {
	c.HTML(http.StatusOK, "authorization.html", nil)
}

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "main.html", Obj{
		"User": CurrentUser,
	})
}

func Default(c *gin.Context) {
	c.HTML(http.StatusOK, "authorization.html", nil)
}
