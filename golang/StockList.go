package kata

import "fmt"
import "strings"
import "strconv"

func StockList(listArt []string, listCat []string) string {
    m := make(map[string]int)
    
    var retArr []string
    retStr := ""
    isEmpty := true
    
    fmt.Print(listArt)
    for _,v := range listArt {
      category, stock := extractInfo(v, " ")
      m[category]+=stock
      isEmpty=false
    }
    
    for _,v := range listCat {
        retArr= append(retArr, fmt.Sprintf("(%s : %d)", v, m[v]))
    }
    
    if(len(retArr) > 0 && !isEmpty){
      retStr = strings.Join(retArr, " - ")
    }
    return retStr
}

func extractInfo(s, sep string) (string, int) {
    x := strings.Split(s, sep)
    num, _ := strconv.Atoi(x[1])
    return string([]rune(x[0])[0]), num
}

/*
A bookseller has lots of books classified in 26 categories 
labeled A, B, ... Z. Each book has a code c of 3, 4, 5 or 
more capitals letters. The 1st letter of a code is the capital 
letter of the book category. In the bookseller's stocklist 
each code c is followed by a space and by a positive integer n 
(int n >= 0) 
which indicates the quantity of books of this code in stock.





your task is to find all the books of listArt with codes 
belonging to each category of listCat
and to sum their quantity according to each category. 
*/