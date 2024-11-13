package main

func palindromeChecker(s string) bool {

	if len(s)%2 == 0 {
		for i := 0; i <= (len(s) / 2); i++ {
			if s[i] != s[len(s)-1-i] {
				return false
			}
		}
	}

	if len(s)%2 != 0 {
		for i := 0; i <= ((len(s) - 1) / 2); i++ {
			if s[i] != s[len(s)-1-i] {
				return false 
			}
		}
	}

	return true 
}