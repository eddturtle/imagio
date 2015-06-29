package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
)

const (
	SessionName = "imagio-tokens"
	TokenLength = 32
)

var (
	store   = sessions.NewCookieStore([]byte("something-very-secret"))
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

type Token struct {
	Value   string
	Created time.Time
}

func generateToken() string {
	rand.Seed(time.Now().Unix())
	b := make([]rune, TokenLength)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func getToken(w http.ResponseWriter, r *http.Request) (t Token) {
	session, _ := store.Get(r, SessionName)
	t = Token{}
	t.Value = generateToken()
	session.Values["csrf"] = t.Value
	session.Save(r, w)
	return t
}

func isValidToken(value string, r *http.Request) (result bool) {
	session, _ := store.Get(r, SessionName)
	if session.Values["csrf"] == value {
		return true
	}
	return false
}
