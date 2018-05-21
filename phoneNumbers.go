//This program takes phone numbers in 4 different formats, and converts them into bare number format, or standard format. 
//Another function returns the area code (first three digits).

package main

import "fmt"
import "strings"
import "unicode"

type PhoneNo interface {
	Number () string
}

type Type1 string

func (t1 Type1) Number() string {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	s1 := strings.FieldsFunc(string(t1), f)
	s := strings.Join(s1, "")
	var b []byte
	b = []byte(s)
	b1 := b[1:]
	s2 := string(b1)
	return s2
}

type Type2 string

func (t2 Type2) Number() string {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	s1 := strings.FieldsFunc(string(t2), f)
	s := strings.Join(s1, "")
	return s
}

type Type3 string

func (t3 Type3) Number() string {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	s1 := strings.FieldsFunc(string(t3), f)
	s := strings.Join(s1, "")
	var b []byte
	b = []byte(s)
	b1 := b[1:]
	s2 := string(b1)
	return s2
}

type Type4 string

func (t4 Type4) Number() string {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	s1 := strings.FieldsFunc(string(t4), f)
	s := strings.Join(s1, "")
	return s
}

func Format (s string) string {
	var b []byte
	b = []byte(s)
	var num1 []byte
	var num2 []byte
	var num3 []byte
	num1 = b[0:3]
	num2 = b[3:6]
	num3 = b[6:10]
	s1 := "(" + string(num1) + ") " + string(num2) + "-" + string(num3)
	return s1
}

func AreaCode (s string) string {
	var b []byte
	b = []byte(s)
	num := b[0:3]
	return string(num)
}

func RetType (s string) string {
	var b []byte
	b = []byte(s)
	s1 := string(b[0:2])
	if s1 == "+1" {
		return "Type1"
	} else if string(b[3:4]) == "-" && string(b[7:8]) == "-"{
		return "Type2"
	} else if string(b[1:2]) == " " && string(b[5:6]) == " " && string(b[9:10]) == " " {
		return "Type3"
	} 
	return "Type4"
	
}

func main () {
	var numbers  = []string{"613-995-0253", "1 613 995 0253", "+1 (613)-995-0253", "613.995.0253"}
	phoneDirectory := make([]PhoneNo, len(numbers))
	
	for i := 0; i < len(numbers); i++ {
		d1 := RetType(numbers[i])
		if strings.Compare(d1, "Type1") == 0 {
			phoneDirectory[i] = Type1(numbers[i])
		} else if strings.Compare(d1, "Type2") == 0 {
			phoneDirectory[i] = Type2(numbers[i])
		} else if strings.Compare(d1, "Type3") == 0 {
			phoneDirectory[i] = Type3(numbers[i])
		} else if strings.Compare(d1, "Type4") == 0 {
			phoneDirectory[i] = Type4(numbers[i])
		}
	}
	
	for i := 0; i < len(phoneDirectory); i++ {
		a := phoneDirectory[i]
		fmt.Println(a.Number())
		fmt.Println(Format(a.Number()))
		fmt.Println(AreaCode(a.Number()))
		fmt.Printf("%T\n\n", a)
	}
}
