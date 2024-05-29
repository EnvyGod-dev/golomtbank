package utils

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

func GenerateAccountNumber(bankName string) int64 {
	var prefix string
	switch bankName {
	case "Голомт банк":
		prefix = getRandomPrefix([]string{"12", "11"})
	case "Хаан банк":
		prefix = getRandomPrefix([]string{"50", "51", "52"})
	case "Mbank":
		prefix = getRandomPrefix([]string{"80", "88"})
	case "Төрийн банк":
		prefix = getRandomPrefix([]string{"10"})
	case "Худалдаа хөгжлийн банк":
		prefix = getRandomPrefix([]string{"70"})
	case "Богд Банк":
		prefix = getRandomPrefix([]string{"60"})
	default:
		return 0
	}

	suffix := generateRandomDigits(8)
	accountNumberStr := prefix + suffix

	accountNumber, err := strconv.ParseInt(accountNumberStr, 10, 64)
	if err != nil {
		return 0
	}
	return accountNumber
}

func getRandomPrefix(prefixes []string) string {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(prefixes))))
	if err != nil {
		return ""
	}
	return prefixes[n.Int64()]
}

func generateRandomDigits(length int) string {
	const digits = "0123456789"
	result := make([]byte, length)
	for i := range result {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return ""
		}
		result[i] = digits[n.Int64()]
	}
	return string(result)
}
