package main

import (
	"cs553s2013/mylexer"
	"github.com/proebsting/cs553s2013/lexer"
	"testing"
)

func tozama_test(t *testing.T) {
	c := make(chan lexer.Token, 100)
	s := ``
	var tok mylexer.Token
	v := []lexer.Token{
	      mylexer.Token{"<EOF>", 0, 1, 1, lexer.EOF},
	}
	go mylexer.Lexer(s, c)
	lexer.CheckTest(s, c, v, t)
}

func main(){
    var *t testing.T
    tozama_test(t)
    
}
