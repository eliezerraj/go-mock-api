package model

import (
    "time"
)

type Balance struct {
	Id					string		`json:"balance_id"`
    Account 			string 		`json:"account"`
	Amount				int32 		`json:"amount"`
    DateBalance  		time.Time 	`json:"date_balance"`
	Description			string 		`json:"description"`
}