package utils

import (
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"strconv"
	"time"
)

func StartOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func EndOfDay(t time.Time) time.Time {
	startOfDay := StartOfDay(t)
	return startOfDay.Add(24 * 3600 * time.Second)
}

func FloatToBigInt(val float64, decimal int32) *big.Int {
	bigval := new(big.Float)
	bigval.SetFloat64(val)

	coin := big.NewFloat(math.Pow10(int(decimal)))
	bigval.Mul(bigval, coin)

	result := new(big.Int)
	result.SetString(fmt.Sprintf("%f", bigval), 10)
	return result
}

func I2JsonString(data interface{}) string {
	body, _ := json.Marshal(data)
	return string(body)
}

func Recover() {
	if err := recover(); err != nil {
		fmt.Println("Recovered from panic: ", err)
	}
}

func GenerateCode(length int) string {
	code := ""
	for i := 0; i < length; i++ {
		code = code + strconv.Itoa(rand.Intn(9-0)+0)
	}
	return code
}
