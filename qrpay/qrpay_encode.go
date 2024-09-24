package qrpay

import (
	"github.com/thanhtinhpas1/emvco_qr/pkg/constants"
	"github.com/thanhtinhpas1/emvco_qr/pkg/models"
)

type OptionFn func(qrPay *QRPay)

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

func BuildBasicQRPay(amount int64, bankBin string, accountNumber string, opts ...OptionFn) *QRPay {
	napasProvider := models.MerchantProvider{
		Id:         constants.NapasIdentifier,
		BankBin:    bankBin,
		TransferTo: accountNumber,
		Method:     models.NapasMethodAccountTransfer,
	}

	merchantInfo := models.MerchantInfo{
		NapasProvider: &napasProvider,
	}

	var initiationMethod = models.InitiationMethodStatic
	if amount > 0 {
		initiationMethod = models.InitiationMethodDynamic
	}

	qrPay := QRPay{
		Version:          constants.NapasDefaultVersion,
		InitiationMethod: initiationMethod,
		CurrencyCode:     constants.VietnameseCurrency,
		Amount:           amount,
		MerchantInfo:     &merchantInfo,
	}

	for _, optFn := range opts {
		optFn(&qrPay)
	}

	return &qrPay
}
