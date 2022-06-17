package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	Print(os.Args[1])
}

func Print(x string) {
	fmt.Println(ConvertWithSpecials(x))
}

func Convert(i string) string {
	var o []string
	for _, b := range i {
		o = append(o, Cyr2Lat(strings.ToLower(string(b))))
	}
	return strings.Join(o[:], "")
}

func ConvertWithSpecials(i string) string {
	for k, v := range c2lspec {
		c2l[k] = v
	}
	return Convert(i)
}
