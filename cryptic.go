package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func wrapAndPrintText(text string, caption string, width int) {

	wrappedtextarray := WrapString(text, width)
	for _, wrappedtext := range wrappedtextarray {
		fmt.Printf("%-16s%s\n", caption, wrappedtext)
		caption = ""
	}
}

func WrapString(s string, maxline int) []string {
	totallength := len(s)
	linesneeded := totallength / maxline
	extralines := totallength % maxline
	if 0 < extralines {
		linesneeded++
	}

	wrapped := make([]string, linesneeded)

	for i := 0; i < linesneeded; i++ {
		start := i * maxline
		end := start + maxline
		if end < totallength {
			wrapped[i] = s[start:end]
		} else {
			wrapped[i] = s[start:]
		}
	}

	return wrapped
}

func main() {

	//The key argument should be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
	key := "abcdefghijklmnopqrstuvwxyz012345" // 16 bytes!

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d bytes NewCipher key with block size of %d bytes\n", len(key), block.BlockSize)
	plaintext := []byte("D4.99999599999999991100119911QR84084020165    007055999Y002554@D4012000010000=15121010004852000VisaNet Merchant         Sterling     VA020000136B248")

	// 16 bytes for AES-128, 24 bytes for AES-192, 32 bytes for AES-256
	ciphertext := []byte("abcdefghijklmnopqrstuvwxyz012345")
	iv := ciphertext[:aes.BlockSize] // const BlockSize = 16

	// encrypt
	encrypter := cipher.NewCFBEncrypter(block, iv)
	encrypted := make([]byte, len(plaintext))
	encrypter.XORKeyStream(encrypted, plaintext)

	wrapAndPrintText(fmt.Sprintf("%s", plaintext), "plaintext ->", 96)
	fmt.Println()
	wrapAndPrintText(fmt.Sprintf("%x", encrypted), "encrypted ->", 96)
	fmt.Println()

	// decrypt
	decrypter := cipher.NewCFBDecrypter(block, iv) // simple!
	decrypted := make([]byte, len(plaintext))
	decrypter.XORKeyStream(decrypted, encrypted)

	wrapAndPrintText(fmt.Sprintf("%s", decrypted), "decrypted ->", 96)

}
