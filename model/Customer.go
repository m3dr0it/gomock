package model

import "net/http"

type Customer struct {
	Username string       `json:"user"`
	Password string       `json:"password"`
	Cookie   *http.Cookie `json:"-"`
}
