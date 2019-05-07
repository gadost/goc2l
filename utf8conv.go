package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/maxn/utf8conv/constants"
)

func main() {
	x := os.Args[1]
	var answ []string
	for _, byteValue := range x {
		answ = append(answ, constants.Utf(strings.ToLower(string(byteValue))))
	}
	fmt.Println(strings.Join(answ[:], ""))

}
