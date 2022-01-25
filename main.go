package main

import (
	"flag"
	"fmt"
	assignment1 "golang1/pckg"
)

func main() {
	var flagValidity bool
	flag.BoolVar(&flagValidity, "test", false, "takes the string as an input, and returns a boolean value if it complies with the format 23-ab-48-caba-56-haha")

	flag.Parse()

	if flagValidity {
		testString := []string{
			"23-ab-48-caba-56-haha",
			"",
			"420-this works-1337-hi-69",
			"ab-12-cava-41-lol-32",
			"-12-ab-",
		}
		var i int
		for i = 0; i < len(testString); i++ {
			if assignment1.TestValidity(testString[i]) {
				fmt.Println("[Success]", "('", testString[i], "')")
			} else {
				fmt.Println("[Failure]", "('", testString[i], "')")
			}
		}
	}

}
