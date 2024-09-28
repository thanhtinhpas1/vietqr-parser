package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/biter777/countries"

	"github.com/thanhtinhpas1/vietqr_parser/pkg/constants"
	"github.com/thanhtinhpas1/vietqr_parser/pkg/models"
	"github.com/thanhtinhpas1/vietqr_parser/qrpay"
)

func main() {
	/** BASIC QR */
	basicQRPay := qrpay.BuildQRPay(constants.Bin_VCB, "0881000458086")
	qr, err := basicQRPay.GenerateQRCode()
	if err != nil {
		fmt.Println("generate qr code failed", err)
		os.Exit(1)
	}

	fmt.Printf("QR Code: %s\n", qr)

	qrPay, err := qrpay.ParseQRPay(qr)
	if err != nil {
		fmt.Printf("parse VietQR ERROR: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Amount: %v\n", qrPay.Amount)
	data, _ := json.Marshal(qrPay)
	fmt.Println(string(data))

	fmt.Println()
	fmt.Println("================================================")
	fmt.Println()

	/** COMPLEX QR*/
	complexQRPay := qrpay.BuildQRPay(
		constants.Bin_VCB,
		"0881000458086",
		qrpay.WithAmount(69000),
		qrpay.WithCurrencyCode(fmt.Sprintf("%d", countries.Singapore.Currency())),
		qrpay.WithDescription("Thanh toán mua sữa"),
		qrpay.WithTipAndFeeType(models.PredefinedTip),
		qrpay.WithTipAndFeeAmount(1000),
		qrpay.WithMerchantInfo(
			qrpay.BuildMerchantInfo(
				"Cửa hàng sữa Vinamilk",
				qrpay.WithCategoryCode(constants.GroceryStoresSupermarkets),
				qrpay.WithMerchantCity("Hồ Chí Minh"),
				qrpay.WithMerchantPostalCode("70000"),
			),
		),
	)

	qr, err = complexQRPay.GenerateQRCode()
	if err != nil {
		fmt.Printf("Error generating QR %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("ComplexQR: %s\n", qr)

	qrPay, err = qrpay.ParseQRPay(qr)
	if err != nil {
		fmt.Printf("parse VietQR ERROR: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Amount: %v\n", qrPay.Amount)
	data, _ = json.Marshal(qrPay)
	fmt.Println(string(data))
}
