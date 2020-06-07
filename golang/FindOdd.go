package kata


func FindOdd(seq []int) int {
    ret:=0
    m:= make(map[int]int)
    
    for i:=0; i<len(seq); i++ {
      m[seq[i]]++
    }
    
    for k, v := range m {
      if v % 2 > 0 {
        ret = k
      }
    }
    
    return ret
}

/*
Given an array, find the integer that appears an 
odd number of times.

There will always be only one integer that 
appears an odd number of times.






*/