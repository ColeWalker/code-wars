package kata

func ValidParentheses(parens string) bool {
    open:= 0
    
    chars := []rune(parens)
    
    for _,v := range chars{
      if(v==')'){
        open --
        if(open < 0){
          return false
        }
      } else {
        open++
      }
    }
    return open == 0
}

/*
Write a function called that takes a string of parentheses, 
and determines if the order of the parentheses is valid. 
The function should return true if 
the string is valid, and false if it's invalid.
*/