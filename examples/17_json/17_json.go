package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type creditCard struct {
	Number         string `json:"number"`
	securityCode   string
	OwnerName      string    `json:"owner_name"`
	ExpirationDate time.Time `json:"expiration_date"`
}

func main() {
	// marshal
	cards := []creditCard{
		{
			Number:         "1234",
			securityCode:   "0000",
			OwnerName:      "Jason",
			ExpirationDate: time.Now().AddDate(1, 0, 0),
		},
		{
			Number:         "abcd",
			securityCode:   "0000",
			OwnerName:      "John",
			ExpirationDate: time.Now().AddDate(2, 0, 0),
		},
	}
	bytesCard, err := json.MarshalIndent(cards, "", "\t")
	if err != nil {
		fmt.Println("err:", err.Error())
	}
	fmt.Println("bytesCard:")
	fmt.Println(string(bytesCard))

	// unmarshal
	jsonCreditCards := `
[
        {
                "number": "5555",
                "owner_name": "Suzy",
                "expiration_date": "2024-11-09T11:52:51.862821122+01:00"
        },
        {
                "number": "6666",
                "owner_name": "Mary",
                "expiration_date": "2025-11-09T11:52:51.862931566+01:00"
        }
]`
	var creditCardsFromJson []creditCard
	if err := json.Unmarshal([]byte(jsonCreditCards), &creditCardsFromJson); err != nil {
		fmt.Println("err:", err.Error())
	}
	fmt.Println("jsonCreditCards")
	for _, v := range creditCardsFromJson {
		fmt.Println(v)
	}
}
