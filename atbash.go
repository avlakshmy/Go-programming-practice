/*******************************

An implementation of the atbash cipher. The Atbash cipher is a simple substitution cipher that relies on transposing all the letters in the alphabet such that the resulting alphabet is backwards. The first letter is replaced with the last letter, the second with the second-last, and so on.

Plain:  abcdefghijklmnopqrstuvwxyz
Cipher: zyxwvutsrqponmlkjihgfedcba

********************************/

package main

import "fmt"
import "unicode"

func crypt (c byte) byte {
	if unicode.IsLower(rune(c)) {
		return 97 + 122 - c
	}
	if unicode.IsUpper(rune(c)) {
		return 65 + 90 - c
	}
	return c
}

func main() {
	var str string
	str = "test"
	l := len(str)
	
	var arr = make([]byte, l)
	for i := 0; i < l; i++ {
		arr[i] = crypt(str[i])
	}
	arr1 := string(arr)
	fmt.Println(arr1)
}
