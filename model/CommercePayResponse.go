package model

import "time"

type Response struct {
	Resutl Result `json:"result"`
}

type Result struct {
	TransactionNumber         string    `json:"transactionNumber"`
	ReferenceCode             string    `json:"referenceCode"`
	Status                    int       `json:"status"`
	CurrencyCode              string    `json:"currencyCode"`
	Amount                    int64     `json:"amount"`
	ChannelID                 int       `json:"channelId"`
	ProviderTransactionNumber string    `json:"providerTransactionNumber"`
	CreationTime              time.Time `json:"creationTime"`
	Remark                    string    `json:"remark"`
	ProviderChannelID         string    `json:"providerChannelId"`
	ProviderPaymentMethod     string    `json:"providerPaymentMethod"`
}
