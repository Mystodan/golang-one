package main

import (
	"flag"
	"fmt"
	assignment1 "golang1/pckg"
)

func main() {
	var flagValidity bool
	flag.BoolVar(&flagValidity, "test", false, "takes the string as an input, and returns a boolean value if it complies with the format 23-ab-48-caba-56-haha")
	var flagAverage bool
	flag.BoolVar(&flagAverage, "avr", false, "takes the string, and returns the average number from all the numbers")
	var flagWStory bool
	flag.BoolVar(&flagWStory, "story", false, "takes the string, and replaces the numbers with ' 'spaces")
	var flagSStats bool
	flag.BoolVar(&flagSStats, "stats", false, "takes the string, and returns four things: the shortest- longest words, in addition to the average wordlength and lists all the words with average w. len.(rounded)")

	flag.Parse()

	testString := []string{
		"23-ab-48-caba-56-haha",
		"",
		"420-this works-1337-hi-69",
		"ab-12-cava-41-lol-32",
		"-12-ab-",
		"1-hello-2-world",
	}

	fmt.Print("\n\n")
	var i int
	for i = 0; i < len(testString); i++ {
		if flagValidity {
			if assignment1.TestValidity(testString[i]) {
				fmt.Println("[Success]", "('", testString[i], "')")
				if flagAverage {
					fmt.Print("\t-")
					assignment1.AverageNumber(testString[i])
				}
				if flagWStory {
					fmt.Print("\t-")
					assignment1.WholeStory(testString[i])
				}
				if flagSStats {
					assignment1.StoryStats(testString[i])
				}
			} else {
				fmt.Println("[Failure]", "('", testString[i], "')")
			}
			fmt.Println()
		} else {
			if flagAverage {
				if assignment1.TestValidity(testString[i]) {
					fmt.Println("[Success]", "('", testString[i], "')")
					fmt.Print("\t-")
					assignment1.AverageNumber(testString[i])
					fmt.Print("\n")
				}
			}
			if flagWStory {
				if assignment1.TestValidity(testString[i]) {
					fmt.Println("[Success]", "('", testString[i], "')")
					fmt.Print("\t-")
					assignment1.WholeStory(testString[i])
					fmt.Print("\n")
				}
			}
			if flagSStats {
				if assignment1.TestValidity(testString[i]) {
					fmt.Println("[Success]", "('", testString[i], "')")
					assignment1.StoryStats(testString[i])
					fmt.Print("\n")
				}
			}
		}
	}
}
