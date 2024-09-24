package utils

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/sigurn/crc16"
	"golang.org/x/text/unicode/norm"
)

const (
	VNRegex = `ÀÁÂÃÈÉÊÌÍÒÓÔÕÙÚÝàáâãèéêìíòóôõùúýĂăĐđĨĩŨũƠơƯưẠạẢảẤấẦầẨẩẪẫẬậẮắẰằẲẳẴẵẶặẸẹẺẻẼẽẾếỀềỂểỄễỆệỈỉỊịỌọỎỏỐốỒồỔổỖỗỘộỚớỜờỞởỠỡỢợỤụỦủỨứỪừỬửỮữỰự`
)

func ConvertStringToAlphabet(input string) string {
	t := strings.ReplaceAll(input, "đ", "d")
	t = strings.ReplaceAll(t, "Đ", "D")
	normalized := norm.NFD.String(t)
	var result []rune
	for _, char := range normalized {
		if char < '\u0300' || char > '\u036F' {
			result = append(result, char)
		}
	}
	return string(result)
}

func ValidateDescriptionWithRegex(input string) bool {
	pattern := fmt.Sprintf("^[a-zA-Z0-9 _\\-.%s]+", VNRegex)
	regexpPattern, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Printf("%s matches the pattern\n", input)
		return false
	}
	return regexpPattern.MatchString(input)
}

func GetCRC(data string) string {
	crcComputed := crc16.Checksum([]byte(data), crc16.MakeTable(crc16.CRC16_CCITT_FALSE))
	return strings.ToUpper(fmt.Sprintf("%04X", crcComputed))
}
