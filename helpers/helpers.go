package helpers

import "strings"

func LeftPad(s string, padStr string, pLen int) string{
	return strings.Repeat(padStr, pLen) + s
}

func LeftPad2Len(s string, padStr string, overallLen int) string{
	var padCountInt int
	padCountInt = 1 + ((overallLen-len(padStr))/len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr)-overallLen):]
}
