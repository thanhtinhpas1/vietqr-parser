package validator

import (
	"errors"
	"strconv"

	"github.com/sigurn/crc16"
)

var (
	crcInvalidError = errors.New("qr code was modified, please check")
)

func ValidateQR(qrCode string) error {
	if !verifyCRC(qrCode) {
		return crcInvalidError
	}

	return nil
}

func verifyCRC(qrCode string) (isValid bool) {
	data := qrCode[0 : len(qrCode)-4]
	crcSig := qrCode[len(data):]
	table := crc16.MakeTable(crc16.CRC16_CCITT_FALSE)
	crcComputed := crc16.Checksum([]byte(data), table)
	i, _ := strconv.ParseUint(crcSig, 16, 64)
	crc := crc16.Complete(uint16(i), table)

	return crcComputed == crc
}
