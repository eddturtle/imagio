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
	Value string
}

func generateToken() string {
	// Seed other-wise you get the same token on 1st launch
	rand.Seed(time.Now().Unix())
	b := make([]rune, TokenLength)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func getToken(w http.ResponseWriter, r *http.Request) (t Token) {
	session, _ := store.Get(r, SessionName)
	t = Token{Value: generateToken()}
	session.Values["csrf"] = t.Value
	session.Save(r, w)
	return t
}

func isValidToken(value string, r *http.Request) bool {
	session, _ := store.Get(r, SessionName)
	return session.Values["csrf"] == value
}
