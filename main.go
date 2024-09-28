package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/biter777/countries"

	"github.com/thanhtinhpas1/emvco_qr/pkg/constants"
	"github.com/thanhtinhpas1/emvco_qr/pkg/models"
	"github.com/thanhtinhpas1/emvco_qr/qrpay"
)

func main() {
	/** BASIC QR */
	basicQRPay := qrpay.BuildQRPay(20000, constants.Bin_VCB, "0881000458086")
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

	/** COMPLEX QR*/
	complexQRPay := qrpay.BuildQRPay(
		50000,
		constants.Bin_VCB,
		"0881000458086",
		qrpay.WithCurrencyCode(countries.Singapore.Currency().String()),
		qrpay.WithDescription("Thanh toán cho cửa hàng sữa"),
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
}
