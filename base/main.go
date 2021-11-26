package main

import (
	"fmt"
)

func main()  {
	a:=strStr("qwewererere","ew")
//	stack=stack[:len(stack)-1]
	fmt.Println(a)
}

func strStr(haystack string, needle string) int {
	if len(haystack) == 0 {
		return 0
	}

	var i,j int
	for i=0;i< len(haystack);i++ {
		for j=0;j<len(needle);j++ {
			if haystack[i+j] != needle[j]{
				break
			}

		}

		if j== len(needle) {
			return i
		}
	}
	return -1



	//
	//for i:=0;i<len(haystack);i++ {
	//	if haystack[i] == needle {
	//
	//	}
	//}
}