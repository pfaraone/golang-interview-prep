import (
	"fmt"
	"strings"
	)
	
	func isValid(char byte) bool{
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9'){
			return true
		}
		return false
	}
	
	func isPalindrome(s string) bool {
		if len(s) < 2{
			return true
		}
		lo := 0
		hi := len(s) - 1
		for lo < hi{
			if !isValid(s[lo]){
				lo ++
				continue
			}
			if !isValid(s[hi]){
				hi --
				continue
			}
			if strings.ToLower(string(s[lo])) != strings.ToLower(string(s[hi])){
				return false
			}
			lo ++
			hi -- 
		}
		return true
	}