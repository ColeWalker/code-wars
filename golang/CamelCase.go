package kata

import("strings")

func ToCamelCase(s string) string {
  var arr []string
  
  if strings.Contains(s, "-") {
    arr = strings.Split(s, "-")
  }else if strings.Contains(s, "_") {
    arr = strings.Split(s, "_")
  }
  
  for i,v := range arr{
    if(i > 0){
      arr[i] = strings.Title(v)
    }
  }
	
	return strings.Join(arr, "")
}

/*
Complete the method/function 
so that it converts dash/underscore 
delimited words into camel casing. 
The first word within the output should 
be capitalized only if the original 
word was capitalized (known 
as Upper Camel Case, also often 
referred to as Pascal case). 
*/