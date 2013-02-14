package main

import (
	"cs553s2013/mylexer"
	"github.com/proebsting/cs553s2013/lexer"
	"fmt"
	"io/ioutil"
)

func main( ){
        var input string
	//fmt.Println("Length of Args: " , len(os.Args))
  /*      if len(os.Args) == 1 {
		input = ""
	}	else {
		input = os.Args[1]
	}
	fmt.Println("Actual arguments: ", input) */
        var output chan <-lexer.Token
        output = make(chan lexer.Token, 1000)
	//fmt.Println("Input: ", input)
	dat, err := ioutil.ReadFile("/home/wallace/test_strings.txt")
 	if ((err == nil) && (dat != nil)) {
		fmt.Println("Try adding this functionality")
	}
	input = string(dat)
	fmt.Println("Value of input is ", input)
        mylexer.Lexer(input, output)
}

