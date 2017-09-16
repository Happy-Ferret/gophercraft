package packet

import (
	"bytes"
	"crypto/rand"
)

func reverseBuffer(input []byte) []byte {
	buf := make([]byte, len(input))
	inc := 0
	for x := len(input) - 1; x > -1; x-- {
		buf[inc] = input[x]
		inc++
	}
	return buf
}

func packetString(input string) []byte {
	data := []byte(input)
	data = bytes.Replace(data, []byte("."), []byte{0}, -1)
	return data
}

func randomBuffer(l int) []byte {
	buf := make([]byte, l)
	rand.Read(buf)
	return buf
}
