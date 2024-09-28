package qrpay

import (
	"github.com/thanhtinhpas1/emvco_qr/pkg/constants"
	"github.com/thanhtinhpas1/emvco_qr/pkg/models"
)

type OptionFn func(qrPay *QRPay)

type MerchantOptionFn func(mc *models.MerchantInfo)

func WithVersion(version string) OptionFn {
	return func(qrPay *QRPay) {
		qrPay.Version = version
	}
}

func WithInitiationMethod(initiationMethod string) OptionFn {
	return func(qrPay *QRPay) {
		qrPay.InitiationMethod = models.InitiationMethod(initiationMethod)
	}
}

func WithMerchantInfo(merchantInfo *models.MerchantInfo) OptionFn {
	return func(qrPay *QRPay) {
		qrPay.MerchantInfo = merchantInfo
	}
}

func WithAmount(amount int64) OptionFn {
	return func(qrPay *QRPay) {
		qrPay.Amount = amount
	}
}

func WithCurrencyCode(currencyCode string) OptionFn {
	return func(qrPay *QRPay) {
		qrPay.CurrencyCode = currencyCode
	}
}

func WithTipAndFeeType(tipAndFeeType models.TipAndFeeType) OptionFn {
	return func(qrPay *QRPay) {
		qrPay.TipAndFeeType = tipAndFeeType
	}
}

func WithTipAndFeeAmount(tipAndFeeAmount int64) OptionFn {
	return func(qrPay *QRPay) {
		qrPay.TipAndFeeAmount = tipAndFeeAmount
	}
}

func WithTipAndFeePercent(tipAndFeePercent int) OptionFn {
	return func(qrPay *QRPay) {
		qrPay.TipAndFeePercent = tipAndFeePercent
	}
}

func WithDescription(description string) OptionFn {
	return func(qrPay *QRPay) {
		qrPay.Description = description
	}
}

func WithMerchantName(name string) MerchantOptionFn {
	return func(mc *models.MerchantInfo) {
		mc.Name = name
	}
}

func WithMerchantCity(city string) MerchantOptionFn {
	return func(mc *models.MerchantInfo) {
		mc.City = city
	}
}

func WithMerchantCountryCode(countryCode string) MerchantOptionFn {
	return func(mc *models.MerchantInfo) {
		mc.CountryCode = countryCode
	}
}

func WithMerchantPostalCode(postalCode string) MerchantOptionFn {
	return func(mc *models.MerchantInfo) {
		mc.PostalCode = postalCode
	}
}

func WithNapasProvider(napasProvider *models.NapasProvider) MerchantOptionFn {
	return func(mc *models.MerchantInfo) {
		mc.NapasProvider = napasProvider
	}
}

func WithMasterAccount(masterAccount string) MerchantOptionFn {
	return func(mc *models.MerchantInfo) {
		mc.MasterAccount = masterAccount
	}
}

func WithVisaAccount(visaAcc string) MerchantOptionFn {
	return func(mc *models.MerchantInfo) {
		mc.VisaAccount = visaAcc
	}
}

func WithJcbAccount(jcbAccount string) MerchantOptionFn {
	return func(mc *models.MerchantInfo) {
		mc.JcbAccount = jcbAccount
	}
}

func WithUpiAccount(upiAcc string) MerchantOptionFn {
	return func(mc *models.MerchantInfo) {
		mc.UpiAccount = upiAcc
	}
}

func WithCategoryCode(categoryCode constants.MerchantCategoryCode) MerchantOptionFn {
	return func(mc *models.MerchantInfo) {
		mc.CategoryCode = categoryCode
	}
}

func BuildMerchantInfo(merchantName string, opts ...MerchantOptionFn) *models.MerchantInfo {
	mcInfo := models.MerchantInfo{
		Name: merchantName,
	}

	for _, optFn := range opts {
		optFn(&mcInfo)
	}

	return &mcInfo
}

func BuildQRPay(amount int64, bankBin string, accountNumber string, opts ...OptionFn) *QRPay {
	napasProvider := models.NapasProvider{
		Id:         constants.NapasIdentifier,
		BankBin:    bankBin,
		TransferTo: accountNumber,
		Method:     models.NapasMethodAccountTransfer,
	}

	qrPay := QRPay{
		Version:      constants.NapasDefaultVersion,
		CurrencyCode: constants.Currency_VND,
		Amount:       amount,
	}

	for _, optFn := range opts {
		optFn(&qrPay)
	}

	if qrPay.MerchantInfo != nil {
		qrPay.MerchantInfo.NapasProvider = &napasProvider
	} else {
		qrPay.MerchantInfo = &models.MerchantInfo{
			NapasProvider: &napasProvider,
		}
	}

	var initiationMethod = models.InitiationMethodStatic
	if amount > 0 {
		initiationMethod = models.InitiationMethodDynamic
	}

	qrPay.InitiationMethod = initiationMethod

	return &qrPay
}
