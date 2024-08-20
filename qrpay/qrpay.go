package qrpay

import (
	"errors"
	"strconv"

	"github.com/spf13/cast"

	"github.com/thanhtinhpas1/emvco_qr/pkg/constants"
	"github.com/thanhtinhpas1/emvco_qr/pkg/models"
	validator "github.com/thanhtinhpas1/emvco_qr/pkg/validators"
)

func ParseQRPay(qrCode string) (*models.QRPay, error) {
	if err := validator.ValidateQR(qrCode); err != nil {
		return nil, err
	}

	return parseQRPay(qrCode)
}

func parseQRPay(qrCode string) (*models.QRPay, error) {
	var qrPay = models.QRPay{}
	rootEmvcoMap := parseEmvcoTag(qrCode)

	// 00: version
	versionVal := rootEmvcoMap[constants.FIELD_ID_Version]
	if len(versionVal) == 0 {
		return nil, errors.New("version field is mandatory")
	}
	qrPay.Version = versionVal

	// 01: initiation method
	method := models.InitiationMethodMap[rootEmvcoMap[constants.FIELD_ID_Method]]
	if len(method) == 0 {
		return nil, errors.New("initiation method is mandatory")
	}
	qrPay.InitiationMethod = models.InitiationMethod(method)

	qrPay.MerchantInfo = parseMerchantInfo(rootEmvcoMap)

	// 53: currency
	qrPay.CurrencyCode = rootEmvcoMap[constants.FIELD_ID_Currency]

	// 54: amount
	qrPay.Amount = cast.ToInt64(rootEmvcoMap[constants.FIELD_ID_Amount])

	qrPay.TipAndFeeType = models.TipAndFeeType(rootEmvcoMap[constants.FIELD_ID_Tip_And_Fee_Type])
	qrPay.TipAndFeeAmount = cast.ToInt64(rootEmvcoMap[constants.FIELD_ID_Tip_And_Fee_Amount])
	qrPay.TipAndFeePercent = cast.ToInt(rootEmvcoMap[constants.FIELD_ID_Tip_And_Fee_Percent])

	qrPay.AdditionData = parseAdditionData(rootEmvcoMap)

	return &qrPay, nil
}

func parseEmvcoTag(content string) map[constants.FieldID]string {
	var (
		emvcoMap = make(map[constants.FieldID]string)
		i        = int(0)
	)
	for i < len(content) {
		tagIndex := i + constants.EmvcoTagLength

		tagName := constants.FieldID(content[i:tagIndex])
		tagLengthVal, _ := strconv.Atoi(content[tagIndex : tagIndex+constants.EmvcoTagLength])
		tagValue := content[tagIndex+constants.EmvcoTagLength : tagIndex+constants.EmvcoTagLength+tagLengthVal]

		emvcoMap[tagName] = tagValue

		i = i + (2*constants.EmvcoTagLength + tagLengthVal) // index + length of a tag (tag name | tag length | tag value)
	}

	return emvcoMap
}

func parseMerchantInfo(emvcoMap map[constants.FieldID]string) models.MerchantInfo {
	var (
		merchantInfo models.MerchantInfo
	)

	// 02: Visa
	merchantInfo.VisaAccount = emvcoMap[constants.Field_ID_Visa]

	// 04: master card
	merchantInfo.MasterAccount = emvcoMap[constants.Field_ID_Master]

	// 13: jcb
	merchantInfo.JcbAccount = emvcoMap[constants.Field_ID_JCB]

	// 15: upi
	merchantInfo.UpiAccount = emvcoMap[constants.Field_ID_UPI]

	// 38: napas provider
	napasEmvcoTag := emvcoMap[constants.FIELD_ID_VietQR]
	napasProvider := parseNapasProvider(napasEmvcoTag)
	merchantInfo.NapasProvider = napasProvider

	// 52: category code
	merchantInfo.CategoryCode = constants.MerchantCategoryCode(emvcoMap[constants.FIELD_ID_Category])

	// 58: country code
	merchantInfo.CountryCode = emvcoMap[constants.FIELD_ID_Country_Code]

	// 59: merchant name
	merchantInfo.Name = emvcoMap[constants.FIELD_ID_Merchant_Name]

	// 60: merchant city
	merchantInfo.City = emvcoMap[constants.FIELD_ID_Merchant_City]

	// 61: postal code
	merchantInfo.PostalCode = emvcoMap[constants.FIELD_ID_Postal_Code]

	return merchantInfo
}

func parseNapasProvider(napasTag string) models.MerchantProvider {
	var napasProvider models.MerchantProvider
	napasEmvcoMap := parseEmvcoTag(napasTag)

	// 00
	napasProvider.Id = napasEmvcoMap[constants.FIELD_ID_Subtag_Id]

	napasData := napasEmvcoMap[constants.FIELD_ID_Subtag_Data]
	// 01
	napasDataEmvcoMap := parseEmvcoTag(napasData)
	napasProvider.BankBin = napasDataEmvcoMap[constants.FIELD_ID_Subtag_Id]
	napasProvider.AccountNumber = napasDataEmvcoMap[constants.FIELD_ID_Subtag_Data]

	// 02
	napasProvider.Method = models.NapasMethod(napasEmvcoMap[constants.FIELD_ID_Subtag_Service])

	return napasProvider
}

func parseAdditionData(emvcoMap map[constants.FieldID]string) map[models.AdditionDataType]string {
	var addMap = make(map[models.AdditionDataType]string)

	additionEmvcoMap := parseEmvcoTag(emvcoMap[constants.FIELD_ID_Additional_Data])
	billNumber := additionEmvcoMap[constants.FieldID(models.BillNumberType)]
	if len(billNumber) > 0 {
		addMap[models.BillNumberType] = billNumber
	}

	mobileNumber := additionEmvcoMap[constants.FieldID(models.MobileNumberType)]
	if len(mobileNumber) > 0 {
		addMap[models.MobileNumberType] = mobileNumber
	}

	storeLabel := additionEmvcoMap[constants.FieldID(models.StoreLabelType)]
	if len(mobileNumber) > 0 {
		addMap[models.StoreLabelType] = storeLabel
	}

	loyaltyNumber := additionEmvcoMap[constants.FieldID(models.LoyaltyNumberType)]
	if len(mobileNumber) > 0 {
		addMap[models.LoyaltyNumberType] = loyaltyNumber
	}

	referenceLabel := additionEmvcoMap[constants.FieldID(models.ReferenceLabelType)]
	if len(mobileNumber) > 0 {
		addMap[models.ReferenceLabelType] = referenceLabel
	}

	customerLabel := additionEmvcoMap[constants.FieldID(models.CustomerLabelType)]
	if len(mobileNumber) > 0 {
		addMap[models.CustomerLabelType] = customerLabel
	}

	terminalLabel := additionEmvcoMap[constants.FieldID(models.TerminalLabelType)]
	if len(mobileNumber) > 0 {
		addMap[models.TerminalLabelType] = terminalLabel
	}

	purposeOfTrans := additionEmvcoMap[constants.FieldID(models.PurposeOfTransactionType)]
	if len(mobileNumber) > 0 {
		addMap[models.PurposeOfTransactionType] = purposeOfTrans
	}

	consumerDataType := additionEmvcoMap[constants.FieldID(models.ConsumerDataType)]
	if len(mobileNumber) > 0 {
		addMap[models.ConsumerDataType] = consumerDataType
	}

	return addMap
}
