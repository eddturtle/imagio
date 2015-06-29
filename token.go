package main

import (
	"net/http"
	"time"
    "math/rand"

	"github.com/gorilla/sessions"
)

const (
	SessionName = "imagio-tokens"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

type Token struct {
	Value   string
	Created time.Time
}

func GenerateToken() string {
	rand.Seed(time.Now().Unix())
    b := make([]rune, 32)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func GetToken(w http.ResponseWriter, r *http.Request) (t Token) {
	session, _ := store.Get(r, SessionName)
	t = Token{Value: ""}
	t.Value = GenerateToken()
	session.Values["csrf"] = t.Value
	session.Save(r, w)
	return t
}

func IsValidToken(value string, r* http.Request) (result bool) {
    session, _ := store.Get(r, SessionName)
    if session.Values["csrf"] == value {
    	return true
    }
    return false
}
