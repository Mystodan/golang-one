package assignment1

/**
 *  Test Validity
 *  takes the string as an input, and returns boolean flag true if the given string complies with the format,
 *  or false if the string does not comply
 *
 *  @param inn  the string it takes in as an input
 */
func TestValidity(inn string) bool {
	if len(inn) < 1 { // if string is empty
		return false
	}
	if inn[0] == '-' || inn[len(inn)-1] == '-' { // if string starts or ends with '-' dash
		return false
	}
	var i int
	j := 0 // counter for bool array containing the states of the strings
	isNum := make([]bool, len(inn))

	for i = 0; i < len(inn); i++ {
		if inn[i] != '-' { // skips check for '-'
			isNum[j] = true // automatically sets it to number
			if !(uint(inn[i]) >= 48 && uint(inn[i]) <= 57) {
				isNum[j] = false // switches if it isnt a number
			}
		} else {
			j++ // changes check to next string
		}
	}
	if isNum[0] == false { // if first string is NAN
		return false
	}
	for i = 0; i < j; i++ { // compares the string if before the '-' dash and after are the same
		if isNum[i] == isNum[i+1] {
			return false
		}
	}
	return true
}
