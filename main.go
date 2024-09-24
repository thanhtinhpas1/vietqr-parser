package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/thanhtinhpas1/emvco_qr/pkg/constants"
	"github.com/thanhtinhpas1/emvco_qr/qrpay"
)

func main() {
	basicQRPay := qrpay.BuildBasicQRPay(20000, constants.Bin_VCB, "0881000458086")
	qr, err := basicQRPay.GenerateQRCode()
	if err != nil {
		fmt.Println("generate qr code failed", err)
		os.Exit(1)
	}

	fmt.Printf("QR Code: %s\n", qr)

	qrPay, err := qrpay.ParseQRPay(qr)
	if err != nil {
		fmt.Printf("parse VietQR ERROR: %v", err)
		return
	}

	fmt.Printf("Amount: %v\n", qrPay.Amount)
	data, _ := json.Marshal(qrPay)
	fmt.Println(string(data))
}
