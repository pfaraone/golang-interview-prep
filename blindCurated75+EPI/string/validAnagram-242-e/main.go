package main

import (
    "fmt" 
)

func isAnagram(s string, t string) bool {
    if len(s) != len(t){
        return false
    }
    // var arr[26] int
    var arr2 = new([26]int)
    // var arr3 = make([]int, 0)
    for i:=0; i<len(s);i++{
        s_ascii := int(s[i])
        arr2[s_ascii-int('a')]++
    }
    for i:=0;i<len(t);i++{
        t_ascii := int(t[i])
        arr2[t_ascii-int('a')]--
        if arr2[t_ascii-int('a')] < 0{
            return false
        }
        // fmt.Printf("char in t is %c \n ",t[i])
    }
    // fmt.Println(arr)
    // fmt.Println(*arr2)
    // fmt.Println(arr3)
    return true
}

func main(){
    fmt.Println(isAnagram("anagram","nagaram"))
    fmt.Println(isAnagram("rat","car"))
}