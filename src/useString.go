package main

import (
	"fmt"
	s "strings"
	"unicode"
)

func main() {
	fmt.Printf("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHAlis"))
	fmt.Printf("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHAli"))

	fmt.Printf("Index: %v\n", s.Index("Mihalis", "ha"))
	fmt.Printf("Index: %v\n", s.Index("Mihalis", "Ha"))

	fmt.Printf("Prefix: %v\n", s.HasPrefix("Mihalis", "Mi"))
	fmt.Printf("Prefix: %v\n", s.HasPrefix("Mihalis", "mi"))
	fmt.Printf("Prefix: %v\n", s.HasSuffix("Mihalis", "is"))
	fmt.Printf("Prefix: %v\n", s.HasSuffix("Mihalis", "Is"))

	t := s.Fields("This is a string!")
	fmt.Printf("%v\n", len(t))
	t = s.Fields("ThisIs a\tstring!")
	fmt.Printf("%v\n", len(t))

	fmt.Printf("%s\n", s.Split("abcd efg", ""))
	fmt.Printf("%s\n", s.Replace("abcd efg", "", "_", -1))
	fmt.Printf("%s\n", s.Replace("abcd efg", "", "_", 4))
	fmt.Printf("%s\n", s.Replace("abcd efg", "", "_", 2))

	fmt.Printf("SplitAfter: %s\n", s.SplitAfter("123++432++", "++"))

	trimFunction := func(c rune) bool {
		fmt.Printf("trimFunc Input: %v\n", c)
		fmt.Printf("is Letter: %t\n", unicode.IsLetter(c))
		return !unicode.IsLetter(c)
	}

	fmt.Printf("TrimFunc: %s\n", s.TrimFunc("123 abc ABC \t .", trimFunction))
	trimFuncResult := s.TrimFunc("123 abc ABC \t .", trimFunction)
	fmt.Printf("%T\n", trimFuncResult)
	fmt.Println(trimFuncResult)
}
