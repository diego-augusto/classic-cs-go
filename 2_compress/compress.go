package compress

import "math/bits"

func compress(gene string) (bites uint) {

	bites = 0b01

	for _, s := range gene {

		bites <<= 2

		if string(s) == "A" {
			bites |= 0b00
		} else if string(s) == "C" {
			bites |= 0b01
		} else if string(s) == "G" {
			bites |= 0b10
		} else if string(s) == "T" {
			bites |= 0b11
		}
	}

	return
}

func decompress(number uint) (gene string) {

	for i := 0; i < bits.Len(number)-1; i += 2 {

		bits := number >> i & 0b11

		if bits == 0b00 {
			gene += "A"
		} else if bits == 0b01 {
			gene += "C"
		} else if bits == 0b10 {
			gene += "G"
		} else if bits == 0b11 {
			gene += "T"
		}
	}

	gene = reserve(gene)

	return
}

func reserve(text string) (result string) {

	r := []rune(text)

	for i := len(r) - 1; i >= 0; i-- {
		result += string(r[i])
	}

	return
}
