package mylexer

import (
	  "github.com/proebsting/cs553s2013/lexer"
	 //"fmt"
  )


func Lexer(input string, out chan<-lexer.Token){
       var output string
       var token Token
       follow :=0
      ext := 0
      match_star := false
       length := len(input)
       lines := 1 
       columns := 1
       if length == 0 {
	    token = Token{value:"<EOF>", location:0, line:1, column:1, enum:lexer.EOF}; out <- token
	    return
      }
      
      i := 0
      for i < length {
	 for input[i] == 10 || input[i] == ' ' || input[i] == 9 || input[i] == 13{
                if input[i] == 10 {
                        lines++; columns = 1; i++
                }       else {
                        columns++; i++
                }
                if !(i < length){
			token = Token{value:"<EOF>", location:0, line: lines, column: columns, enum:lexer.EOF}; out<-token
			return
		}
        }
	 if ! (i < length) {
	      token = Token{value:"<EOF>", location:0, line: lines, column: columns, enum:lexer.EOF}; out<-token
		return
	 }
	 tok, size, match := lexer.Operator(input[i:])
	 if  IsALetter(input[i]) {
	    output = Reading_Identifier(input[i:])
	    lexed, paired := lexer.Keyword(output)
	    if paired {
		  token = Token{value: output, location:0, line: lines, column: columns, enum:lexed}; out<- token
	    } else {
	    token = Token{value: output, location:0, line: lines, column: columns, enum:lexer.IDENT}; out<-token
	    }
	    columns += len(output)
	    i += len(output)
	    continue
	 } else if IsADigit(input[i]) {
	      output = IsAnId(input[i])
	      if len(input[i:]) < 2 {
		token = Token{value: output, location:0, line: lines, column: columns, enum:lexer.INTEGER}; out<-token
		columns++
	        token = Token{value:"<EOF>", location:0, line: lines, column: columns, enum:lexer.EOF}; out<-token
	        return
	      }
	      number, bel := Reading_Number(input[i:]); output = number
	      token = Token{value: output, location:0, line: lines, column: columns, enum:bel}; out <- token
	      columns += len(output); i += len(output)
	      continue
	  } else if input[i] == '"' {
	    //we are reading a string
	    if len(input[i:]) < 2 {
	      token = Token{value: "\"", location:0, line:lines, column: columns, enum:lexer.ERROR}; out <- token
	      return
	    }
	    output = Reading_String(input[i+1:])
	    j := 0
	    error_flag := 0
		for j < len(output) {
		      if output[j] == '"' {
			error_flag++
		      }
		      j++
		}

		if error_flag == 1 {
		    token = Token{value: output, location:0, line:lines, column: columns, enum:lexer.ERROR}; out <- token
		    return
		}
		token = Token{value: output, location:0, line: lines, column: columns, enum:lexer.STRING}; out <- token
		j = 0
			for j < len(output) {
			    if output[j] == 10 {
				  lines++; columns = 1
			    } else {
			      columns++ 
			    }
			    j++
			}	
		i += len(output)
		continue
	  } else if  match {
	  if i < (len(input) - 1){
		    follow, ext, match_star = lexer.Operator(input[i+1:])
		}    
	  if (tok == lexer.LPAREN) && follow == lexer.STAR {
			if match_star {
			ext++
			}
			output = Reading_Comment(input[i:])
			if output == "bad comment" {
			  token = Token{value: output, location:0, line:lines, column: columns, enum:lexer.ERROR}; out<-token
			  return
			}
			j := 0
			for j < len(output) {
			    if output[j] == 10 {
				  lines++; columns = 1
			    }  else {
			      columns++ 
			    }
			    j++
			}
			i += len(output)
			if i == length {
			    token = Token{value:"<EOF>", location:0, line:lines, column:columns, enum:lexer.EOF}; out <- token
			    return
			}
	 } else {
		output = input[i:i+size]
		token = Token{value: output, location:0, line: lines, column: columns, enum:tok}; out <- token
		columns += size; i += size
		if i == length {
			    token = Token{value:"<EOF>", location:0, line:lines, column:columns, enum:lexer.EOF}; out <- token
			    return
			}
		continue
	    }
	   
	} else {
	    token = Token{value:"-1", location:0, line:lines, column:columns, enum:lexer.ERROR}; out <- token
	    return
	}
	  
	  
      }
      
      token = Token{value:"<EOF>", location:0, line:lines, column:columns, enum:lexer.EOF}; out <- token
}
      
       
       