package utils

import (
	"github.com/go-resty/resty/v2"
	"gomock/model"
)

func doPost(apiService model.ApiService) {
	client := resty.New()

	client.SetBaseURL(apiService.BaseUrl)
	for key, value := range apiService.Header {
		client.SetHeader(key, value)
	}

	if apiService.Cookie != nil {
		client.SetCookie(apiService.Cookie)
	}

}
