//This program takes a string of words, and prints out the acronym for the string.

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func acronym (s string) {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	
	var slice []string
	slice = strings.FieldsFunc(s, f)

	for i := 0; i < len(slice); i++ {
		fmt.Print(string(unicode.ToUpper(rune([]rune(slice[i])[0]))))
	}
	fmt.Println();
}

func main(){
	str := "Indian Standard Time"
	acronym(str)
}
