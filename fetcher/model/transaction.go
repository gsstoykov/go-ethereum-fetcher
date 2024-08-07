package model

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	TransactionHash   string `json:"transactionHash" gorm:"column:transaction_hash;unique;not null"`
	TransactionStatus int    `json:"transactionStatus" gorm:"column:transaction_status"`
	BlockHash         string `json:"blockHash" gorm:"column:block_hash"`
	BlockNumber       int    `json:"blockNumber" gorm:"column:block_number"`
	From              string `json:"from" gorm:"column:from_address"`
	To                string `json:"to" gorm:"column:to_address"`
	ContractAddress   string `json:"contractAddress" gorm:"column:contract_address"`
	LogsCount         int    `json:"logsCount" gorm:"column:logs_count"`
	Input             string `json:"input" gorm:"column:input"`
	Value             string `json:"value" gorm:"column:value"`
}
