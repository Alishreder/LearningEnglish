package data

import (
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2"
	"sync"
)

var TokensLock sync.Mutex
var Tokens = map[string]User{}

const SecretWord = "letsgo"

type Obj map[string]interface{}

type WordTranslate struct {
	Word            string `json:"word" bson:"word"`
	Translate       string `json:"translate" bson:"translate"`
	WordId          uint64 `json:"word_id" bson:"word_id"`
	FirstAlgorithm  bool   `json:"first_algorithm" bson:"first_algorithm"`
	SecondAlgorithm bool   `json:"second_algorithm" bson:"second_algorithm"`
	ThirdAlgorithm  bool   `json:"third_algorithm" bson:"third_algorithm"`
	// FourthAlgorithm bool            `json:"fourth_algorithm" bson:"fourth_algorithm"`

}

type User struct {
	Id               uint64          `json:"id" bson:"_id"`
	Email            string          `json:"login" bson:"login"`
	Pass             string          `json:"pass" bson:"pass"`
	Dictionary       []WordTranslate `json:"dictionary" bson:"dictionary"`
	DictionarySize   uint64          `json:"dictionary_size" bson:"dictionary_size"`
	WordsForLearning []WordTranslate `json:"words_for_learning" bson:"words_for_learning"`
	LastWordId       uint64          `json:"last_word_id" bson:"last_word_id"`
}

type TokenStruct struct {
	Token string `json:"token" binding:"required"`
}

type AccessTokenClaims struct {
	UserID uint64
	jwt.StandardClaims
}

// For Database

var Session *mgo.Session
var DB *mgo.Database
var Collection *mgo.Collection
