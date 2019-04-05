package golangexamples

import "github.com/ehteshamz/greetings"
import (
	"fmt"
)

func ConcatSlice(sliceToConcat []byte) string{
	concatinated := ""
	for i := 0; i< len(sliceToConcat);i++{
			concatinated = concatinated + string(sliceToConcat[i]) + "-"
	}
	return concatinated
}

func Encrypt(sliceToEncrypt []byte, ceaserCount int){
	encrypted:= ""
	for i := 0; i< len(sliceToEncrypt);i++{
			encrypted = encrypted + string(int(sliceToEncrypt[i]) + ceaserCount)
	}
	fmt.Println(encrypted)

}

func EZGreetings(name string) string{
	return (greetings.PrintGreetings(name))
}
