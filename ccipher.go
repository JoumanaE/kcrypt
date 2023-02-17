package ccypher

import (
	"bytes"
	"unicode"
)

type Cipher struct {
	Key int
}

func (c *Cipher) Encipher(s string) string {
	var b bytes.Buffer
	var ciphered string
	for _, r := range s {
		r = ShiftRune(r, 1)
		b.WriteRune(r)
		ciphered = b.String()
	}
	return ciphered
}

func ShiftRune(r rune, Shift int) rune {
	if !unicode.IsLetter(r) {
		return r
	} else {
		r = unicode.ToUpper(r)
	}
	if r+rune(Shift) >= 'Z' {
		return r + rune(Shift) - 26
	}
	return r + rune(Shift)
}

func (c *Cipher) Decipher(s string) string {
	var deciphered string
	var b bytes.Buffer
	for _, r := range s {
		r = ShiftRune(r, -1)
		b.WriteRune(r)
		deciphered = b.String()
	}
	return deciphered
}

func New(key int) *Cipher {
	return &Cipher{Key: key}
}
