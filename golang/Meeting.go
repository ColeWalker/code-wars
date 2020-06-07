package kata

import("strings")
import("sort")
import("fmt")
func Meeting(s string) string {
    arr := strings.Split(s, ";")
    people := []string{}
    
    for _,v := range arr{
      last, first := ExtractInfo(v)
      str := fmt.Sprintf("(%s, %s)", last, first)
      people = append(people, str)
    }
    
    sort.Strings(people)
    
    return strings.Join(people, "")
    
}

func ExtractInfo(s string) (string, string){
  arr:= strings.Split(s, ":")
  last:= strings.ToUpper(arr[1])
  first:= strings.ToUpper(arr[0])
  
  return last, first
}

/*
John has invited some friends. His list is:

s = "Fred:Corwill;Wilfred:Corwill;Barney:Tornbull;Betty:Tornbull;Bjon:Tornbull;Raphael:Corwill;Alfred:Corwill";

Could you make a program that

    makes this string uppercase
    gives it sorted in alphabetical order by last name.

When the last names are the same, 
sort them by first name. Last name and 
first name of a guest come in the result 
between parentheses separated by a comma.

So the result of function meeting(s) will be:

"(CORWILL, ALFRED)(CORWILL, FRED)(CORWILL, RAPHAEL)(CORWILL, WILFRED)(TORNBULL, BARNEY)(TORNBULL, BETTY)(TORNBULL, BJON)"

It can happen that in two distinct families 
with the same family name two people have the same 
first name too.

Notes

    You can see another examples in the "Sample tests".


*/