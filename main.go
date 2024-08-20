package main

import (
	"encoding/json"
	"fmt"

	"github.com/thanhtinhpas1/emvco_qr/qrpay"
)

func main() {
	qrPay, err := qrpay.ParseQRPay("0002010102111531272007040052044600000006215014926480708804357610010vn.zalopay0111ZP-89849ED2020300150430011SACOMBANKQR011621129977623751279904A769520458125303704550202560105802VN5910ZLP*MBHIOT6006QUAN 7623307088043576101172312180000000000463044798")
	if err != nil {
		fmt.Printf("parse VietQR ERROR: %v", err)
		return
	}

	fmt.Printf("Amount: %v\n", qrPay.Amount)
	data, _ := json.Marshal(qrPay)
	fmt.Println(string(data))
}
