package qrpay

import (
	"errors"
	"fmt"
	"strings"

	"github.com/biter777/countries"
	"github.com/spf13/cast"

	"github.com/thanhtinhpas1/emvco_qr/pkg/constants"
	"github.com/thanhtinhpas1/emvco_qr/pkg/models"
	"github.com/thanhtinhpas1/emvco_qr/pkg/utils"
)

type TagValue struct {
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

func NewTag(tagId constants.FieldID, value string) TagValue {
	return TagValue{
		Tag:   string(tagId),
		Value: value,
	}
}

func (v TagValue) String() string {
	if len(v.Value) > 0 {
		return fmt.Sprintf("%s%02d%s", v.Tag, len(v.Value), v.Value)
	}

	return ""
}

type QRPay struct {
	// Version of the QR
	Version string

	// Method of QR, fill Dynamic if has amount greater than zero, otherwise fill Static
	InitiationMethod models.InitiationMethod `json:"initiation_method"`

	// Merchant Info to fill in, such as bank bin acquirer, beneficiary number, etc
	MerchantInfo *models.MerchantInfo `json:"merchant_info"`

	// Currency code of QR, refenrece to https://en.wikipedia.org/wiki/ISO_4217
	CurrencyCode string `json:"currency_code"`

	// Amount of money which will be transferred
	Amount int64 `json:"amount"`

	// Tip and fee type of QR
	TipAndFeeType models.TipAndFeeType `json:"tip_and_fee_type"`

	// Tip and Fee type was defined, tip and fee amount will be take first if both amount and percentage are provided
	TipAndFeeAmount int64 `json:"tip_and_fee_amount"`

	// Percent of amount will be tip
	TipAndFeePercent int `json:"tip_and_fee_percent"`

	// Description when transferring
	Description string `json:"description"`

	// This field will be overwritten when build QR, so don't initialize it when building QR
	AdditionData map[models.AdditionDataType]string `json:"addition_data"`
}

func (qrPay *QRPay) Validate() error {
	if qrPay == nil {
		return models.QrPayWasNilError
	}

	// normalize before validation
	if qrPay.MerchantInfo != nil {
		qrPay.MerchantInfo.Name = utils.ConvertStringToAlphabet(qrPay.MerchantInfo.Name)
		qrPay.MerchantInfo.City = utils.ConvertStringToAlphabet(qrPay.MerchantInfo.City)
	}

	if len(qrPay.Description) > 0 {
		qrPay.Description = utils.ConvertStringToAlphabet(qrPay.Description)
	}

	if len(qrPay.AdditionData) > 0 {
		purposeOfTrans, exists := qrPay.AdditionData[models.PurposeOfTransactionType]
		if exists {
			qrPay.AdditionData[models.PurposeOfTransactionType] = utils.ConvertStringToAlphabet(purposeOfTrans)
		}
	} else {
		qrPay.AdditionData = make(map[models.AdditionDataType]string)
		qrPay.AdditionData[models.PurposeOfTransactionType] = qrPay.Description
	}

	if len(qrPay.Version) > models.VersionMaxLength {
		return errors.Join(models.ExceededMaxLengthError, fmt.Errorf("version max length %d", models.VersionMaxLength))
	}

	if len(qrPay.InitiationMethod) > models.InitiationMethodMaxLength {
		return errors.Join(models.ExceededMaxLengthError, fmt.Errorf("initiation method max length %d", models.InitiationMethodMaxLength))
	}

	if qrPay.MerchantInfo != nil && len(qrPay.MerchantInfo.CategoryCode) > models.MerchantCategoryCodeMaxLength {
		return errors.Join(models.ExceededMaxLengthError, fmt.Errorf("category code max length %d", models.MerchantCategoryCodeMaxLength))
	}

	if len(qrPay.CurrencyCode) > models.TransactionCurrencyMaxLength {
		return errors.Join(models.ExceededMaxLengthError, fmt.Errorf("currency code max length %d", models.TransactionCurrencyMaxLength))
	}

	if len(cast.ToString(qrPay.Amount)) > models.AmountMaxLength {
		return errors.Join(models.ExceededMaxLengthError, fmt.Errorf("amount max length %d", models.AmountMaxLength))
	}

	if len(cast.ToString(qrPay.TipAndFeeType)) > models.TipOrConvenienceMaxLength {
		return errors.Join(models.ExceededMaxLengthError, fmt.Errorf("tip or convenience type max length %d", models.TipOrConvenienceMaxLength))
	}

	if len(cast.ToString(qrPay.TipAndFeeAmount)) > models.FeeFixedMaxLength {
		return errors.Join(models.ExceededMaxLengthError, fmt.Errorf("tip and fee amount max length %d", models.FeeFixedMaxLength))
	}

	if len(cast.ToString(qrPay.TipAndFeePercent)) > models.FeePercentageMaxLength {
		return errors.Join(models.ExceededMaxLengthError, fmt.Errorf("tip and fee percentage max length %d", models.FeePercentageMaxLength))
	}

	if qrPay.MerchantInfo != nil && len(cast.ToString(qrPay.MerchantInfo.CountryCode)) > models.CountryCodeMaxLength {
		return errors.Join(models.ExceededMaxLengthError, fmt.Errorf("country code max length %d", models.CountryCodeMaxLength))
	}

	if qrPay.MerchantInfo != nil && len(cast.ToString(qrPay.MerchantInfo.Name)) > models.MerchantNameMaxLength {
		return errors.Join(models.ExceededMaxLengthError, fmt.Errorf("merchant name max length %d", models.MerchantNameMaxLength))
	}

	if qrPay.MerchantInfo != nil && len(cast.ToString(qrPay.MerchantInfo.City)) > models.MerchantCityMaxLength {
		return errors.Join(models.ExceededMaxLengthError, fmt.Errorf("merchant city max length %d", models.MerchantCityMaxLength))
	}

	if qrPay.MerchantInfo != nil && len(cast.ToString(qrPay.MerchantInfo.PostalCode)) > models.PostalCodeMaxLength {
		return errors.Join(models.ExceededMaxLengthError, fmt.Errorf("postal code max length %d", models.PostalCodeMaxLength))
	}

	if len(qrPay.AdditionData) > 0 {
		if len(qrPay.AdditionData[models.BillNumberType]) > models.SubAdditionDataMaxLength {
			return errors.Join(models.ExceededMaxLengthError, fmt.Errorf("bill number max length %d", models.SubAdditionDataMaxLength))
		}

		if len(qrPay.AdditionData[models.MobileNumberType]) > models.SubAdditionDataMaxLength {
			return errors.Join(models.ExceededMaxLengthError, fmt.Errorf("mobile number max length %d", models.SubAdditionDataMaxLength))
		}

		if len(qrPay.AdditionData[models.StoreLabelType]) > models.SubAdditionDataMaxLength {
			return errors.Join(models.ExceededMaxLengthError, fmt.Errorf("store label max length %d", models.SubAdditionDataMaxLength))
		}

		if len(qrPay.AdditionData[models.LoyaltyNumberType]) > models.SubAdditionDataMaxLength {
			return errors.Join(models.ExceededMaxLengthError, fmt.Errorf("loyalty number max length %d", models.SubAdditionDataMaxLength))
		}

		if len(qrPay.AdditionData[models.ReferenceLabelType]) > models.SubAdditionDataMaxLength {
			return errors.Join(models.ExceededMaxLengthError, fmt.Errorf("reference label max length %d", models.SubAdditionDataMaxLength))
		}

		if len(qrPay.AdditionData[models.CustomerLabelType]) > models.SubAdditionDataMaxLength {
			return errors.Join(models.ExceededMaxLengthError, fmt.Errorf("customer label max length %d", models.SubAdditionDataMaxLength))
		}

		if len(qrPay.AdditionData[models.TerminalLabelType]) > models.SubAdditionDataMaxLength {
			return errors.Join(models.ExceededMaxLengthError, fmt.Errorf("terminal max length %d", models.SubAdditionDataMaxLength))
		}

		if len(qrPay.AdditionData[models.PurposeOfTransactionType]) > models.SubAdditionDataMaxLength {
			return errors.Join(models.ExceededMaxLengthError, fmt.Errorf("purpose of transaction max length %d", models.SubAdditionDataMaxLength))
		}

		if len(qrPay.AdditionData[models.ConsumerDataType]) > models.SubAdditionConsumerDataMaxLength {
			return errors.Join(models.ExceededMaxLengthError, fmt.Errorf("consumer data request max length %d", models.SubAdditionConsumerDataMaxLength))
		}
	}

	if qrPay.Amount < 0 {
		return models.AmountMustPostiveError
	}

	if len(qrPay.TipAndFeeType) > 0 && qrPay.TipAndFeeType != models.UserInputTip && qrPay.TipAndFeeType != models.PredefinedTip {
		return models.TipAndFeeTypeError
	}

	if qrPay.TipAndFeeAmount < 0 {
		return models.TipAndFeeAmountError
	}

	if qrPay.TipAndFeePercent < 0 || qrPay.TipAndFeePercent > 100 {
		return models.TipAndFeePercentError
	}

	if len(qrPay.TipAndFeeType) > 0 && qrPay.TipAndFeeAmount == 0 && qrPay.TipAndFeePercent == 0 {
		return models.TipAndFeeRequiredWhiteTypeWasDefinedError
	}

	if qrPay.InitiationMethod != models.InitiationMethodDynamic && qrPay.InitiationMethod != models.InitiationMethodStatic {
		return models.InitiationMethodError
	}

	if qrPay.MerchantInfo == nil {
		return models.MerchantInfoWasNilError
	}

	if len(qrPay.CurrencyCode) == 0 {
		return models.CurrencyCodeWasEmptyError
	}

	if err := qrPay.MerchantInfo.Validate(); err != nil {
		return err
	}

	return nil
}

func (qp *QRPay) GenerateQRCode() (string, error) {
	if err := qp.Validate(); err != nil {
		return "", err
	}

	// TAG 00
	versionTag := NewTag(constants.FIELD_ID_Version, qp.Version)

	// TAG 01
	initiationMethodTag := NewTag(constants.FIELD_ID_Method, string(qp.InitiationMethod))

	// => START TAG 38 VIETQR
	napasIdentifyTag := NewTag(constants.FIELD_ID_Subtag_Id, constants.NapasIdentifier)
	napasMethodTag := NewTag(constants.FIELD_ID_Subtag_Service, string(qp.MerchantInfo.NapasProvider.Method))

	bankBinTag := NewTag(constants.FIELD_ID_Subtag_Id, qp.MerchantInfo.NapasProvider.BankBin)
	transferToTag := NewTag(constants.FIELD_ID_Subtag_Data, qp.MerchantInfo.NapasProvider.TransferTo)
	accountInfoTag := NewTag(constants.FIELD_ID_Subtag_Data, fmt.Sprintf("%s%s", bankBinTag, transferToTag))

	vietQRTag := NewTag(constants.FIELD_ID_VietQR, fmt.Sprintf("%s%s%s", napasIdentifyTag, accountInfoTag, napasMethodTag))

	if len(vietQRTag.String()) > models.MerchantAccountMaxLength+2 {
		return "", errors.Join(models.ExceededMaxLengthError, fmt.Errorf("vietqr tag max length %d", models.MerchantAccountMaxLength))
	}
	// => END TAG 38 VIETQR

	merchantCategoryTag := NewTag(constants.FIELD_ID_Category, string(qp.MerchantInfo.CategoryCode))

	// TAG 53
	currencyTag := NewTag(constants.FIELD_ID_Currency, qp.CurrencyCode)

	// TAG 54
	amountTag := NewTag(constants.FIELD_ID_Amount, cast.ToString(qp.Amount))

	// => START TAG 55, 56, 57: TIP OR CONVENIENCE INDICATOR
	tipAndFeeTypeTag := NewTag(constants.FIELD_ID_Tip_And_Fee_Type, string(qp.TipAndFeeType))

	var (
		tipAndFeeAmountTag, tipAndFeePercentTag TagValue
	)
	if qp.TipAndFeeType == models.PredefinedTip {
		if qp.TipAndFeeAmount > 0 {
			tipAndFeeAmountTag = NewTag(constants.FIELD_ID_Tip_And_Fee_Amount, cast.ToString(qp.TipAndFeeAmount))
		} else if qp.TipAndFeePercent > 0 {
			tipAndFeePercentTag = NewTag(constants.FIELD_ID_Tip_And_Fee_Percent, cast.ToString(qp.TipAndFeePercent))
		}
	}
	// => END TAG 55, 56, 57

	// TAG 58
	var countryTag TagValue
	if len(qp.MerchantInfo.CountryCode) > 0 {
		countryTag = NewTag(constants.FIELD_ID_Country_Code, qp.MerchantInfo.CountryCode)
	} else {
		countryTag = NewTag(constants.FIELD_ID_Country_Code, cast.ToString(countries.Vietnam.Info().Alpha2))
	}

	// TAG 59
	merchantName := NewTag(constants.FIELD_ID_Merchant_Name, qp.MerchantInfo.Name)

	// TAG 60
	cityTag := NewTag(constants.FIELD_ID_Merchant_City, qp.MerchantInfo.City)

	// TAG 61
	postalTag := NewTag(constants.FIELD_ID_Postal_Code, qp.MerchantInfo.PostalCode)

	// => START TAG 62: Addition SUB TAG 08 => DESCRIPTION <=> Purpose of transaction
	billNumberTag := NewTag(constants.FieldID(string(models.BillNumberType)), qp.AdditionData[models.BillNumberType])
	mobileNumberTag := NewTag(constants.FieldID(string(models.MobileNumberType)), qp.AdditionData[models.MobileNumberType])
	storeLabelTag := NewTag(constants.FieldID(string(models.StoreLabelType)), qp.AdditionData[models.StoreLabelType])
	loyaltyTag := NewTag(constants.FieldID(string(models.LoyaltyNumberType)), qp.AdditionData[models.LoyaltyNumberType])
	referenceTag := NewTag(constants.FieldID(string(models.ReferenceLabelType)), qp.AdditionData[models.ReferenceLabelType])
	customerLabelTag := NewTag(constants.FieldID(string(models.CustomerLabelType)), qp.AdditionData[models.CustomerLabelType])
	terminalLabelTag := NewTag(constants.FieldID(string(models.TerminalLabelType)), qp.AdditionData[models.TerminalLabelType])
	purposeOfTransaction := NewTag(constants.FieldID(string(models.PurposeOfTransactionType)), qp.AdditionData[models.PurposeOfTransactionType])
	consumerDataRequestTag := NewTag(constants.FieldID(string(models.ConsumerDataType)), qp.AdditionData[models.ConsumerDataType])
	strBuilder := strings.Builder{}
	strBuilder.WriteString(fmt.Sprintf("%s%s%s%s%s%s%s%s%s", billNumberTag, mobileNumberTag, storeLabelTag, loyaltyTag, referenceTag, customerLabelTag, terminalLabelTag, purposeOfTransaction, consumerDataRequestTag))

	// build for rfu, unreversed fields
	for key, val := range qp.AdditionData {
		if cast.ToInt(key) <= 10 {
			continue
		}

		newTag := NewTag(constants.FieldID(key), val)
		strBuilder.WriteString(newTag.String())
	}

	additionTag := NewTag(constants.FIELD_ID_Additional_Data, strBuilder.String())

	// => END TAG 62

	// TAG 64: MERCHANT INFORMATION - Language Template
	// TDB

	// TAG 65 -> 79: RFU for EMVCo (register by EMVCo)
	// TDB

	// TAG 80 -> 99:  Unreserved Template (register for future usage)
	// TDB

	qrContent := strings.Join([]string{
		versionTag.String(),
		initiationMethodTag.String(),
		vietQRTag.String(),
		merchantCategoryTag.String(),
		currencyTag.String(),
		amountTag.String(),
		tipAndFeeTypeTag.String(),
		tipAndFeeAmountTag.String(),
		tipAndFeePercentTag.String(),
		countryTag.String(),
		merchantName.String(),
		cityTag.String(),
		postalTag.String(),
		additionTag.String(),
		string(constants.FIELD_ID_Crc),
		"04",
	}, "")

	// TAG 64: CRC (Cyclic Redundancy Check)
	tag63Str := utils.GetCRC(qrContent)

	return fmt.Sprintf("%s%s", qrContent, tag63Str), nil
}
