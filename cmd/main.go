package main

import (
	"bigint/pkg/bigint"
	"fmt"
)

func main() {
	a, err := bigint.NewInt("988847123412385995937737458959")
	if err != nil {
		panic(err)
	}

	b, err := bigint.NewInt("21231231231231231231231231233")
	if err != nil {
		panic(err)
	}

	err = b.Set("1") // b = "1"
	if err != nil {
		panic(err)
	}

	c := bigint.Add(a, b)      // c = "988847123412385995937737458960"
	d := bigint.Sub(a, b)      // d = "988847123412385995937737458958"
	e := bigint.Multiply(a, b) // e = "988847123412385995937737458959"
	f := bigint.Mod(a, b)      // f = "0"
	l := b.Abs()               // always returns pos num

	fmt.Println(c, d, e, f, l)
	// fmt.Println(a, b, c, d, e, l)
}
