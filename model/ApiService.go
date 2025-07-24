package model

import "net/http"

type ApiService struct {
	Name    string
	BaseUrl string
	Port    string
	Path    string
	Header  map[string]string
	Body    any
	Env     string
	Cookie  *http.Cookie
	Jwt     string
}
