package models

import (
	"github.com/thanhtinhpas1/emvco_qr/pkg/constants"
)

type InitiationMethod string

var (
	InitiationMethodDynamic InitiationMethod = "11"
	InitiationMethodStatic  InitiationMethod = "12"
)

var InitiationMethodMap = map[string]InitiationMethod{
	"11": InitiationMethodDynamic,
	"12": InitiationMethodStatic,
}

type NapasMethod string

var (
	NapasMethodAccountTransfer NapasMethod = "QRIBFTTA"
	NapasMethodCardTransfer    NapasMethod = "QRIBFTTC"
)

type TipAndFeeType string

var (
	UserInputTip  TipAndFeeType = "01" // User will input the tip amount
	PredefinedTip TipAndFeeType = "02" // Predefined tip amount will be determined by in tag 56
)

type MerchantProvider struct {
	Id            string      `json:"id"`
	BankBin       string      `json:"bank_bin"`
	AccountNumber string      `json:"account_number"`
	CardNumber    string      `json:"card_number"`
	Method        NapasMethod `json:"method"`
}

type MerchantInfo struct {
	Name          string                         `json:"name" valid:"required,maxstringlength(25)"`
	City          string                         `json:"city"`
	CountryCode   string                         `json:"country_code"` // Reference to: https://developer.mastercard.com/card-issuance/documentation/code-and-formats/iso-country-and-currency-codes/
	PostalCode    string                         `json:"postal_code"`  // Zip code
	NapasProvider MerchantProvider               `json:"napas_provider"`
	MasterAccount string                         `json:"master_account"`
	VisaAccount   string                         `json:"visa_account"`
	JcbAccount    string                         `json:"jcb_account"`
	UpiAccount    string                         `json:"upi_account"`
	CategoryCode  constants.MerchantCategoryCode `json:"category_code"`
}

type QRPay struct {
	Version          string
	InitiationMethod InitiationMethod            `json:"initiation_method"`
	MerchantInfo     MerchantInfo                `json:"merchant_info"`
	CurrencyCode     string                      `json:"currency_code"` // https://vi.wikipedia.org/wiki/ISO_4217
	Amount           int64                       `json:"amount"`
	TipAndFeeType    TipAndFeeType               `json:"tip_and_fee_type"`
	TipAndFeeAmount  int64                       `json:"tip_and_fee_amount"`
	TipAndFeePercent int                         `json:"tip_and_fee_percent"`
	AdditionData     map[AdditionDataType]string `json:"addition_data"`
}
