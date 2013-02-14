package mylexer

import (
	  "github.com/proebsting/cs553s2013/lexer"
      )

type Token struct{	
	value string
	location int
	line int
	column int
	enum int
}

func (t Token) Value() string{
	return t.value
}

func (t Token) Location() int{
	return t.location
}

func (t Token) Line() int{
	return t.line
}

func (t Token) Column() int{
	return t.column 
}

func (t Token) Enum() int{
        return t.enum
}

func Reading_Identifier(input string) string {
      i:= 0
      var output string
      length := len(input) 
      for (IsALetter(input[i]) || IsADigit(input[i])) {
	 output += IsAnId(input[i])
	 if len(output) < length {
	      i++
	 } else {
	    break;
	 }
      }
      return output
}

func Reading_Number(input string) (string, int) {
     i := 0
     var output string
     size := len(input)
     dots := 0
     letters := 0
     var exit bool
     exit = false
     for i < size {
	  switch(input[i]) {
	    case '1': output += IsAnId(input[i])
		      break;
	    case '2': output += IsAnId(input[i])
		      break;
	    case '3': output += IsAnId(input[i])
		      break;
	    case '4': output += IsAnId(input[i])
		      break;
	    case '5': output += IsAnId(input[i])
		      break;
	    case '6': output += IsAnId(input[i])
		      break;
	    case '7': output += IsAnId(input[i])
		      break;
	    case '8': output += IsAnId(input[i])
		      break;
	    case '9': output += IsAnId(input[i])
		      break;
	    case '0': output += IsAnId(input[i])
		      break;
	    case 'A': if dots == 1{
			  exit = true
			  break
		      }
		      output += IsAnId(input[i])
		      letters++
		      break;
	    case 'B': if dots == 1{
			  exit = true
			  break
		      }
		      output += IsAnId(input[i])
		      letters++
		      break;
	    case 'C': if dots == 1{
			  exit = true
			  break
		      }
		      output += IsAnId(input[i])
		      letters++
		      break;
	    case 'D': if dots == 1{
			  flag := false
			  //means appending to a real number
			  j := 0; for j < len(output) { 
			      if (output[j] == 'E') || output[j] == 'D' {
				flag = true
				break
				}
				j++
			    }
			    if flag { exit = true; break }
		      }
		      output += IsAnId(input[i])
		      letters++
		      break
	    case 'E': if dots == 1{
			  flag :=false			  
			  j := 0; for j < len(output) { 
			      if output[j] == 'E' || output[j] == 'D' {
				flag = true
				}
				j++
			    }
			  if flag == true { exit = true; break }
		      }
		      output += IsAnId(input[i])
		      letters++
		      break;
	    case 'F': if dots == 1{
			  exit = true
			  break
		      }
		      output += IsAnId(input[i])
		      letters++
		      break; 
	    case 'H': if dots == 1 { exit = true; break }
		      exit = true;
		      output += "H"
		      letters++
		      break;
	    case 'X': if dots == 1 {
			exit = true;
			break
		      }
		      letters++
		      exit = true
		      output += "X"
		      break
	    case '+': length := len(output)
		      if !((IsAnId(output[length - 1]) == "E")  || (IsAnId(output[length - 1]) == "D")) {
			//means we are done reading real number
			exit = true
			break
		      }
		      output += "+"
		      break
	    case '-': length := len(output)
		      if !((IsAnId(output[length - 1]) == "E") || (IsAnId(output[length - 1]) == "D")) {
			//means we are done reading real number
			exit = true
			break
		      }
		      output += "-"
		      break
	    case '.': if dots != 0 || letters > 0{
			  exit = true
		      } else {
			dots++
			output += "."
		      }
			
	    default: exit = true
		    }	
		  i++
		  if exit {
		      if (dots == 0 && letters == 0) {
			//means that we are returning an integer type
			return output, lexer.INTEGER
		      } else if dots == 1{
			//possible reading of real type
			if output[len(output) - 1] == '.' || output[len(output) - 1] == 'D' || output[len(output) - 1] == 'E' || output[len(output) - 1] == '+'|| output[len(output) - 1] == '-' {
			    return output, lexer.ERROR
			} else {
			    return output, lexer.REAL
			}
		      } else {
			  //means there is at least one letter in input
			  if output[len(output) - 1] == 'H' {
			      return output, lexer.INTEGER
			  } else if output[len(output) - 1] == 'X'{
			      return output, lexer.STRING
			  } else {
			      return output, lexer.ERROR
			  }
			  
		      }
		      
		  }
	    }
	    //at end of input
	    if dots == 1 { if output[len(output) - 1] == '.' || output[len(output) - 1] == 'D' || output[len(output) - 1] == 'E' || output[len(output) - 1] == '+'|| output[len(output) - 1] == '-' {
			    return output, lexer.ERROR
			} else {
			    return output, lexer.REAL
			}
	    }
	    if output[len(output) - 1] == 'H' {
			      return output, lexer.INTEGER
			  } else if output[len(output) - 1] == 'X'{
			      return output, lexer.STRING
			  }
	    return output, lexer.ERROR
} 

