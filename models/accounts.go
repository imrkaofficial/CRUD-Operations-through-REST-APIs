package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	UID                string  `json:"uid" gorm:"type:varchar(255);uniqueIndex"`
	First              string  `json:"first" gorm:"type:varchar(255)"`
	Last               string  `json:"last" gorm:"type:varchar(255)"`
	Email              string  `json:"email" gorm:"type:varchar(255);uniqueIndex"`
	Phone1             string  `json:"phone1" gorm:"type:varchar(255)"`
	Phone2             string  `json:"phone2" gorm:"type:varchar(255)"`
	Address            string  `json:"address" gorm:"type:varchar(255)"`
	Balance            float64 `json:"balance_money" gorm:"type:decimal(12,6)"`
	Type               string  `json:"type" gorm:"type:varchar(255)"`
	Country            string  `json:"country" gorm:"type:varchar(255)"`
	GSTNumber          string  `json:"gst_number" gorm:"type:varchar(255)"`
	City               string  `json:"city" gorm:"type:varchar(255)"`
	PinCode            string  `json:"pin_code" gorm:"type:varchar(255)"`
	RegisteredCompName string  `json:"registered_company_name" gorm:"type:varchar(255)"`
	DataStoreConsent   bool    `json:"data_store_consent"`
	IsDisabled         bool    `json:"is_disabled"`
}
