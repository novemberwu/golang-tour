package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	
	var result = make(map[string]int)
	for _,word := range strings.Fields(s){
		//_, ok := result[word]
		//if ok{
			result[word] = result[word]+1
		//}else{
		
		//	result[word] = 1
		//}
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
