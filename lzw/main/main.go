package main

import (
	"lzw/lzw"
	"fmt"
)


func main() {
	//str := "ABABABA"
	//str := "ABRACADABRABRABRA"
	str := "123 qwe NI9 2^32"
	fmt.Println("The string being compressed :", str)

	fmt.Print("The code generated :")
	lzw.Compress(str)

	//codeInput := []int16 {65, 66, 129, 131, 128}
	//codeInput := []int16 {65, 66, 82, 65, 67, 65, 68, 129, 131, 130, 136, 65, 128}
	codeInput := []int16 {49, 50, 51, 32, 113, 119, 101, 32, 78, 73, 57, 32, 50, 94, 51, 50, 128}
	fmt.Print("The string expanded from the code:")
	lzw.Expand(codeInput)
}
