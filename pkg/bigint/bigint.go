package bigint

import (
	"errors"
	"math/big"
	"strings"
)

type Bigint struct {
	Value string
}

var (
	ErrorBadInput = errors.New("bad input, please input only number")
)

func Validation(num string) bool {
	allowed := "1234567890-+"
	var err bool

	arr := strings.Split(num, "")
	for _, v := range arr {
		if !strings.Contains(allowed, v) {
			err = true
		}
	}

	if err {
		return true
	} else {
		return false
	}
}

func CleanUp(num string) (string, bool) {
	verify := "+-"

	arr := strings.Split(num, "")
	for i := 1; i < len(arr); i++ {
		if strings.Contains(verify, arr[i]) {
			return num, true
		}
	}
	if strings.HasPrefix(num, "+") {
		num = strings.Replace(num, "+", "", 1)
	}
	if strings.HasPrefix(num, "0") {
		num = strings.TrimLeft(num, "0")
	}
	return num, false
}

func NewInt(num string) (Bigint, error) {
	wrongVal := Validation(num)
	cleanedVal, err := CleanUp(num)

	if err || wrongVal {
		return Bigint{Value: num}, ErrorBadInput
	} else {
		return Bigint{Value: cleanedVal}, nil
	}
}

func (z *Bigint) Set(num string) error {
	wrongVal := Validation(num)
	cleanedVal, err := CleanUp(num)

	if err || wrongVal {
		return ErrorBadInput
	}
	z.Value = cleanedVal
	return nil
}

func isHexString(s string) bool {
	return strings.ContainsAny(s, "abc")
}

func toBigInt(s string) *big.Int {
	x, _ := new(big.Int).SetString("0", 16)

	if strings.HasPrefix(s, "0x") {
		s = strings.Replace(s, "0x", "", 1)
		x, _ = new(big.Int).SetString(s, 16)
	} else if strings.HasPrefix(s, "0") {
		x, _ = new(big.Int).SetString(s, 8)
	} else if isHexString(s) {
		x, _ = new(big.Int).SetString(s, 16)

	} else {
		x, _ = new(big.Int).SetString(s, 10)

	}
	return (x)
}

func Add(a, b Bigint) Bigint {
	x1 := toBigInt(a.Value)
	x2 := toBigInt(b.Value)
	result := new(big.Int).Add(x1, x2).String()
	return Bigint{Value: result}
}

func Sub(a, b Bigint) Bigint {
	x1 := toBigInt(a.Value)
	x2 := toBigInt(b.Value)
	result := new(big.Int).Sub(x1, x2).String()
	return Bigint{Value: result}
}

func Multiply(a, b Bigint) Bigint {
	x1 := toBigInt(a.Value)
	x2 := toBigInt(b.Value)
	result := new(big.Int).Mul(x1, x2).String()
	return Bigint{Value: result}
}

func Mod(a, b Bigint) Bigint {
	x1 := toBigInt(a.Value)
	x2 := toBigInt(b.Value)
	result := new(big.Int).Mod(x1, x2).String()
	return Bigint{Value: result}
}

func (x Bigint) Abs() Bigint {
	if x.Value[0] == '-' {
		return Bigint{
			Value: x.Value[1:],
		}
	}
	if x.Value[0] == '+' {
		return Bigint{
			Value: x.Value[1:],
		}
	}
	return Bigint{
		Value: x.Value,
	}
}
