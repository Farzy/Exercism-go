package cipher

import "strings"

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.
type shift int

type vigenere string

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance < -25 || distance == 0 || distance > 25 {
		return nil
	}

	return shift(distance)
}

func modulo(a, b rune) rune {
	return (a%b + b) % b
}

func (c shift) Encode(input string) string {
	var out strings.Builder
	out.Grow(len(input))

	for _, chr := range input {
		if 'a' <= chr && chr <= 'z' {
			chr = 'a' + modulo(chr-'a'+rune(c), 26)
			out.WriteRune(chr)
		}
		if 'A' <= chr && chr <= 'Z' {
			chr = 'a' + modulo(chr-'A'+rune(c), 26)
			out.WriteRune(chr)
		}
	}
	return out.String()
}

func (c shift) Decode(input string) string {
	return shift(-c).Encode(input)
}

func NewVigenere(key string) Cipher {
	if len(key) == 0 {
		return nil
	}
	aCount := 0
	for _, chr := range key {
		if chr < 'a' || chr > 'z' {
			return nil
		}
		if chr == 'a' {
			aCount++
		}
	}
	if aCount == len(key) {
		return nil
	}
	return vigenere(key)
}

func (v vigenere) Encode(input string) string {
	var out strings.Builder
	out.Grow(len(input))

	keyPosition := 0
	for _, chr := range input {
		if 'a' <= chr && chr <= 'z' {
			chr = 'a' + modulo(chr-'a'+rune(v[keyPosition]-'a'), 26)
			out.WriteRune(chr)
			keyPosition = (keyPosition + 1) % len(v)
		}
		if 'A' <= chr && chr <= 'Z' {
			chr = 'a' + modulo(chr-'A'+rune(v[keyPosition]-'a'), 26)
			out.WriteRune(chr)
			keyPosition = (keyPosition + 1) % len(v)
		}
	}
	return out.String()
}

func (v vigenere) Decode(input string) string {
	var out strings.Builder
	out.Grow(len(input))

	keyPosition := 0
	for _, chr := range input {
		if 'a' <= chr && chr <= 'z' {
			chr = 'a' + modulo(chr-'a'-rune(v[keyPosition]-'a'), 26)
			out.WriteRune(chr)
			keyPosition = (keyPosition + 1) % len(v)
		}
		if 'A' <= chr && chr <= 'Z' {
			chr = 'a' + modulo(chr-'A'-rune(v[keyPosition]-'a'), 26)
			out.WriteRune(chr)
			keyPosition = (keyPosition + 1) % len(v)
		}
	}
	return out.String()
}
