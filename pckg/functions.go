package assignment1

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

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
	j := 0 // counter for bool array containing the states of the strings
	isNum := make([]bool, uint(len(inn)/3))

	for i := 0; i < len(inn); i++ {
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
	for i := 0; i < j; i++ { // compares the string if before the '-' dash and after are the same
		if isNum[i] == isNum[i+1] {
			return false
		}
	}
	return true
}

/**
 *  Average Number
 *  takes the string, and returns the average number from all the numbers
 *
 *  @param inn  the string it takes in as an input
 */
func AverageNumber(inn string) uint {
	if len(inn) < 1 { // if string is empty
		return 0
	}
	if inn[0] == '-' || inn[len(inn)-1] == '-' { // if string starts or ends with '-' dash
		return 0
	}
	j := 0 // counter for bool array containing the states of the strings
	isNum := make([]bool, uint(len(inn)/3))
	for i := 0; i < len(inn); i++ {
		if inn[i] != '-' { // skips check for '-'
			isNum[j] = true // automatically sets it to number
			if !(uint(inn[i]) >= 48 && uint(inn[i]) <= 57) {
				isNum[j] = false // switches if it isnt a number
			}
		} else {
			j++ // changes check to next string
		}
	}
	strInn := strings.Split(inn, "-")
	if isNum[0] == false { // if first string is NAN
		return 0
	}
	var avgNum uint
	div := uint(0)
	fmt.Print("The average number of the numbers:[")
	for i := 0; i < j; i++ {
		if isNum[i] == true { // if its a number
			tNum, err := strconv.Atoi(strInn[i]) // convert from ascii to integer
			if err != nil {
				fmt.Println(err)
				return 0
			}
			fmt.Print(" ", tNum, " ")
			avgNum += uint(tNum) // adds it to the answer
			div++
		}
	}
	fmt.Print("] is ")
	fmt.Println((avgNum / div))
	return (avgNum / div) // divides the answer by count of numbers and returns it
}

/**
 *  WholeStory
 *  that takes the string, and returns a text that is composed from all the text words separated by spaces,
 *  e.g. story called for the string 1-hello-2-world should return text:
 *  "hello world"
 *
 *  @param inn  the string it takes in as an input
 */
func WholeStory(inn string) string {
	if len(inn) < 1 { // if string is empty
		return ""
	}
	if inn[0] == '-' || inn[len(inn)-1] == '-' { // if string starts or ends with '-' dash
		return ""
	}
	j := 0 // counter for bool array containing the states of the strings
	isNum := make([]bool, uint(len(inn)/3))
	for i := 0; i < len(inn); i++ {
		if inn[i] != '-' { // skips check for '-'
			isNum[j] = true // automatically sets it to number
			if !(uint(inn[i]) >= 48 && uint(inn[i]) <= 57) {
				isNum[j] = false // switches if it isnt a number
			}
		} else {
			j++ // changes check to next string
		}
	}
	strInn := strings.Split(inn, "-")
	if isNum[0] == false { // if first string is NAN
		return ""
	}
	var rString string
	fmt.Print("The string with numbers replaced :[ ")
	for i := 0; i <= j; i++ {
		if isNum[i] == true { // if its a number
			if i != 0 && i != j {
				strInn[i] = " "
			} else {
				strInn[i] = ""
			}
		}
		rString += strInn[i]
	}
	fmt.Print(rString)
	fmt.Println(" ]")
	return (rString) // the string with dashes replaced
}

/**
 *  StoryStats
 *  returns four things:
 *	- the shortest word
 *	- the longest word
 *	- the average word length
 *	- the list (or empty list) of all words from the story that have the length the same as the average length rounded up and down.
 *
 *  @param inn  the string it takes in as an input
 */
func StoryStats(inn string) (string, string, float64, []string) {
	if len(inn) < 1 { // if string is empty
		return "", "", 0.0, []string{""}
	}
	if inn[0] == '-' || inn[len(inn)-1] == '-' { // if string starts or ends with '-' dash
		return "", "", 0.0, []string{""}
	}
	j := 0 // counter for bool array containing the states of the strings
	isNum := make([]bool, uint(len(inn)/3))
	for i := 0; i < len(inn); i++ {
		if inn[i] != '-' { // skips check for '-'
			isNum[j] = true // automatically sets it to number
			if !(uint(inn[i]) >= 48 && uint(inn[i]) <= 57) {
				isNum[j] = false // switches if it isnt a number
			}
		} else {
			j++ // changes check to next string
		}
	}
	// splitting string into strings
	strInn := strings.Split(inn, "-")
	if isNum[0] == false { // if first string is NAN
		return "", "", 0.0, []string{""}
	}
	//handler for string length
	shortestString := strInn[1]
	longestString := strInn[1]
	for i := 0; i < len(strInn); i++ {
		if isNum[i] == false {
			if len(shortestString) > len(strInn[i]) {
				shortestString = strInn[i]
			}
			if len(longestString) < len(strInn[i]) {
				longestString = strInn[i]
			}
		}
	}
	//handler for average word length
	counter := 0
	averageLength := 0.0

	for i := 0; i < len(strInn); i++ {
		if isNum[i] == false {
			averageLength += float64(len(strInn[i])) // adds it to the answer
			counter++
		}
	}
	if counter > 0 {
		//	fmt.Print("counter:", averageLength/float64(counter))
		averageLength = float64(averageLength) / float64(counter)
	}
	rString := []string{}
	// handler for list rounding
	for i := 0; i < len(strInn); i++ {

		if isNum[i] == false {
			fmt.Println(averageLength, ":", math.Round(averageLength))
			if float64(len(strInn[i])) == math.Floor(averageLength) {
				rString = append(rString, strInn[i])
			} else if float64(len(strInn[i])) == math.Ceil(averageLength) {
				rString = append(rString, strInn[i])
			}
		}
	}

	fmt.Print(
		"\t-Shortest string: ", shortestString, "\n",
		"\t-Longest string : ", longestString, "\n",
		"\t-Average Length : ", averageLength, "\n",
		"\t-List(between ", math.Floor(averageLength), " - ", math.Ceil(averageLength), "): \n",
	)
	for i := 0; i < len(rString); i++ {
		fmt.Println("\t\t - ", rString[i])
	}

	fmt.Println()
	return shortestString, longestString, averageLength, rString // divides the answer by count of numbers and returns it
}
