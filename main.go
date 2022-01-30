package main

import (
	"flag"
	"fmt"
	assignment1 "golang1/pckg"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	var flagValidity bool
	flag.BoolVar(&flagValidity, "test", false, "takes the string as an input, and returns a boolean value if it complies with the format 23-ab-48-caba-56-haha")
	var flagAverage bool
	flag.BoolVar(&flagAverage, "avr", false, "takes the string, and returns the average number from all the numbers")
	var flagWStory bool
	flag.BoolVar(&flagWStory, "story", false, "takes the string, and replaces the numbers with ' 'spaces")
	var flagSStats bool
	flag.BoolVar(&flagSStats, "stats", false, "takes the string, and returns four things: the shortest- longest words, in addition to the average wordlength and lists all the words with average w. len.(rounded)")
	var flagGenTrue bool
	flag.BoolVar(&flagGenTrue, "gen_true", false, " returns a string value which complies with the format f. ex. 23-ab-48-caba-56-haha")
	var flagGenFalse bool
	flag.BoolVar(&flagGenFalse, "gen_false", true, "returns a string value which does not complies with the format f. ex. -av-a2-48-caba-56-haha")

	flag.Parse()

	var testString []string
	if len(flag.Arg(0)) > 0 {
		testString = []string{
			"23-ab-48-caba-56-haha",
			"",
			"420-this works-1337-hi-69",
			"ab-12-cava-41-lol-32",
			"-12-ab-",
			"1-hello-2-world",
			assignment1.Generate(true),
			assignment1.Generate(false),
			assignment1.Generate(true),
			flag.Arg(0),
		}
	} else {
		testString = []string{
			"23-ab-48-caba-56-haha",
			"",
			"420-this works-1337-hi-69",
			"ab-12-cava-41-lol-32",
			"-12-ab-",
			"1-hello-2-world",
			assignment1.Generate(true),
			assignment1.Generate(false),
			assignment1.Generate(true),
		}
	}

	fmt.Print("\n\n")
	if flagGenTrue {
		fmt.Println(assignment1.Generate(true))
	} else {
		fmt.Println(assignment1.Generate(false))
	}
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
