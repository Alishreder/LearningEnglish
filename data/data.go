package data

import (
	"gopkg.in/mgo.v2"
	"sync"
)

var UsersLock sync.Mutex

// var TokensLock sync.Mutex

var Tokens = map[string]User{}
// var Users = map[User]struct{}{}

const SecretWord = "letsgo"

type Obj map[string]interface{}

type WordTranslate struct {
	Word      string `json:"word" bson:"word"`
	Translate string `json:"" bson:""`
	WantToLearn bool `json:"" bson:""`
	WordId uint64 `json:"" bson:""`
}

type User struct {
	Id         uint64          `json:"id" bson:"_id"`
	Email      string          `json:"login" bson:"login"`
	Pass       string          `json:"pass" bson:"pass"`
	Dictionary []WordTranslate `json:"dictionary" bson:"dictionary"`
	LastWordId uint64 `json:"" bson:""`
}

type TokenStruct struct {
	Token string `json:"token" binding:"required"`
}

var CurrentUser User
var CurrId uint64

// For Database

var Session *mgo.Session
var DB *mgo.Database
var Collection *mgo.Collection
