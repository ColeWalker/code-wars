package kata

import("strconv")
import("strings")

func DigitalRoot(n int) int {
  arr := strings.Split(strconv.Itoa(n),"")
  sum := 0
  
  for _,v := range arr{
    value,_ := strconv.Atoi(v)
    sum+= value
  }
  
  if sum >= 10{
    return DigitalRoot(sum)
  } else{
    return sum
  }
}

/*
Digital root is the recursive sum of all the digits in a number.

Given n, take the sum of the digits of n. 
If that value has more than one digit, continue 
reducing in this way until a single-digit number is produced. 
This is only applicable to the natural numbers.

*/