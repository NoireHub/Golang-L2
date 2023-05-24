package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)


func stringUnpack(s string) (string, error){
	if len(s) == 0 {
		return "", nil
	}

	if '0' <= s[0] && s[0] <= '9' {
		return "" , errors.New("invalid string")
	}

	var sBuilder strings.Builder
	var prevCh rune
	var escape bool

	for _, ch := range s {

		if ch == '\\' {
			if escape {
				sBuilder.WriteRune(ch)
				escape = false
			}else {
				escape = true
			}

			prevCh = ch	
			continue
		}

		if '0' <= ch && ch <= '9' && !escape {	
			num, err:= strconv.Atoi(string(ch))
			if err != nil || num < 1 {
				return "" , errors.New("invalid number")
			}
			

			sBuilder.WriteString(strings.Repeat(string(prevCh),num - 1))
		}else{
			escape = false
			sBuilder.WriteRune(ch)	
		}

		prevCh = ch		
	}
 
	return sBuilder.String() , nil
}

func main() {
	
	fmt.Println(stringUnpack("a4bc2d0e"))
	fmt.Println(stringUnpack("qwe\\4\\5"))
	fmt.Println(stringUnpack("a4bc2d5e"))
	fmt.Println(stringUnpack("qwe\\\\5"))
}
