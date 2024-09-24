package constants

const EmvcoTagLength = 2

type FieldID string

const (
	FIELD_ID_Version             FieldID = "00"
	FIELD_ID_Method              FieldID = "01"
	Field_ID_Visa                FieldID = "02"
	Field_ID_Master              FieldID = "04"
	Field_ID_JCB                 FieldID = "13"
	Field_ID_UPI                 FieldID = "15"
	FIELD_ID_VietQR              FieldID = "38"
	FIELD_ID_Category            FieldID = "52"
	FIELD_ID_Currency            FieldID = "53"
	FIELD_ID_Amount              FieldID = "54"
	FIELD_ID_Tip_And_Fee_Type    FieldID = "55"
	FIELD_ID_Tip_And_Fee_Amount  FieldID = "56"
	FIELD_ID_Tip_And_Fee_Percent FieldID = "57"
	FIELD_ID_Country_Code        FieldID = "58"
	FIELD_ID_Merchant_Name       FieldID = "59"
	FIELD_ID_Merchant_City       FieldID = "60"
	FIELD_ID_Postal_Code         FieldID = "61"
	FIELD_ID_Additional_Data     FieldID = "62"
	FIELD_ID_Crc                 FieldID = "63"
)

const (
	FIELD_ID_Subtag_Id      FieldID = "00"
	FIELD_ID_Subtag_Data    FieldID = "01"
	FIELD_ID_Subtag_Service FieldID = "02"
)

const (
	FIELD_ID_Subtag_Addition_Description FieldID = "08"
)
