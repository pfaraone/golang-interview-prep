package main

func isValid(s string) bool {
	if len(s) < 2 {
		return false
	}
	bracketDict := map[string]string{")": "(",
		"}": "{",
		"]": "["}
	bracketStack := make([]string, 0)
	for char := range s {
		stringVersion := string(char)
		bracket, inMap := bracketDict[stringVersion]
		// case : opening bracket -> push onto stack
		if !inMap {
			bracketStack = append(bracketStack, stringVersion)
		} else {
			// pop off stack most recent pushed on open bracket
			var lastOpeningBracket string
			if len(bracketStack) > 0 {
				lastOpeningBracket = bracketStack[len(bracketStack)-1]
				bracketStack = bracketStack[:len(bracketStack)-1]
			} else {
				lastOpeningBracket = "a"
			}
			if bracket != string(lastOpeningBracket) {
				return false
			}
		}
		// fmt.Println(bracketStack,i,string(char))
	}
	// fmt.Println(bracketStack)
	if len(bracketStack) == 0 {
		return true
	} else {
		return false
	}
}

func main() {

}