func Reading_String(input string) string{
	i := 0
	size := len(input)
	var output string
	output = "\""
	for i < size {
	      if (input[i] == '"')  {
		output += input[0:i+1]
		return output
	      }
	      i++
	}
	output += input[0:]
	return output
}

func Reading_Comment(input string) string{
      credit := 1
      length := len(input)
      if length == 2 {
	//means we return error token
	return input
      }
      i := 2
      for i < length {
	if input[i] == '*' {
	      if (i+1) < length {
		if input[i+1] == ')' {
		    //means we have possibly found our match
		    credit--
		    if credit == 0 {
			  return input[:i+2]
		    }
		}
	      }
	} else if input[i] == '(' {
		if (i+1) < length {
		  if input[i+1] == '*' {
		      //means we are reading nested comments
		      credit++
		  }
		}
	} 
		i++

      }
	//getting to this region of code means we have a possible debt
	//so we return an error token
	return "bad comment"
}

func IsALetter(literal uint8) bool {
      switch(literal){
		case 'a' : return true
		case 'b' : return true 
		case 'c' : return true
		case 'd' : return true
		case 'e' : return true
		case 'f' : return true
		case 'g' : return true
		case 'h' : return true
		case 'i' : return true
		case 'j' : return true
		case 'k' : return true
		case 'l' : return true
		case 'm' : return true
		case 'n' : return true
		case 'o' : return true
		case 'p' : return true
		case 'q' : return true
		case 'r' : return true
		case 's' : return true
		case 't' : return true
		case 'u' : return true
		case 'v' : return true
		case 'w' : return true  
		case 'x' : return true
		case 'y' : return true
		case 'z' : return true
		case 'A' : return true
		case 'B' : return true
		case 'C' : return true
		case 'D' : return true
		case 'E' : return true
		case 'F' : return true
		case 'G' : return true
		case 'H' : return true
		case 'I' : return true
		case 'J' : return true
		case 'K' : return true
		case 'L' : return true
		case 'M' : return true
		case 'N' : return true
		case 'O' : return true
		case 'P' : return true
		case 'Q' : return true
		case 'R' : return true
		case 'S' : return true
		case 'T' : return true
		case 'U' : return true
		case 'V' : return true
		case 'W' : return true
		case 'X' : return true
		case 'Y' : return true
		case 'Z' : return true
		
      }
      return false;
}

func IsADigit(literal uint8) bool{
    switch(literal){
		case '1': return true
		case '2': return true
		case '3': return true
		case '4': return true
		case '5': return true
		case '6': return true
		case '7': return true
		case '8': return true
		case '9': return true
		case '0': return true
		}
	 return false

}

func IsAnId(literal uint8) string {
	switch literal {
		case 'a' : return "a"
		case 'b' : return "b" 
		case 'c' :return "c"
		case 'd' :return "d"
		case 'e' :return "e"
		case 'f' :return "f"
		case 'g' :return "g"
		case 'h' :return "h"
		case 'i' :return "i"
		case 'j' :return "j"
		case 'k' :return "k"
		case 'l' :return "l"
		case 'm' :return "m"
		case 'n' :return "n"
		case 'o' :return "o"
		case 'p' :return "p"
		case 'q' :return "q"
		case 'r' :return "r"
		case 's' :return "s"
		case 't' :return "t"
		case 'u' :return "u"
		case 'v' :return "v"
		case 'w' :return "w"
		case 'x' :return "x"
		case 'y' :return "y"
		case 'z' :return "z"
		case 'A' :return "A"
		case 'B' :return "B"
		case 'C' :return "C"
		case 'D' :return "D"
		case 'E' :return "E"
		case 'F' :return "F"
		case 'G' :return "G"
		case 'H' :return "H"
		case 'I' :return "I"
		case 'J' :return "J"
		case 'K' :return "K"
		case 'L' :return "L"
		case 'M' :return "M"
		case 'N' :return "N"
		case 'O' :return "O"
		case 'P' :return "P"
		case 'Q' :return "Q"
		case 'R' :return "R"
		case 'S' :return "S"
		case 'T' :return "T"
		case 'U' :return "U"
		case 'V' :return "V"
		case 'W' :return "W"
		case 'X' :return "X"
		case 'Y' :return "Y"
		case 'Z' :return "Z"
		case '1' :return "1"
		case '2' :return "2"
		case '3' :return "3"
		case '4' :return "4"
		case '5' :return "5"
		case '6' :return "6"
		case '7' :return "7"
		case '8' :return "8"
		case '9' :return "9"
		case '0' :return "0"
		default  ://fall through 
	}
		return ""
}