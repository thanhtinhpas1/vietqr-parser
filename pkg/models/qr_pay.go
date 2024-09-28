package models

import (
	"errors"

	"github.com/thanhtinhpas1/vietqr-parser/pkg/constants"
)

var (
	QrPayWasNilError                          = errors.New("qr pay object must not nil")
	AmountMustPostiveError                    = errors.New("amount must greater than zero")
	TipAndFeeTypeError                        = errors.New("invalid tip and fee type")
	TipAndFeeAmountError                      = errors.New("tip and fee amount must greater than zero")
	TipAndFeePercentError                     = errors.New("tip and fee percent must between 0 - 100")
	TipAndFeeRequiredWhiteTypeWasDefinedError = errors.New("tip and fee required white type was defined")
	InitiationMethodError                     = errors.New("initiation method was invalid")
	CurrencyCodeWasEmptyError                 = errors.New("currency code must not be empty")
	NapasMethodError                          = errors.New("napas method was invalid")
	MerchantInfoWasNilError                   = errors.New("merchant info must be specified")
	NapasProviderWasNilError                  = errors.New("napas provider must be specified")
	NapasProviderIdEmptyError                 = errors.New("napas provider id must be specified")
	NapasProviderBankBinEmptyError            = errors.New("napas provider bank bin must be specified")
	NapasProviderInfoTransferToEmptyError     = errors.New("napas provider account/card number must be specified")
	NapasProviderCardNumberEmptyError         = errors.New("napas provider card number must be specified")
	NapasProviderMethodError                  = errors.New("napas provider method was invalid")
	ExceededMaxLengthError                    = errors.New("data has exceeded maximum length")
)

const (
	VersionMaxLength                 = 2
	InitiationMethodMaxLength        = 2
	MerchantAccountMaxLength         = 99
	MerchantCategoryCodeMaxLength    = 4
	TransactionCurrencyMaxLength     = 3
	AmountMaxLength                  = 13
	TipOrConvenienceMaxLength        = 2
	FeeFixedMaxLength                = 13
	FeePercentageMaxLength           = 5
	CountryCodeMaxLength             = 2
	MerchantNameMaxLength            = 25
	MerchantCityMaxLength            = 15
	PostalCodeMaxLength              = 10
	AdditionDataFieldMaxLength       = 99
	MerchantInfoTemplateMaxLength    = 99
	SubAdditionDataMaxLength         = 25
	SubAdditionConsumerDataMaxLength = 3
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

type NapasProvider struct {
	Id         string      `json:"id"`
	BankBin    string      `json:"bank_bin"`
	TransferTo string      `json:"transfer_to"`
	Method     NapasMethod `json:"method"`
}

type MerchantInfo struct {
	Name          string                         `json:"name" valid:"required,maxstringlength(25)"`
	City          string                         `json:"city"`
	CountryCode   string                         `json:"country_code"` // Reference to: https://developer.mastercard.com/card-issuance/documentation/code-and-formats/iso-country-and-currency-codes/
	PostalCode    string                         `json:"postal_code"`  // Zip code
	NapasProvider *NapasProvider                 `json:"napas_provider"`
	MasterAccount string                         `json:"master_account"`
	VisaAccount   string                         `json:"visa_account"`
	JcbAccount    string                         `json:"jcb_account"`
	UpiAccount    string                         `json:"upi_account"`
	CategoryCode  constants.MerchantCategoryCode `json:"category_code"`
}

func (mc *MerchantInfo) Validate() error {
	if mc.NapasProvider == nil {
		return NapasProviderWasNilError
	}

	if err := mc.NapasProvider.Validate(); err != nil {
		return err
	}

	return nil
}

func (np *NapasProvider) Validate() error {
	if len(np.Id) == 0 {
		return NapasProviderIdEmptyError
	}

	if len(np.BankBin) == 0 {
		return NapasProviderBankBinEmptyError
	}

	if len(np.Method) == 0 || (np.Method != NapasMethodAccountTransfer && np.Method != NapasMethodCardTransfer) {
		return NapasProviderMethodError
	}

	if len(np.TransferTo) == 0 {
		return NapasProviderInfoTransferToEmptyError
	}

	return nil
}
